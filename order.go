package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Order struct {
	Id         int   `json:"id"`
	Items      []int `json:"items"`
	Priority   int   `json:"priority"`
	MaxWait    int   `json:"max-wait"`
	PickUpTime int   `json:"pick-up-time"`
}

func sendOrder(order Order) {
	for {
		for !ordersAggregator.isEmpty() {
			go performPostRequest(order)
			time.Sleep(time.Second * 1)
		}
	}
}

func performPostRequest(order Order) {
	if !ordersAggregator.isEmpty() {
		const myUrl = "http://localhost:8080/aggregator"

		// return the first order form the queue
		order := ordersAggregator.Dequeue()

		var requestBody, _ = json.Marshal(order)
		time.Sleep(time.Second * 3)

		fmt.Printf("\nData was sent to the Aggregtor\n")
		response, err := http.Post(myUrl, "application/json", bytes.NewBuffer(requestBody))

		if err != nil {
			panic(err)
		}

		defer response.Body.Close()
		time.Sleep(time.Second * 1)
	}
}
