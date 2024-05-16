package main

import (
	"errors"
	"net/http"

	"github.com/vbrenister/green-toys/internal/models"
)

func (app *application) getToyByID(w http.ResponseWriter, r *http.Request) {
	id, err := app.readParamID(r)

	if err != nil {
		app.badRequest(w, err)
		return
	}

	toy, err := app.toys.GetByID(id)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrRecordNotFound):
			app.notFound(w)
		default:
			app.internalServerError(w, err)
		}

		return
	}

	err = app.writeJSON(w, http.StatusOK, toy)
	if err != nil {
		app.internalServerError(w, err)
	}
}

// TODO: Add validation rules
func (app *application) createToy(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		Category    string  `json:"category"`
	}

	err := app.readJSON(r, &input)
	if err != nil {
		app.badRequest(w, err)
		return
	}

	toy := &models.Toy{
		Title:       input.Title,
		Description: input.Description,
		Price:       input.Price,
		Category:    input.Category,
	}

	err = app.toys.Create(toy)
	if err != nil {
		app.internalServerError(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, toy)
	if err != nil {
		app.internalServerError(w, err)
	}
}

func (app *application) getAllToys(w http.ResponseWriter, r *http.Request) {
	toys, err := app.toys.GetAll()
	if err != nil {
		app.internalServerError(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, toys)
	if err != nil {
		app.internalServerError(w, err)
	}
}
