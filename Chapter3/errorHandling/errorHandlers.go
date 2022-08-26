package main

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

type NameNotFound struct {
	Code int
	Err  error
}

func (nameNotFoundError NameNotFound) Error() string {
	return nameNotFoundError.Err.Error()
}

type WrapperHandler func(w http.ResponseWriter, r *http.Request) error

func (wrapperHandler WrapperHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := wrapperHandler(w, r)
	if err != nil {
		switch e := err.(type) {
		case NameNotFound:
			log.Printf("HTTP %s - %d", e.Err, e.Code)
			http.Error(w, e.Err.Error(), e.Code)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

func getName(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	name := vars["name"]
	if strings.EqualFold(name, "foo") {
		fmt.Fprintf(w, "Your name is foo")
		return nil
	} else {
		return NameNotFound{Code: 500, Err: errors.New("name not found")}
	}
}

func main() {
	router := mux.NewRouter()
	router.Handle("/employee/get/{name}", WrapperHandler(getName)).Methods("GET")
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}
