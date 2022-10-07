// Filename:
package main

import (
	"fmt"
	"net/http"
	

	"quiz2.amagwuladesire.net/internal/data"
	"quiz2.amagwuladesire.net/internal/validator"
)

// CreateSchoolHandler for the POST /v1/schools" endpoint

func (app *application) createInfoHandler(w http.ResponseWriter, r *http.Request) {
	// Our target decode destination fmt.Fprintln(w, "create a new school..")
	var input struct {
		Name    string   `json:"name"`
		Nationality string   `json:"nationality"`
		Phone   string   `json:"phone"`
		Email   string   `json:"email"`
		Address string   `json:"address"`
		Mode    []string `json:"mode"`
	}

	// Initialize a new json.Decoder instance
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Copy the values from the input struct to a new school struct
	entry := &data.Entry{
		Name:    input.Name,
		Nationality: input.Nationality,
		Phone:   input.Phone,
		Email:   input.Email,
		Address: input.Address,
		Mode:    input.Mode,
	}

	//Initialize a new validator instance
	v := validator.New()

	// Check the map to determine if there were any validation errors
	if data.ValidateEntries(v, entry); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	// Display the request
	fmt.Fprintf(w, "%+v\n", input)

}

func (app *application) showRandHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	value:=int(id)
	tools:=&data.Tools{
		Int: value,
	}
	v := validator.New()
	if data.ValidateInt(v,tools); !v.Valid(){
		app.failedValidationResponse(w,r,v.Errors)
		return
	}
	randstring := tools.GenerateRandomString(value)
	data := envelope{
		"id": value,
		"randomstring": randstring,
	}

	err = app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
