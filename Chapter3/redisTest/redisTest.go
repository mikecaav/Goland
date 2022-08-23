package main

import (
	"fmt"
	"gopkg.in/boj/redistore.v1"
	redisStore "gopkg.in/boj/redistore.v1"
	"log"
	"net/http"
)

var (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

var store *redistore.RediStore
var err error

func init() {
	store, err = redisStore.NewRediStore(10, "tcp", ":6379", "",
		[]byte("secret-key"))
	if err != nil {
		log.Fatal("error getting redis store : ", err)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	session.Values["authenticated"] = true
	session.Save(r, w)
	fmt.Fprintf(w, "You have succesfully logged in")
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	session.Values["authenticated"] = false
	session.Save(r, w)
	fmt.Fprintf(w, "You have succesfully logged out")
}

func home(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	var authenticated = session.Values["authenticated"]
	if authenticated != nil {
		isAuthenticated := session.Values["authenticated"].(bool)
		if !isAuthenticated {
			http.Error(w, "You're unauthorized to see this page", http.StatusForbidden)
			return
		}
		fmt.Fprintf(w, "Home page, you've succesfully been authorized, congratulation!")
	} else {
		http.Error(w, "You're unauthorized to see this page", http.StatusForbidden)
		return
	}
}

func main() {
	defer store.Close()
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/home", home)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("Error: ", err)
		return
	}
}
