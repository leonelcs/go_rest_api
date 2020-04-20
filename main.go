package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	//preparing routes to be called
	router.HandleFunc("/", homeLink).Methods("GET")
	router.HandleFunc("/events", getAllEvents).Methods("GET")
	router.HandleFunc("/events/{id}", getEventByIDHandler).Methods("GET")
	//start the server listening to the port
	log.Fatal(http.ListenAndServe(":8080", router))
}
