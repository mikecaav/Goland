package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	CONN_HOST        = "localhost"
	CONN_PORT        = "8080"
	DRIVER_NAME      = "mysql"
	DATA_SOURCE_NAME = "root:admin@/test"
)

var db *sql.DB
var connectionError error

func init() {
	db, connectionError = sql.Open(DRIVER_NAME, DATA_SOURCE_NAME)
	if connectionError != nil {
		log.Fatal("Cannot connect to the DB", connectionError)
	}
}

func getCurrentDb(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT DATABASE() as db")
	if err != nil {
		log.Println("error on query: ", err)
		return
	}
	var dbName string
	for rows.Next() {
		rows.Scan(&dbName)
	}
	fmt.Fprintf(w, "dbName: %s", dbName)
}

func main() {
	defer db.Close()
	var router = mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/get-db", getCurrentDb)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
	if err != nil {
		log.Fatal("Server could not be initialized: ", err)
	}
}
