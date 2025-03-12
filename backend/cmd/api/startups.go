package main

import (
	"net/http"
	"startups/internal/data"
	"startups/internal/validator"
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
