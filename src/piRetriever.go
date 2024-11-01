package main

import (
	"encoding/json"
	"net/http"

	"github.com/ethan6077/pi-007/src/models"
)

func RetrievePi(w http.ResponseWriter, r *http.Request) {
	println("Retrieving Pi")

	var retrieveReq models.RetrieveReq
	payloadError := json.NewDecoder(r.Body).Decode(&retrieveReq)
	if payloadError != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		panic(payloadError)
	}

	decryptedForm := DecryptPiData(retrieveReq.Body, retrieveReq.Key)

	w.Header().Set("Content-Type", "application/json")

	jsonEncoderErr := json.NewEncoder(w).Encode(decryptedForm)

	if jsonEncoderErr != nil {
		http.Error(w, jsonEncoderErr.Error(), http.StatusInternalServerError)
	}
}
