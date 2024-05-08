package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) internalServerError(w http.ResponseWriter, err error) {
	app.logger.Error(err.Error())
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) notFound(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func (app *application) badRequest(w http.ResponseWriter, err error) {
	app.logger.Error(err.Error())
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}

func (app *application) writeJSON(w http.ResponseWriter, statusCode int, data any) error {
	json, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(json)

	return nil
}

func (app *application) readParamID(r *http.Request) (int, error) {
	idParam := httprouter.ParamsFromContext(r.Context()).ByName("id")

	id, err := strconv.Atoi(idParam)
	if id < 1 || err != nil {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}
