package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/getsentry/raven-go"
	"golang_boilerplate/logger"
)

func NewDB(logger logger.Log, url string, maxConn int) *sqlx.DB {
	db, err := sqlx.Open("postgres", url)
	if err != nil {
		raven.CaptureError(err, map[string]string{"failed to load the database ": err.Error()})
		logger.Fatalf("failed to load the database: %s", err)
	}

	if err = db.Ping(); err != nil {
		raven.CaptureError(err, map[string]string{"ping to the database host failed ": err.Error()})
		logger.Fatalf("ping to the database host failed: %s", err)
	}

	db.SetMaxOpenConns(maxConn)
	return db
}
