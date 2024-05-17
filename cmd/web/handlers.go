package main

import (
	"errors"
	"net/http"

	"github.com/vbrenister/green-toys/internal/models"
	"github.com/vbrenister/green-toys/internal/validation"
)

func (app *application) getToyByID(w http.ResponseWriter, r *http.Request) {
	id, err := app.readParamID(r)

	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	toy, err := app.toys.GetByID(id)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrRecordNotFound):
			app.notFound(w, r)
		default:
			app.internalServerError(w, r, err)
		}

		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"toy": toy})
	if err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) createToy(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		Category    string  `json:"category"`
	}

	err := app.readJSON(r, &input)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	toy := &models.Toy{
		Title:       input.Title,
		Description: input.Description,
		Price:       input.Price,
		Category:    input.Category,
	}

	v := validation.New()

	if models.ValidateToy(v, toy); !v.Valid() {
		app.failedValidation(w, r, v.Errors)
		return
	}

	err = app.toys.Create(toy)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"toy": toy})
	if err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) getAllToys(w http.ResponseWriter, r *http.Request) {
	toys, err := app.toys.GetAll()
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"toys": toys})
	if err != nil {
		app.internalServerError(w, r, err)
	}
}
