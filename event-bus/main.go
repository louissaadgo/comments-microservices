package main

import (
	"encoding/json"
	"fmt"
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
}

//Routes Incoming Requests
func eventBus(w http.ResponseWriter, r *http.Request) {
	newComment := comment{}
	json.NewDecoder(r.Body).Decode(&newComment)
	fmt.Println(newComment)
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
