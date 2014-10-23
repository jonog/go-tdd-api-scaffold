package main

import (
	"log"
	"net/http"

	"github.com/coopernurse/gorp"
	"github.com/gorilla/mux"
)

type Api struct {
	Router *mux.Router
	DB     *gorp.DbMap
}

var api Api

func main() {

	api.InitDB()
	defer api.DB.Db.Close()

	api.InitRoutes()
	http.Handle("/", api.Router)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
