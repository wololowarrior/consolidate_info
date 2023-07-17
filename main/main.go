package main

import (
	"fmt"
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
	myRouter := mux.NewRouter().StrictSlash(true)
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
	fmt.Println("Starting server..")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
