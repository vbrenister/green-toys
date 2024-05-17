package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFound)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowed)

	router.HandlerFunc(http.MethodGet, "/toys/:id", app.getToyByID)
	router.HandlerFunc(http.MethodPost, "/toys", app.createToy)
	router.HandlerFunc(http.MethodGet, "/toys", app.getAllToys)
	router.HandlerFunc(http.MethodPatch, "/toys/:id", app.rentToy)

	return app.recoverPanic(app.requestLogging(router))
}
