package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/toys/:id", app.getToyByID)
	router.HandlerFunc(http.MethodPost, "/toys", app.createToy)
	router.HandlerFunc(http.MethodGet, "/toys", app.getAllToys)

	return router
}
