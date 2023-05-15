package main

import (
	"encoding/json"
	"log"
	"net/http"
	"rest-vs-grpc/internal/constants"
)

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(constants.Response)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
