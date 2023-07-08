package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {
	r := mux.NewRouter()

	r.HandleFunc("/", HomePage)
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", GetUserByID).Methods("GET")
	r.HandleFunc("/users/{id}", UpdateUserByID).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUserByID).Methods("DELETE")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
