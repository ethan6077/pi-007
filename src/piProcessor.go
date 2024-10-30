package main

import (
	"net/http"
	// "github.com/ethan6077/pi-007/src/models"
)

func processPi(w http.ResponseWriter, r *http.Request) {
	println("Processing Pi")

	// form := models.Form{
	// 	Name:            "Ethan Zhang",
	// 	Phone:           "1234567890",
	// 	Email:           "ethan.zhang@gmail.com",
	// 	Content:         "I am very interested in this property. When can I inspect it?",
	// 	EnquiryCategory: "Inspection",
	// 	ListingId:       "6666666",
	// 	AgentId:         "007",
	// 	AgencyID:        "MI6",
	// }

	w.Header().Set("Content-Type", "application/json")

	// err := json.NewEncoder(w).Encode(form)

	w.Write([]byte("Hello, World!"))

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

}
