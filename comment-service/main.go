package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

//Port 4000
const address string = ":4000"

//Comment Structure
type comment struct {
	ID       string `json:"ID"`
	Content  string `json:"Content"`
	Verified bool   `json:"Verified"`
}

//Generates a new ID
func getID() string {
	return uuid.New().String()
}

//Creates new comments
func postComment(w http.ResponseWriter, r *http.Request) {
	newComment := comment{
		ID:       getID(),
		Verified: false,
	}
	//Decodes the request body into newComment
	err := json.NewDecoder(r.Body).Decode(&newComment)
	if err != nil {
		log.Fatalln("Error decoding rquest body: ", err)
	}
	fmt.Fprintf(w, "Comment Received! \n %v", newComment)
	//Sends an event to the Event-Bus
	bs, err := json.Marshal(newComment)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.Post("http://localhost:4003/bus", "application/json", bytes.NewBuffer(bs))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
}

//Handles incoming requests
func handleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/comment", postComment).Methods("POST")
	log.Fatal(http.ListenAndServe(address, r))
}

func main() {
	handleRequests()
}
