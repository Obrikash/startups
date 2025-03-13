package main

import "net/http"

func (app *application) routes() http.Handler {
    router := http.NewServeMux()
    
    router.HandleFunc("GET /api/startups", app.listStartupsHandler)
    router.HandleFunc("GET /api/startups/{id}", app.showStartupHandler)

    return router
}
