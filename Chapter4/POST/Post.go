package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

type Employee struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Employees []Employee

type Route struct {
	Path        string
	Name        string
	HandlerFunc http.HandlerFunc
	Method      string
}

type Routes []Route

var employees Employees
var routes Routes

func init() {
	employees = Employees{
		Employee{1, "Miguel", "Caballero"},
		Employee{2, "Benjamin", "Caballero"},
		Employee{3, "Zaira", "Avendaño"},
	}
	routes = Routes{
		Route{Name: "getEmployees", HandlerFunc: getEmployees, Path: "/employees", Method: "GET"},
		Route{Name: "addEmployee", HandlerFunc: addEmployee, Path: "/add-employee", Method: "POST"},
	}
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(employees)
}

func addEmployee(w http.ResponseWriter, r *http.Request) {
	employee := Employee{}
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		log.Print("error occurred while decoding employee data :: ", err)
		return
	}
	employees = append(employees, employee)
	json.NewEncoder(w).Encode(employees)
}

func AddRoute(muxRouter *mux.Router) *mux.Router {
	for _, route := range routes {
		muxRouter.Name(route.Name).Path(route.Path).Methods(route.Method).Handler(route.HandlerFunc)
	}

	return muxRouter
}

func main() {
	var muxRouter = mux.NewRouter().StrictSlash(true)
	muxRouter = AddRoute(muxRouter)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, muxRouter)
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}
