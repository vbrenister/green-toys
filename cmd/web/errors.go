package main

import "net/http"

func (app *application) logError(r *http.Request, err error) {
	var (
		path   = r.URL.Path
		method = r.Method
	)
	app.logger.Error(err.Error(), "method", method, "path", path)
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
	envelope := envelope{"error": message}

	err := app.writeJSON(w, statusCode, envelope)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (app *application) notFound(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

func (app *application) badRequest(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}
