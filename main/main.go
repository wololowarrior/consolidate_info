package main

import (
	"log"
	"net/http"

	"accuknox/dao/datastore"
	"accuknox/service"
	"github.com/gorilla/mux"
)

func main() {
	handleRequests()
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	users := datastore.NewUsers()
	notes := datastore.NewNotes()
	userService := service.NewUserService(users)
	noteService := service.NewNotesService(notes, users)
	myRouter.HandleFunc("/signup", userService.Signup).Methods("POST")
	myRouter.HandleFunc("/login", userService.Login).Methods("POST")
	myRouter.HandleFunc("/logout", userService.Logout).Methods("POST")
	myRouter.HandleFunc("/notes", noteService.Create).Methods("POST")
	myRouter.HandleFunc("/notes", noteService.Delete).Methods("DELETE")
	myRouter.HandleFunc("/notes", noteService.List).Methods("GET")
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
