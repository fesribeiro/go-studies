package main

import (
	"basic-crud/server"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func main() { 

	router := mux.NewRouter()

	router.HandleFunc("/users", server.CreateUser).Methods(http.MethodPost)

	fmt.Println("Starting the server on port 5000...")
	log.Fatal(http.ListenAndServe(":5000", router))
}
