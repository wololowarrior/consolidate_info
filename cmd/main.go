package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"accuknox/dao/datastore"
	"accuknox/service"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	db := connectDatastore()
	handleRequests(db)
}

func handleRequests(db *sql.DB) {
	myRouter := mux.NewRouter().StrictSlash(true)
	users := datastore.NewUsers(db)
	userService := service.NewUserService(users)
	myRouter.HandleFunc("/identify", userService.Signup).Methods("POST")

	fmt.Println("Starting server..")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func connectDatastore() *sql.DB {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DATABASE")
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	// Open a connection to the PostgreSQL database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the PostgreSQL database!")
	return db
}
