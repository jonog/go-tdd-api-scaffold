package main

import (
	"log"
	"net/http"

	"github.com/go-gorp/gorp"
	"github.com/gorilla/mux"
)

type Api struct {
	Router *mux.Router
	DB     *gorp.DbMap
}

type apiHandler struct {
	*Api
	handler func(*Api, http.ResponseWriter, *http.Request)
}

func (ah apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ah.handler(ah.Api, w, r)
}

func main() {

	api := &Api{}
	api.InitDB()
	defer api.DB.Db.Close()

	api.InitRoutes()
	http.Handle("/", api.Router)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
