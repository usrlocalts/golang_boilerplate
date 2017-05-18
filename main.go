package main

import (
	_ "expvar"
	"os"

	raven "github.com/getsentry/raven-go"
	"github.com/urfave/cli"
	"golang_boilerplate/appcontext"
	"golang_boilerplate/config"
	"golang_boilerplate/console"
	"golang_boilerplate/logger"
	"golang_boilerplate/server"
	"golang_boilerplate/db"
)

func handleInitError(logger logger.Log) {
	if e := recover(); e != nil {
		logger.Fatalf("Failed to load the app due to error : %s", e)
	}
}

func main() {
	config := config.Load()
	logger := logger.SetupLogger(config)
	db := db.NewDB(logger, config.DBConfig().Url(), config.DBConfig().MaxConn())
	ctx := appcontext.NewAppContext(logger, config)

	defer handleInitError(logger)

	if config.Sentry().Enabled() {
		raven.SetDSN(config.Sentry().Dsn())
	}

	clientApp := cli.NewApp()
	clientApp.Name = "golang_boilerplate"
	clientApp.Version = "0.0.1"
	clientApp.Commands = []cli.Command{
		{
			Name:        "start",
			Description: "Start HTTP api server",
			Action: func(c *cli.Context) error {
				server.StartAPIServer(ctx, db)
				return nil
			},
		},
		{
			Name:        "migrate",
			Description: "Run database migrations",
			Action: func(c *cli.Context) error {
				return console.RunDatabaseMigrations(config)
			},
		},
		{
			Name:        "rollback",
			Description: "Rollback latest database migration",
			Action: func(c *cli.Context) error {
				return console.RollbackLatestMigration(config)
			},
		},
	}

	if err := clientApp.Run(os.Args); err != nil {
		panic(err)
	}
}
