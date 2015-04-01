package main

import "github.com/gorilla/mux"

func (api *Api) InitRoutes() {

	api.Router = mux.NewRouter()

	api.Router.Handle("/widgets", apiHandler{api, WidgetsIndex}).Methods("GET")
	api.Router.Handle("/widgets", apiHandler{api, WidgetsCreate}).Methods("POST")
	api.Router.Handle("/widgets/{id}", apiHandler{api, WidgetsShow}).Methods("GET")
	api.Router.Handle("/widgets/{id}", apiHandler{api, WidgetsUpdate}).Methods("PUT")
	api.Router.Handle("/widgets/{id}", apiHandler{api, WidgetsDestroy}).Methods("DELETE")

}
