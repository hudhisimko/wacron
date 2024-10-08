//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/mysql"
)

var ZLembagaWhatsapp = newZLembagaWhatsappTable("simko_data", "z_lembaga_whatsapp", "")

type zLembagaWhatsappTable struct {
	mysql.Table

	// Columns
	LembagaID     mysql.ColumnString
	NamaLembaga   mysql.ColumnString
	KodeCab       mysql.ColumnString
	Provider      mysql.ColumnString
	Data          mysql.ColumnString
	Header        mysql.ColumnString
	SleepAfter    mysql.ColumnInteger
	SleepDuration mysql.ColumnInteger
	SendDelay     mysql.ColumnInteger
	LastSent      mysql.ColumnTimestamp
	Idle          mysql.ColumnInteger
	LastSleep     mysql.ColumnTimestamp
	Ready         mysql.ColumnBool
	Sleeping      mysql.ColumnBool
	RunDur        mysql.ColumnInteger
	DateExpired   mysql.ColumnDate

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type ZLembagaWhatsappTable struct {
	zLembagaWhatsappTable

	NEW zLembagaWhatsappTable
}

// AS creates new ZLembagaWhatsappTable with assigned alias
func (a ZLembagaWhatsappTable) AS(alias string) *ZLembagaWhatsappTable {
	return newZLembagaWhatsappTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ZLembagaWhatsappTable with assigned schema name
func (a ZLembagaWhatsappTable) FromSchema(schemaName string) *ZLembagaWhatsappTable {
	return newZLembagaWhatsappTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ZLembagaWhatsappTable with assigned table prefix
func (a ZLembagaWhatsappTable) WithPrefix(prefix string) *ZLembagaWhatsappTable {
	return newZLembagaWhatsappTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ZLembagaWhatsappTable with assigned table suffix
func (a ZLembagaWhatsappTable) WithSuffix(suffix string) *ZLembagaWhatsappTable {
	return newZLembagaWhatsappTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newZLembagaWhatsappTable(schemaName, tableName, alias string) *ZLembagaWhatsappTable {
	return &ZLembagaWhatsappTable{
		zLembagaWhatsappTable: newZLembagaWhatsappTableImpl(schemaName, tableName, alias),
		NEW:                   newZLembagaWhatsappTableImpl("", "new", ""),
	}
}

func newZLembagaWhatsappTableImpl(schemaName, tableName, alias string) zLembagaWhatsappTable {
	var (
		LembagaIDColumn     = mysql.StringColumn("lembaga_id")
		NamaLembagaColumn   = mysql.StringColumn("nama_lembaga")
		KodeCabColumn       = mysql.StringColumn("kode_cab")
		ProviderColumn      = mysql.StringColumn("provider")
		DataColumn          = mysql.StringColumn("data")
		HeaderColumn        = mysql.StringColumn("header")
		SleepAfterColumn    = mysql.IntegerColumn("sleep_after")
		SleepDurationColumn = mysql.IntegerColumn("sleep_duration")
		SendDelayColumn     = mysql.IntegerColumn("send_delay")
		LastSentColumn      = mysql.TimestampColumn("last_sent")
		IdleColumn          = mysql.IntegerColumn("idle")
		LastSleepColumn     = mysql.TimestampColumn("last_sleep")
		ReadyColumn         = mysql.BoolColumn("ready")
		SleepingColumn      = mysql.BoolColumn("sleeping")
		RunDurColumn        = mysql.IntegerColumn("run_dur")
		DateExpiredColumn   = mysql.DateColumn("date_expired")
		allColumns          = mysql.ColumnList{LembagaIDColumn, NamaLembagaColumn, KodeCabColumn, ProviderColumn, DataColumn, HeaderColumn, SleepAfterColumn, SleepDurationColumn, SendDelayColumn, LastSentColumn, IdleColumn, LastSleepColumn, ReadyColumn, SleepingColumn, RunDurColumn, DateExpiredColumn}
		mutableColumns      = mysql.ColumnList{LembagaIDColumn, NamaLembagaColumn, KodeCabColumn, ProviderColumn, DataColumn, HeaderColumn, SleepAfterColumn, SleepDurationColumn, SendDelayColumn, LastSentColumn, IdleColumn, LastSleepColumn, ReadyColumn, SleepingColumn, RunDurColumn, DateExpiredColumn}
	)

	return zLembagaWhatsappTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		LembagaID:     LembagaIDColumn,
		NamaLembaga:   NamaLembagaColumn,
		KodeCab:       KodeCabColumn,
		Provider:      ProviderColumn,
		Data:          DataColumn,
		Header:        HeaderColumn,
		SleepAfter:    SleepAfterColumn,
		SleepDuration: SleepDurationColumn,
		SendDelay:     SendDelayColumn,
		LastSent:      LastSentColumn,
		Idle:          IdleColumn,
		LastSleep:     LastSleepColumn,
		Ready:         ReadyColumn,
		Sleeping:      SleepingColumn,
		RunDur:        RunDurColumn,
		DateExpired:   DateExpiredColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
