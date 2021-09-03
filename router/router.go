package router

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/masdikaid-rakamin-homework-movie-service/handler"
)

func NewRouter() http.Handler {
	r := mux.NewRouter()

	movieRouter(r)

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	return loggedRouter
}

func movieRouter(router *mux.Router) {
	router.Methods(http.MethodPost).Path("/").Handler(handler.CreateMovie())
	router.Methods(http.MethodGet).Path("/{slug}").Handler(handler.GetMovie())
	router.Methods(http.MethodDelete).Path("/{slug}").Handler(handler.DeleteMovie())
}
