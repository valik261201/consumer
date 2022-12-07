package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var order Order

func postOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var order Order

	_ = json.NewDecoder(r.Body).Decode(&order)

	// add orders to the end of the queue
	ordersConsumer.Enqueue(order)

	json.NewEncoder(w).Encode(&order)

	fmt.Print("\nConsumer recieved the data:\n", order)

	//go performPostRequest(order)
}

func main() {

	router := mux.NewRouter()

	go sendOrder(order)

	router.HandleFunc("/consumer", postOrder).Methods("POST")
	http.ListenAndServe(":5050", router)
}
