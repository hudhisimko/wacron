package messenger

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/go-jet/jet/v2/mysql"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"simko/wacron/mysql/jetgen/simko_data/model"
	"simko/wacron/mysql/jetgen/simko_data/table"
	enumgw "simko/wacron/mysql/jetgen/simko_gateway/enum"
	modelgw "simko/wacron/mysql/jetgen/simko_gateway/model"
	tablegw "simko/wacron/mysql/jetgen/simko_gateway/table"
	"simko/wacron/worker"
)

var (
	PendingMessagesIterationLimit    = 20
	PendingMessagesIterationInterval = 20 * time.Second
)

type OperationMode struct {
	Mode  string   `koanf:"mode"` // ""/all | deny | allow
	Allow []string `koanf:"allow"`
	Deny  []string `koanf:"deny"`
}

// Messenger fetch org configs and messages, then send based on the config
//
// Notes: Org and Tenant ic actually similar thing.
// Org is returned object or config of tenant.
// Tenant itself is from real live perspective
type Messenger struct {
	DataSchema    string
	GatewaySchema string
	DataDB        *sql.DB
	GatewayDB     *sql.DB

	maxWorker  int
	workers    map[string]*worker.Worker
	workerLock sync.Mutex

	OperationMode OperationMode

	log zerolog.Logger

	done chan struct{}

	orgProcess     map[string]struct{}
	orgProcessLock sync.Mutex
}

func NewMessenger(dataDB, gatewayDB *sql.DB, dataSchema, gatewaySchema string, workerPerTenant int) *Messenger {
	m := Messenger{
		DataSchema:    dataSchema,
		GatewaySchema: gatewaySchema,
		DataDB:        dataDB,
		GatewayDB:     gatewayDB,
		maxWorker:     workerPerTenant,

		workers:    make(map[string]*worker.Worker),
		orgProcess: make(map[string]struct{}),
	}
	return &m
}

// Start the messenger
func (m *Messenger) Start() error {
	m.log = log.With().
		Str("logger", "messenger").
		Logger()
	m.done = make(chan struct{})

	go func() {
	loop:
		for {
			select {
			case <-m.done:
				break loop
			default:

			}

			m.updateSleeping()
			m.iterate()
			time.Sleep(PendingMessagesIterationInterval)
		}
	}()

	return nil
}

func (m *Messenger) Stop() error {
	close(m.done)

	var wg sync.WaitGroup
	wg.Add(len(m.workers))
	for _, w := range m.workers {
		go func(w *worker.Worker) {
			<-w.Stop()
			wg.Done()
		}(w)
	}
	wg.Wait()

	return nil
}

func (m *Messenger) iterate() {
	lwa := table.ZLembagaWhatsapp.FromSchema(m.DataSchema)

	orgsLookup := mysql.Bool(true)
	switch {
	case m.OperationMode.Mode == "allow":
		var ids []mysql.Expression
		for _, v := range m.OperationMode.Allow {
			ids = append(ids, mysql.String(v))
		}

		orgsLookup = orgsLookup.AND(lwa.LembagaID.IN(ids...))

	case m.OperationMode.Mode == "deny":
		var ids []mysql.Expression
		for _, v := range m.OperationMode.Allow {
			ids = append(ids, mysql.String(v))
		}

		orgsLookup = orgsLookup.AND(lwa.LembagaID.NOT_IN(ids...))
	}

	orgsStmt := lwa.SELECT(lwa.LembagaID, lwa.KodeCab).WHERE(mysql.AND(
		orgsLookup,
		lwa.Ready,
	))

	var orgs []model.ZLembagaWhatsapp
	if err := orgsStmt.Query(m.DataDB, &orgs); err != nil {
		return
	}

	for _, org := range orgs {
		m.orgProcessLock.Lock()
		_, ok := m.orgProcess[orgKey(org)]
		if ok {
			m.orgProcessLock.Unlock()
			continue
		}
		m.orgProcess[orgKey(org)] = struct{}{}
		m.orgProcessLock.Unlock()

		go m.iterateOrg(org)
	}
}

func (m *Messenger) iterateOrg(org model.ZLembagaWhatsapp) {
	msgs, err := m.pendingMessages(org)
	if err != nil {
		m.log.Err(err).Send()
	}

	var wg sync.WaitGroup
	wg.Add(len(msgs))

	for _, msg := range msgs {
		msg := msg

		m.worker(org.LembagaID, org.KodeCab).AddTask(&worker.Task{
			ID: int(msg.ID),
			Run: func(workerID int) {
				// metric: add received messages
				messageReceivedTotal.WithLabelValues(msg.KodeLembaga, msg.KodeCab).Add(1)

				m.sendMessage(msg)
				wg.Done()
			},
		})
	}

	wg.Wait()

	// release process lock
	m.orgProcessLock.Lock()
	defer m.orgProcessLock.Unlock()
	delete(m.orgProcess, orgKey(org))
}

func (m *Messenger) worker(tenantID, branch string) *worker.Worker {
	key := fmt.Sprint(tenantID, "+", branch)

	m.workerLock.Lock()
	defer m.workerLock.Unlock()

	w, ok := m.workers[key]
	if ok {
		return w
	}

	// run new worker
	w = worker.NewWorker(m.maxWorker)
	w.Run()

	log.Info().
		Str("tenant_id", tenantID).
		Str("branch", branch).
		Str("key", key).
		Int("count", m.maxWorker).
		Msg("starting worker")

	m.workers[key] = w
	return w
}

func (m *Messenger) pendingMessages(org model.ZLembagaWhatsapp) ([]modelgw.Pesan, error) {
	Pesan := tablegw.Pesan.FromSchema(m.GatewaySchema)
	lwa := table.ZLembagaWhatsapp.FromSchema(m.DataSchema)

	msgStmt := func(filter mysql.BoolExpression) mysql.SelectStatement {
		stmt := Pesan.
			SELECT(Pesan.AllColumns).
			FROM(Pesan.INNER_JOIN(
				// Select which has config only
				lwa,
				lwa.LembagaID.EQ(Pesan.KodeLembaga),
			)).
			WHERE(mysql.AND(
				Pesan.Status.EQ(enumgw.PesanStatus.MenungguJadwal),
				Pesan.Proses.EQ(mysql.Int(0)),
				Pesan.Jadwal.GT_EQ(mysql.RawTimestamp("CURDATE()")),
				filter,
			)).
			ORDER_BY(Pesan.Jadwal.ASC()).
			LIMIT(int64(PendingMessagesIterationLimit))

		return stmt
	}

	stmt := msgStmt(mysql.AND(
		Pesan.KodeLembaga.EQ(mysql.String(org.LembagaID)),
		Pesan.KodeCab.EQ(mysql.String(org.KodeCab)),
	))

	var models []modelgw.Pesan
	if err := stmt.Query(m.GatewayDB, &models); err == nil {
		log.Info().
			Str("tenant", org.LembagaID).
			Str("branch", org.KodeCab).
			Int("count", len(models)).
			Msg("fetched messages")
	}

	return models, nil
}

func (m *Messenger) updateSleeping() {
	// update "should-be-sleeping" provider users
	lwa := table.ZLembagaWhatsapp.FromSchema(m.DataSchema)
	updateStmt := lwa.UPDATE(lwa.LastSleep).
		SET(mysql.NOW()).
		WHERE(mysql.AND(
			lwa.RunDur.GT_EQ(lwa.SleepAfter),
		))

	updateStmt.Exec(m.DataDB)
}

func (m *Messenger) updateLastSent(msg modelgw.Pesan) error {
	lwa := table.ZLembagaWhatsapp.FromSchema(m.DataSchema)

	tx, err := m.GatewayDB.Begin()
	if err != nil {
		return err
	}

	// update slow down throttle last sent
	where := mysql.AND(
		lwa.LembagaID.EQ(mysql.String(msg.KodeLembaga)),
		lwa.KodeCab.EQ(mysql.String(msg.KodeCab)),
	)
	lockStmt := lwa.SELECT(lwa.AllColumns).WHERE(where).FOR(mysql.UPDATE())
	updateStmt := lwa.UPDATE(lwa.LastSent).
		SET(mysql.NOW()).
		WHERE(where)

	if _, err := lockStmt.Exec(tx); err != nil {
		tx.Rollback()
		return err
	}

	if _, err := updateStmt.Exec(tx); err != nil {
		return err
	}

	return tx.Commit()
}

func (m *Messenger) sendMessage(msg modelgw.Pesan) {
	setFailed := func(errs ...error) {
		msg.Status = modelgw.PesanStatus_Gagal
		m.updateMessageStatus(&msg)

		// metric: add failed
		messageFailedTotal.
			WithLabelValues(msg.KodeLembaga, msg.KodeCab, errors.Join(errs...).Error()).
			Add(1)
	}

	// check quota if 0 fail the job
	quota, err := m.getQuota(msg.KodeLembaga)
	if err != nil {
		// setFailed()
		// don't fail wait later
		log.Err(err).Str("tenant", msg.KodeLembaga).Msg("error quota check")
		return
	}

	if quota <= 0 {
		msg.Status = modelgw.PesanStatus_QuotaHabis
		m.updateMessageStatus(&msg)
		return
	}

	// get org configuration
	lwa := table.ZLembagaWhatsapp.FromSchema(m.DataSchema)
	lwp := table.ZLembagaWhatsappProvider.FromSchema(m.DataSchema)

	stmt := lwa.SELECT(lwa.AllColumns, lwp.AllColumns).
		FROM(lwa.INNER_JOIN(
			lwp,
			lwp.Name.EQ(lwa.Provider),
		)).
		WHERE(mysql.AND(
			lwa.LembagaID.EQ(mysql.String(msg.KodeLembaga)),
			lwa.KodeCab.EQ(mysql.String(msg.KodeCab)),

			// Ensure that we use time throttle and sleep condition
			lwa.Ready,
		)).
		LIMIT(1)

	var conf struct {
		Config   model.ZLembagaWhatsapp
		Provider model.ZLembagaWhatsappProvider
	}
	if err := stmt.Query(m.DataDB, &conf); err != nil {
		// setFailed()
		// don't fail here
		return
	}

	log.Debug().Int32("message_id", msg.ID).Caller().Send()

	// set use time right after that
	if err := m.updateLastSent(msg); err != nil {
		log.Debug().Caller().Err(err).Msg("maybe table locked")
		return
	}

	// stop send message if it get expire date
	if err := m.stopMessage(); err != nil {
		log.Error().Caller().Err(err).Msg("Failed update expire date")
	}

	// parse additional data
	var additionalData map[string]string
	if err := json.Unmarshal([]byte(conf.Config.Data), &additionalData); err != nil {
		// setFailed()
		// don't fail here
		return
	}
	var providerHeader map[string]string
	if err := json.Unmarshal([]byte(conf.Provider.Header), &providerHeader); err != nil {
		return
	}
	var confHeader map[string]string
	if err := json.Unmarshal([]byte(conf.Config.Header), &confHeader); err != nil {
		return
	}

	// create provider
	p := NewHTTPProvider()
	p.URL = conf.Provider.URL
	p.Method = conf.Provider.Method
	p.ContentType = conf.Provider.Content
	p.MessageTmpl = conf.Provider.Data
	p.Headers = map[string]string{}

	// combine header from provider and org-provider-config
	for k, v := range providerHeader {
		p.Headers[k] = v
	}
	for k, v := range confHeader {
		p.Headers[k] = v
	}

	if err := p.SendMessage(Message{
		From: msg.Sender,
		To:   msg.Nomor,
		Body: msg.Pesan,
	}, additionalData); err != nil {
		log.Err(err).
			Int("id", int(msg.ID)).
			Str("tenant", msg.KodeLembaga).
			Str("br", msg.KodeCab).
			Msg("unable send message")

		var netError net.Error
		if errors.As(err, &netError) {
			return
		}

		if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) {
			return
		}

		setFailed(err)
		return
	}

	msg.Status = modelgw.PesanStatus_Terkirim
	m.updateMessageStatus(&msg)
	msg.Proses = 1
	m.updateMessageProses(&msg)
	m.reduceQuota(msg.KodeLembaga, 1)

	// metric: track success
	messageSentTotal.
		WithLabelValues(msg.KodeLembaga, msg.KodeCab, conf.Config.Provider).
		Add(1)
}

func (m *Messenger) updateMessageProses(msg *modelgw.Pesan) error {
	Pesan := tablegw.Pesan.FromSchema(m.GatewaySchema)
	stmt := Pesan.UPDATE(Pesan.Proses).
		MODEL(msg).
		WHERE(Pesan.ID.EQ(mysql.Int(int64(msg.ID))))
	if _, err := stmt.Exec(m.GatewayDB); err != nil {
		return err
	}
	return nil
}

func (m *Messenger) updateMessageStatus(msg *modelgw.Pesan) error {
	Pesan := tablegw.Pesan.FromSchema(m.GatewaySchema)
	stmt := Pesan.UPDATE(Pesan.Status).
		MODEL(msg).
		WHERE(Pesan.ID.EQ(mysql.Int(int64(msg.ID))))
	if _, err := stmt.Exec(m.GatewayDB); err != nil {
		return err
	}
	return nil
}

func (m *Messenger) getQuota(tenantID string) (int, error) {
	type dst struct {
		Quota int `alias:"lembaga.quota"`
	}

	Lembaga := table.Lembaga.FromSchema(m.DataSchema)
	stmt := Lembaga.
		SELECT(Lembaga.KuotaAwalWa.SUB(Lembaga.KuotaTransWa).AS("lembaga.quota")).
		WHERE(Lembaga.Memberid.EQ(mysql.String(tenantID)))

	var dest dst
	if err := stmt.Query(m.DataDB, &dest); err != nil {
		return 0, err
	}

	return dest.Quota, nil
}

func (m *Messenger) reduceQuota(tenantID string, n int) error {
	Lembaga := table.Lembaga.FromSchema(m.DataSchema)
	stmt := Lembaga.
		UPDATE(Lembaga.KuotaTransWa, Lembaga.KuotaAkhirWa).
		SET(
			Lembaga.KuotaTransWa.SET(Lembaga.KuotaTransWa.ADD(mysql.Int(int64(n)))),
			Lembaga.KuotaAkhirWa.SET(Lembaga.KuotaAwalWa.SUB(Lembaga.KuotaTransWa)),
		).
		WHERE(Lembaga.Memberid.EQ(mysql.String(tenantID)))

	if _, err := stmt.Exec(m.DataDB); err != nil {
		return err
	}
	return nil
}

func (m *Messenger) stopMessage() error {
	lwa := table.ZLembagaWhatsapp.FromSchema(m.DataSchema)
	setStmt := lwa.UPDATE(lwa.Ready).SET(mysql.Int(0)).WHERE(lwa.DateExpired.GT_EQ(mysql.CURRENT_DATE()).AND(lwa.Ready))

	if _, err := setStmt.Exec(m.DataDB); err != nil {
		return err
	}

	return nil
}

func orgKey(org model.ZLembagaWhatsapp) string {
	return fmt.Sprint(org.LembagaID + "+" + org.KodeCab)
}
