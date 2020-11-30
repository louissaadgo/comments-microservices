package main

import (
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

//Stores all comments
var comments = []comment{}

//Sends all approved comments
func getComments(w http.ResponseWriter, r *http.Request) {

}

//Handles incoming requests
func handleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/comments", getComments).Methods("GET")
	log.Fatal(http.ListenAndServe(address, r))
}

func main() {
	handleRequests()
}
