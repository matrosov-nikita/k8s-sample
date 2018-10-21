package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// For private use only!!!
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if request.Method == "OPTIONS" {
			return
		}

		var b struct {
			Name string `json:"name"`
		}

		if err := json.NewDecoder(request.Body).Decode(&b); err != nil {
			http.Error(writer, "could not parse request body", http.StatusBadRequest)
			return
		}

		defer func() {
			log.Println("name:", b.Name)
		}()


		resp, err := json.Marshal(b)
		if err != nil {
			http.Error(writer, "could not build response", http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.Write(resp)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}