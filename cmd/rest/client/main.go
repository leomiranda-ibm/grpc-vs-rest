package main

import (
	"encoding/json"
	"log"
	"net/http"
	"rest-vs-grpc/internal/cli"
	"rest-vs-grpc/internal/constants"
	"rest-vs-grpc/internal/entities"
	"time"
)

func main() {
	log.Println("---- REST CLIENT TEST ----")
	log.Printf("There are %v users in response\n", len(constants.Response.Users))

	qtyRequests, err := cli.HandleParams()
	if err != nil {
		log.Fatal(err)
	}

	result, err := manyCalls(qtyRequests)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("average of %v requests is: %v/request", qtyRequests, result)
}

func manyCalls(numCalls int) (time.Duration, error) {
	var callsSum time.Duration
	for i := 0; i < numCalls; i++ {
		startCall := time.Now()
		_, err := callAPI()
		if err != nil {
			return 0, err
		}
		requestTime := time.Since(startCall)
		callsSum = callsSum + requestTime
	}

	return callsSum / time.Duration(numCalls), nil
}

func callAPI() (result *entities.Response, err error) {
	resp, err := http.Get("http://localhost:8080/users")
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
