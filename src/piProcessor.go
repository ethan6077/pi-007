package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/ethan6077/pi-007/src/models"
)

func ProcessPi(w http.ResponseWriter, r *http.Request) {
	println("Processing Pi")

	missionId, uuidErr := uuid.NewRandom()
	if uuidErr != nil {
		fmt.Printf("Failed to generate UUID: %v\n", uuidErr)
		panic(uuidErr)
	}

	var payload models.Form
	payloadError := json.NewDecoder(r.Body).Decode(&payload)
	if payloadError != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		panic(payloadError)
	}

	encryptedForm := EncryptPiData(payload)

	missionCard := models.MissionCard{
		Status:         "Stealth",
		PiDataDetected: true,
		MissionId:      missionId,
		Body:           encryptedForm,
	}

	w.Header().Set("Content-Type", "application/json")

	jsonEncoderErr := json.NewEncoder(w).Encode(missionCard)

	if jsonEncoderErr != nil {
		http.Error(w, jsonEncoderErr.Error(), http.StatusInternalServerError)
	}

}
