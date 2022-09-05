package main

import (
	"database/sql"
	"encoding/json"
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

var dbConnection *sql.DB
var errorConenction error

type Employee struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Employees []Employee

var employees Employees

type Route struct {
	name            string
	path            string
	method          string
	handlerFunction http.HandlerFunc
}

type Routes []Route

var routes Routes

func getRoutes() Routes {
	return Routes{
		Route{handlerFunction: getRecords, name: "getRecords", path: "/get", method: "GET"},
		Route{handlerFunction: addRecord, name: "addRecord", path: "/add", method: "POST"},
		Route{handlerFunction: updateRecord, name: "updateRecord", path: "/update", method: "PUT"},
	}
}

func init() {
	employees = Employees{}
	routes = getRoutes()
	dbConnection, errorConenction = sql.Open(DRIVER_NAME, DATA_SOURCE_NAME)
	if errorConenction != nil {
		log.Fatal("The connection to the DB could not be established: ", errorConenction)
	}
}

func AddRoutes(router *mux.Router) *mux.Router {
	for _, route := range routes {
		router.Methods(route.method).Name(route.name).Handler(route.handlerFunction).Path(route.path)
	}
	return router
}

func getRecords(w http.ResponseWriter, r *http.Request) {
	rows, err := dbConnection.Query("SELECT * FROM MyTable")
	if err != nil {
		log.Println("An error ocurred while trying to fetch the values from MyTable")
	}
	var employees Employees
	for rows.Next() {
		var uid int
		var name string
		rows.Scan(&uid, &name)
		employees = append(employees, Employee{Id: uid, Name: name})
	}
	json.NewEncoder(w).Encode(employees)
}

func addRecord(w http.ResponseWriter, r *http.Request) {
	query := "INSERT INTO MyTable SET Name=?"
	var employee Employee
	json.NewDecoder(r.Body).Decode(&employee)
	stmt, err := dbConnection.Prepare(query)
	if err != nil {
		log.Println("Could not insert new employee")
	}
	result, err := stmt.Exec(employee.Name)
	if err != nil {
		log.Println("Found a problem executing query: \"%s\"", query)
	}
	id, err := result.LastInsertId()
	fmt.Fprintf(w, "inserted id: %d", id)
}

func updateRecord(w http.ResponseWriter, r *http.Request) {
	var employee Employee
	json.NewDecoder(r.Body).Decode(&employee)
	stmt, err := dbConnection.Prepare("UPDATE myTable SET Name=? WHERE Id=?")
	if err != nil {
		log.Println("Failed to update employee")
	}
	result, err := stmt.Exec(employee.Name, employee.Id)
	rowsAffected, err := result.RowsAffected()
	fmt.Printf("Rows affected: %d", rowsAffected)
	getRecords(w, r)
}

func main() {
	router := mux.NewRouter()
	AddRoutes(router.PathPrefix("/employee").Subrouter())
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
	if err != nil {
		log.Fatal("The server could not be started. Error: ", err)
	}
}
