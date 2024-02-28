package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Palaszontko/go-luhn-credit-validator/internal/luhn"
)

func validateCardHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is accepted!", http.StatusMethodNotAllowed)
		return
	}

	var req validationRequest
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	isValid := luhn.ValidateCard(req.CardNumber)
	cardNetwork := luhn.CardNetwork(req.CardNumber)

	response := responseStruct{
		IsValid:    isValid,
		CardNetwork: cardNetwork,
	}

	responseJSON, err := json.Marshal(response)

	if err != nil {
		http.Error(w, "Error while processing request", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

type responseStruct struct {
	IsValid    bool   `json:"isValid"`
	CardNetwork string `json:"cardNetwork"`
}

type validationRequest struct {
	CardNumber string `json:"cardNumber"`
}

func main() {

	http.HandleFunc("/validate", validateCardHandler)
	log.Println("Server is listening on port 8080")
	// Uruchom serwer na porcie 8080 i nasłuchuj na żądania
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
