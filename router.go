package main

import "github.com/gorilla/mux"

func (api *Api) InitRoutes() {

	api.Router = mux.NewRouter()
	api.Router.HandleFunc("/widgets", WidgetsIndex).Methods("GET")
	api.Router.HandleFunc("/widgets", WidgetsCreate).Methods("POST")
	api.Router.HandleFunc("/widgets/{id}", WidgetsShow).Methods("GET")
	api.Router.HandleFunc("/widgets/{id}", WidgetsUpdate).Methods("PUT")
	api.Router.HandleFunc("/widgets/{id}", WidgetsDestroy).Methods("DELETE")

}
