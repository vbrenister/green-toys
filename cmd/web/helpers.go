package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type envelope map[string]any

func (app *application) writeJSON(w http.ResponseWriter, statusCode int, data envelope) error {
	json, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(json)

	return nil
}

// TODO: Handle JSON decoding errors
func (app *application) readJSON(r *http.Request, v any) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func (app *application) readParamID(r *http.Request) (int, error) {
	idParam := httprouter.ParamsFromContext(r.Context()).ByName("id")

	id, err := strconv.Atoi(idParam)
	if id < 1 || err != nil {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}
