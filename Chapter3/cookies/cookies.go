package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

var (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

var store *sessions.CookieStore

func init() {
	store = sessions.NewCookieStore([]byte("secret-key"))
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
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/home", home)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("Error: ", err)
		return
	}
}
