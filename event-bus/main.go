package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Port 4003
const address string = ":4003"

//Forwards events to all services
func eventBus(w http.ResponseWriter, r *http.Request) {
	//Forwards to the query service
	resp, err := http.Post("http://localhost:4002/comments", "application/json", r.Body)
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
