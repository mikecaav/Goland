package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"log"
	"net/http"
	"time"
)

var cacheStore *cache.Cache

func init() {
	cacheStore = cache.New(5*time.Minute, 10*time.Minute)
}

func readCache(w http.ResponseWriter, r *http.Request) {
	foo, found := cacheStore.Get("foo")
	if found {
		log.Println("Key found :: ", foo.(string))
		fmt.Fprintf(w, "key found in Cache "+foo.(string))
	} else {
		log.Println("Key not found we")
		fmt.Fprintf(w, "key not found in Cache")
	}
}

func createCache(w http.ResponseWriter, r *http.Request) {
	cacheStore.Set("foo", "bar", cache.DefaultExpiration)
	fmt.Fprintf(w, "Created foo bar")
}

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

func main() {
	http.HandleFunc("/", readCache)
	http.HandleFunc("/create", createCache)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}
