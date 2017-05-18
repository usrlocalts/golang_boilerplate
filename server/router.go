package server

import (
	"net/http"
	"path"

	"github.com/gorilla/mux"
	h "golang_boilerplate/handler"
	"golang_boilerplate/services"
	"golang_boilerplate/posts_creation"
	"golang_boilerplate/appcontext"
	"github.com/jmoiron/sqlx"
)

func Router(ctx *appcontext.AppContext, db *sqlx.DB) *mux.Router {
	logger := ctx.GetLogger()
	config := ctx.GetConfig()
	services := services.New(logger, db)
	router := mux.NewRouter()
	router.HandleFunc("/ping", h.PingHandler).Methods("GET")
	router.HandleFunc("/posts", posts_creation.PostsCreationHandler(services.PostsService, logger, ctx.GetConfig().Newrelic().Enabled)).Methods("POST")
	// Swagger Docs
	router.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", http.FileServer(http.Dir(config.DocsPath()))))
	router.HandleFunc("/swagger.yml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path.Join(config.DocsPath(), "swagger.yml"))
	})
	router.NotFoundHandler = http.HandlerFunc(h.NotFoundHandler)
	return router
}

func apiDocHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/docs" {
		http.Redirect(w, r, "/docs/", http.StatusFound)
		return
	}
}
