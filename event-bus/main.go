package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Port 4003
const address string = ":4003"

//Forwards events to all services
func eventBus(w http.ResponseWriter, r *http.Request) {
	go moderation(r.Body)
	go query(r.Body)
}

//Forwards event to the moderation service
func moderation(body io.ReadCloser) {
	res, err := http.Post("http://localhost:4001/moderation", "application/json", body)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
}

//Forwards event to the query service
func query(body io.ReadCloser) {
	resp, err := http.Post("http://localhost:4002/comments", "application/json", body)
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
