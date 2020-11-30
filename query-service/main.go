package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Port 4002
const address string = ":4002"

//Comment Structure
type comment struct {
	ID       string `json:"ID"`
	Content  string `json:"Content"`
	Verified bool   `json:"Verified"`
}

var newComment = comment{}

//Stores all comments
var comments = []comment{}

//Sends all approved comments
func getComments(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(comments)
}

//Receivec events from the event bus
func postComments(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode(&newComment)
	if newComment.Verified == true {
		comments = append(comments, newComment)
	}
}

//Handles incoming requests
func handleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/comments", getComments).Methods("GET")
	r.HandleFunc("/comments", postComments).Methods("POST")
	log.Fatal(http.ListenAndServe(address, r))
}

func main() {
	handleRequests()
}
