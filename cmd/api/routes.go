// Filename

package main 

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)


func (app *application) routes() *httprouter.Router {
	// Create
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/info", app.createInfoHandler)
	router.HandlerFunc(http.MethodGet, "/v1/random/:id", app.showRandHandler)

	return router
}