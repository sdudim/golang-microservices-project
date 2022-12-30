package main

import (
	"log-service/data"
	"net/http"
)

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	// read josn into var
	var requestedPayload JSONPayload
	_ = app.readJSON(w, r, &requestedPayload)

	// insert data
	event := data.LogEntry{
		Name: requestedPayload.Name,
		Data: requestedPayload.Data,
	}

	err := app.Models.LogEntry.Insert(event)

	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := jsonResponse{ // this struct object defined in the helpers.go file
		Error:   false,
		Message: "logged(Log service)",
	}

	app.writeJSON(w, http.StatusAccepted, resp)

}
