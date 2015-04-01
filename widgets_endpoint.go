package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func WidgetsIndex(w http.ResponseWriter, r *http.Request) {

	widgets, err := GetAllWidgets()
	if err != nil {
		Error(w, "Internal Server Error", 500)
		return
	}

	widgetsJSON, err := json.Marshal(ExportWidgets(widgets))
	if err != nil {
		Error(w, "Internal Server Error", 500)
		return
	}

	Respond(w, widgetsJSON, 200)
}

func WidgetsCreate(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Error(w, "Internal Server Error", 500)
		return
	}

	params := &WidgetParams{}
	err = json.Unmarshal(body, params)
	if err != nil {
		Error(w, "Internal Server Error", 500)
		return
	}

	widget := BuildWidget(params)
	err = widget.Validate()
	if err != nil {
		Error(w, "Invalid parameters", 400)
		return
	}

	err = widget.Save()
	if err != nil {
		Error(w, "Internal Server Error", 500)
		return
	}

	widgetJSON, err := json.Marshal(widget.Export())
	if err != nil {
		Error(w, "Internal Server Error", 500)
		return
	}

	Respond(w, widgetJSON, 201)
}

func WidgetsShow(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		Error(w, "Internal Server Error", 500)
		return
	}

	widget, err := FindWidget(id)
	if RecordNotFoundError(err) {
		Error(w, "Not Found", 404)
		return
	}

	widgetJSON, err := json.Marshal(widget.Export())
	if err != nil {
		Error(w, "Internal Server Error", 500)
		return
	}

	Respond(w, widgetJSON, 200)

}

func WidgetsUpdate(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Error(w, "Internal Server Error", 500)
		return
	}

	var params WidgetParams
	err = json.Unmarshal(body, &params)
	if err != nil {
		Error(w, "Internal Server Error", 500)
		return
	}

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		Error(w, "Internal Server Error", 500)
		return
	}

	widget, err := FindWidget(id)
	if RecordNotFoundError(err) {
		Error(w, "Not Found", 404)
		return
	}

	widget.Name = params.Name
	err = widget.Validate()
	if err != nil {
		Error(w, "Invalid parameters", 400)
		return
	}

	widget.Name = params.Name
	widget.Save()

	widgetJSON, err := json.Marshal(widget.Export())
	if err != nil {
		Error(w, "Internal Server Error", 500)
		return
	}

	Respond(w, widgetJSON, 200)
}

func WidgetsDestroy(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		Error(w, "Internal Server Error", 500)
		return
	}

	widget, err := FindWidget(id)
	if RecordNotFoundError(err) {
		Error(w, "Not Found", 404)
		return
	}

	err = widget.Delete()
	if err != nil {
		Error(w, "Internal Server Error", 500)
		return
	}
	Respond(w, []byte(`{}`), 200)

}
