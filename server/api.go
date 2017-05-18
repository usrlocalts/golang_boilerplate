package server

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Sirupsen/logrus"
	nrgorilla "github.com/newrelic/go-agent/_integrations/nrgorilla/v1"
	"golang_boilerplate/appcontext"
	"golang_boilerplate/logger"
	"github.com/urfave/negroni"
	"github.com/jmoiron/sqlx"
)

func StartAPIServer(ctx *appcontext.AppContext, db *sqlx.DB) {
	logger := ctx.GetLogger()
	config := ctx.GetConfig()
	newRelicApp := ctx.NewrelicApp()

	router := Router(ctx, db)
	handlerFunc := router.ServeHTTP
	server := negroni.New(negroni.NewRecovery())
	server.Use(httpStatLogger(logger))
	server.UseHandlerFunc(handlerFunc)
	portInfo := ":" + strconv.Itoa(config.Port())
	logger.Info("Starting Golang Boilerplate and listening on port: ", config.Port())
	http.ListenAndServe(portInfo, nrgorilla.InstrumentRoutes(router, newRelicApp))
	server.Run(portInfo)
}

func httpStatLogger(logger logger.Log) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		startTime := time.Now()
		next(rw, r)
		responseTime := time.Now()
		deltaTime := responseTime.Sub(startTime).Seconds() * 1000

		if r.URL.Path != "/ping" {
			logger.WithFields(logrus.Fields{
				"RequestTime":   startTime.Format(time.RFC3339),
				"ResponseTime":  responseTime.Format(time.RFC3339),
				"DeltaTime":     deltaTime,
				"RequestUrl":    r.URL.Path,
				"RequestMethod": r.Method,
				"RequestProxy":  r.RemoteAddr,
				"RequestSource": r.Header.Get("X-FORWARDED-FOR"),
			}).Debug("Http Logs")
		}
	})
}
