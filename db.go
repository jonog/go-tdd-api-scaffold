package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq"
)

func getConnURL() string {
	var connURL string
	if os.Getenv("PG_CONNECTION_URL") != "" {
		connURL = os.Getenv("PG_CONNECTION_URL")
	} else {
		connURL = "postgres://" +
			os.Getenv("PG_USER_PASS") + "@" +
			os.Getenv("PG_PORT_5432_TCP_ADDR") + ":" +
			os.Getenv("PG_PORT_5432_TCP_PORT") + "/" +
			os.Getenv("PG_DB_NAME") + "?sslmode=disable"
	}
	return connURL
}

func (api *Api) InitDB() {

	// example: postgres://user:pass@host/db?sslmode=disable
	db, err := sql.Open("postgres", getConnURL())

	checkErr(err, "sql.Open failed")
	api.DB = &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	api.DB.AddTableWithName(Widget{}, "widgets").SetKeys(true, "Id")

	// Remove & add migrations in production setup
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
	return err != nil && err == sql.ErrNoRows
}
