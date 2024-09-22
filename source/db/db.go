package db

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-jet/jet/v2/mysql"
	mysqldrv "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
)

type Connection struct {
	Identifier string
	Connection string
}

func connDSN(c Connection, db string) string {
	return fmt.Sprintf("%s/%s?charset=utf8mb4&parseTime=true", c.Connection, db)
}

func Open(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(15)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)
	return db, err
}

func openDB(conns []Connection, dbName string) (*sql.DB, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(len(conns))

	resultCh := make(chan *sql.DB, len(conns))

	// concurrently try open connection
	for _, conn := range conns {
		go func(ctx context.Context, conn Connection, dbName string) {
			defer wg.Done()

			select {
			case <-ctx.Done():
				return
			default:
				db, err := Open(connDSN(conn, dbName))
				if err != nil {
					log.Warn().
						Str("conn_id", conn.Identifier).
						Str("db_name", dbName).
						Err(err).
						Msg("cannot open db connection")
					resultCh <- nil
					return
				}

				log.Info().
					Str("conn_id", conn.Identifier).
					Str("db_name", dbName).
					Msg("db connection established")
				resultCh <- db
				cancel()
			}
		}(ctx, conn, dbName)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for db := range resultCh {
		if db != nil {
			return db, nil
		}
	}

	return nil, fmt.Errorf("no db")
}

// DSNSchema returns db name from dsn
func DSNSchema(dsn string) (string, error) {
	cfg, err := mysqldrv.ParseDSN(dsn)
	if err != nil {
		return "", err
	}
	return cfg.DBName, nil
}

func EnableQueryLog() {
	mysql.SetQueryLogger(func(ctx context.Context, info mysql.QueryInfo) {
		f, l, _ := info.Caller()
		errStr := ""
		if info.Err != nil {
			errStr = info.Err.Error()
		}

		if id := middleware.GetReqID(ctx); id != "" {
			log.Printf("sql: request-id: %s", id)
		}
		log.Printf("[%.3fms] [rows:%d] %s:%d %s %s", float64(info.Duration.Nanoseconds())/1e6, info.RowsProcessed, f, l, errStr, info.Statement.DebugSql())
	})
}
