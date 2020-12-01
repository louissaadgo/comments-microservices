package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

//Port 4001
const address string = ":4001"

//Comment Structure
type comment struct {
	ID       string `json:"ID"`
	Content  string `json:"Content"`
	Verified bool   `json:"Verified"`
	Valid    bool   `json:"Valid"`
}

//Checks if the comment is valid
func moderation(w http.ResponseWriter, r *http.Request) {
	newComment := comment{}
	json.NewDecoder(r.Body).Decode(&newComment)
	if newComment.Verified == false {
		newComment.Verified = true
		if strings.Contains(newComment.Content, "ugly") == false {
			newComment.Valid = true
		}
		go sendEvent(newComment)
		return
	}
	return
}

//Sends an event to the event bus
func sendEvent(newComment comment) {
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
	r.HandleFunc("/moderation", moderation).Methods("POST")
	log.Fatal(http.ListenAndServe(address, r))
}

func main() {
	handleRequests()
}
