package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Rand    int    `json:"rand"`
}

// API handler for reversing a string
func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var data Data
		err := decoder.Decode(&data)
		if err != nil {
			http.Error(w, "Bad Request",
				http.StatusBadRequest)
		}

		// Call to reverse method
		if err := data.reverse(); err != nil {
			http.Error(w, "Some error occurred.",
				http.StatusInternalServerError)
		}

		resp := &Response{data.Message, rand.Int()}

		res, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, "Some error occurred.",
				http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
