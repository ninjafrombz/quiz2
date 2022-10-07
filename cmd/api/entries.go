// Filename:
package main

import (
	"fmt"
	"net/http"

	"quiz2.desireamagwula.net/internals/data"
	"quiz2.desireamagwula.net/internals/validator"
)

// CreateSchoolHandler for the POST /v1/schools" endpoint

func (app *application) createSchoolHandler(w http.ResponseWriter, r *http.Request) {
	// Our target decode destination fmt.Fprintln(w, "create a new school..")
	var input struct {
		Name    string   `json:"name"`
		Contact string   `json:"contact"`
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
	school := &data.School{
		Name:    input.Name,
		Contact: input.Contact,
		Phone:   input.Phone,
		Email:   input.Email,
		Address: input.Address,
		Mode:    input.Mode,
	}

	//Initialize a new validator instance
	v := validator.New()

	// Check the map to determine if there were any validation errors
	if data.ValidateSchool(v, school); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	// Display the request
	fmt.Fprintf(w, "%+v\n", input)

}

func (app *application) showSchoolHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"school": school}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
