package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/coopernurse/gorp"
	_ "github.com/lib/pq"
)

func (api *Api) InitDB() {

	// example: postgres://user:pass@host/db?sslmode=disable
	db, err := sql.Open("postgres", os.Getenv("PG_CONNECTION_URL"))

	checkErr(err, "sql.Open failed")
	api.DB = &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	api.DB.AddTableWithName(Widget{}, "widgets").SetKeys(true, "Id")

	// fix for production
	fmt.Println("creating tables if they dont' exist")
	err = api.DB.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func RecordNotFoundError(err error) bool {
	return err == sql.ErrNoRows
}
