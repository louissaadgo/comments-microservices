package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Port 4003
const address string = ":4003"

//Comment Structure
type comment struct {
	ID       string `json:"ID"`
	Content  string `json:"Content"`
	Verified bool   `json:"Verified"`
	Valid    bool   `json:"Valid"`
}

//Forwards events to all services
func eventBus(w http.ResponseWriter, r *http.Request) {
	newComment := comment{}
	err := json.NewDecoder(r.Body).Decode(&newComment)
	if err != nil {
		log.Fatalln(err)
	}
	bs, err := json.Marshal(newComment)
	if err != nil {
		log.Fatalln(err)
	}
	go moderation(bs)
	go query(bs)
}

//Forwards event to the moderation service
func moderation(bs []byte) {
	res, err := http.Post("http://localhost:4001/moderation", "application/json", bytes.NewBuffer(bs))
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
}

//Forwards event to the query service
func query(bs []byte) {
	resp, err := http.Post("http://localhost:4002/comments", "application/json", bytes.NewBuffer(bs))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
}

//Handles incoming requests
func handleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/bus", eventBus).Methods("POST")
	log.Fatal(http.ListenAndServe(address, r))
}

func main() {
	handleRequests()
}
