package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (*application) routes() http.Handler {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("Hello world!"))
	})
	return router
}
