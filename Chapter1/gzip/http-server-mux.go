package main

import (
	"github.com/gorilla/handlers"
	"io"
	"log"
	"net/http"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = ":8080"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorld)
	err := http.ListenAndServe(CONN_HOST+CONN_PORT, handlers.CompressHandler(mux))
	if err != nil {
		log.Fatal("error starting http server: ", err)
		return
	}
}
