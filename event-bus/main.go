package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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

//Forwards events to all services
func eventBus(w http.ResponseWriter, r *http.Request) {
	newComment := comment{}
	json.NewDecoder(r.Body).Decode(&newComment)
	bs, _ := json.Marshal(newComment)
	if newComment.Verified == false {
		go mod(bs)
	} else {
		go query(bs)
	}
}
func mod(bs []byte) {
	//Forwards to the moderation service
	res, err := http.Post("http://localhost:4001/moderation", "application/json", bytes.NewBuffer(bs))
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	ioutil.ReadAll(res.Body)
}
func query(bs []byte) {
	//Forwards to the query service
	resp, err := http.Post("http://localhost:4002/comments", "application/json", bytes.NewBuffer(bs))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	ioutil.ReadAll(resp.Body)
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
