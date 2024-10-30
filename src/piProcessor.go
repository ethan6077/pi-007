package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/ethan6077/pi-007/src/models"
)

func processPi(w http.ResponseWriter, r *http.Request) {
	println("Processing Pi")

	missionId, uuidErr := uuid.NewRandom()
	if uuidErr != nil {
		fmt.Printf("Failed to generate UUID: %v\n", uuidErr)
		panic(uuidErr)
	}

	form := models.Form{
		Name:            "Vesper Lynd",
		Phone:           "0482663672",
		Email:           "vesper.lynd@gmail.com",
		Content:         "I am very interested in this property. When can I inspect it?",
		EnquiryCategory: "Inspection",
		ListingId:       "listing77777",
		AgentId:         "007",
		AgencyID:        "MI6",
	}

	missionCard := models.MissionCard{
		Status:         "Stealth",
		PiDataDetected: true,
		MissionId:      missionId,
		MissionKey:     "007",
		Body:           form,
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(missionCard)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
