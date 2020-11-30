package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const address string = ":4001"

//Comment Structure
type comment struct {
	ID       string `json:"ID"`
	Content  string `json:"Content"`
	Verified bool   `json:"Verified"`
}

var newComment = comment{}

func moderation(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode(&newComment)
	newComment.Verified = true
	go check()
}

func check() {
	bs, err := json.Marshal(newComment)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.Post("http://localhost:4003/bus", "application/json", bytes.NewBuffer(bs))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	ioutil.ReadAll(resp.Body)
}

func handleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/moderation", moderation).Methods("POST")
	log.Fatal(http.ListenAndServe(address, r))
}

func main() {
	handleRequests()
}
