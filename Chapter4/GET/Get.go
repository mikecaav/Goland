package main

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

type WrapperHandler func(w http.ResponseWriter, r *http.Request) error

func (wrapperHanlder WrapperHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := wrapperHanlder(w, r)
	if err != nil {
		switch e := err.(type) {
		case EmployeeNotFound:
			log.Printf("HTTP %s - %d", e.Err, e.Code)
			http.Error(w, e.Err.Error(), e.Code)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler http.Handler
}

type Routes []Route

var routes = Routes{
	Route{
		Name:    "getEmployees",
		Method:  "GET",
		Pattern: "/employees",
		Handler: WrapperHandler(getEmployees),
	},
	Route{
		Name:    "getEmployee",
		Method:  "GET",
		Pattern: "/employee/{id}",
		Handler: WrapperHandler(getEmployee),
	},
}

type Employee struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Employees []Employee

var employees Employees

type EmployeeNotFound struct {
	Code int
	Err  error
}

func (employeeNotFound EmployeeNotFound) Error() string {
	return employeeNotFound.Err.Error()
}

func init() {
	employees = Employees{
		Employee{Id: "1", FirstName: "Foo", LastName: "Bar"},
		Employee{Id: "2", FirstName: "Baz", LastName: "Qux"},
	}
}

func getEmployees(w http.ResponseWriter, r *http.Request) error {
	json.NewEncoder(w).Encode(employees)
	return nil
}

func getEmployee(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)

	for _, employee := range employees {
		if employee.Id == vars["id"] {
			json.NewEncoder(w).Encode(employee)
			return nil
		}
	}

	return EmployeeNotFound{500, errors.New("employee not found")}
}

func AddRoutes(router *mux.Router) *mux.Router {
	for _, route := range routes {
		router.Methods(route.Method).Handler(route.Handler).Name(route.Name).Path(route.Pattern)
	}
	return router
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router = AddRoutes(router)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
	if err != nil {
		log.Fatal("error starting http server :: ", err)
		return
	}
}
