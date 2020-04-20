package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Festa da Uva de Vinhedo",
		Description: "Venham tomar o melhor do vinho",
	},
	{
		ID:          "2",
		Title:       "Festa da noz de campinas",
		Description: "Venham todos para essa festa comer noz",
	},
}

//home uri
func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func getEventByID(id string) (*event, error) {
	for _, firstEvent := range events {
		if firstEvent.ID == id {
			return &firstEvent, nil
		}
	}
	return nil, errors.New("Event not found")
}

func getEventByIDHandler(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	res, err := getEventByID(eventID)

	if err != nil {
		log.Fatal("Event not found")
	}
	json.NewEncoder(w).Encode(res)
}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)
}
