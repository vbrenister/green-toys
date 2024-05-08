package main

import (
	"net/http"

	"github.com/vbrenister/green-toys/internal/models"
)

func (app *application) getToyByID(w http.ResponseWriter, r *http.Request) {
	id, err := app.readParamID(r)

	if err != nil {
		app.badRequest(w, err)
		return
	}

	toy := models.Toy{
		ID:          id,
		Title:       "Toy",
		Description: "A toy",
		Price:       10.00,
		Rating:      4.5,
		Condition:   models.GoodCondition,
		Category:    "Toy",
	}

	err = app.writeJSON(w, http.StatusOK, toy)
	if err != nil {
		app.internalServerError(w, err)
	}
}
