package main

import (
	"errors"
	"fmt"
	"net/http"
	"startups/internal/data"
	"startups/internal/validator"
	"strconv"
)

func (app *application) listStartupsHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string
		data.Filters
	}

	qs := r.URL.Query()

	v := validator.New()

	input.Title = app.readString(qs, "title", "")

	input.Page = app.readInt(qs, "page", 1, v)
	input.PageSize = app.readInt(qs, "page_size", 20, v)

	input.Sort = app.readString(qs, "sort", "created_at")
	input.SortSafeList = []string{"created_at", "-created_at", "views", "-views", "title", "-title"}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	startups, metadata, err := app.models.Startups.GetAll(input.Title, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"startups": startups, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

}

func (app *application) showStartupHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	if idString == "" {
		app.badRequestResponse(w, r, errors.New("no id provided"))
		return
	}

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	startup, err := app.models.Startups.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"startup": startup}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) updateViewsHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	if idString == "" {
		app.badRequestResponse(w, r, errors.New("no id provided"))
		return
	}

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Startups.UpdateViews(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (app *application) createStartupHandler(w http.ResponseWriter, r *http.Request) {
    type Slug struct {
            Type    string `json:"type"`
            Current string `json:"current"`
    }

	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Category    string `json:"category"`
        Image       string `json:"image"`
        Slug        Slug
        Author      data.Author
        Pitch       string `json:"pitch"`
	}

    err := app.readJSON(w, r, &input)
    if err != nil {
        app.badRequestResponse(w, r, err)
        return
    }
    fmt.Println("here")
}
