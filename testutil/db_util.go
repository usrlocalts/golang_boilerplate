package testutil

import (
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"github.com/jmoiron/sqlx"
	"golang_boilerplate/config"
	"testing"
	"github.com/stretchr/testify/assert"
	"golang_boilerplate/logger"
	_ "github.com/lib/pq"
	"fmt"
)

var TestDBUrl = "postgres://boilerplate@localhost:5432/golang_boilerplate_test?sslmode=disable"

func SetupMockTestDB(t *testing.T) (*sqlx.DB, sqlmock.Sqlmock) {
	config := config.Load()
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	sqlxDB := sqlx.NewDb(db, "postgres")

	sqlxDB.SetMaxOpenConns(config.DatabaseConfig().MaxConn())
	return sqlxDB, mock
}

func CloseTestDB(db *sqlx.DB) {
	db.Close()
}

func SetupTestDB(logger logger.Log) *sqlx.DB {
	db, err := sqlx.Open("postgres", TestDBUrl)
	if err != nil {
		logger.Fatalf("failed to load the database: %s", err)
	}
	if err = db.Ping(); err != nil {
		logger.Fatalf("ping to the database host failed: %s", err)
	}

	db.SetMaxOpenConns(50)
	return db
}

func WithCleanTable(block func(), db *sqlx.DB, tableName string) {

	db.Exec(fmt.Sprintf("delete from %v", tableName))
	block()
	db.Exec(fmt.Sprintf("delete from %v", tableName))
	db.Close()
}
