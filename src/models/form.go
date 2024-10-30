package main

type Form struct {
	Name            string `json:"name"`
	Phone           string `json:"phone"`
	Email           string `json:"email"`
	Content         string `json:"content"`
	EnquiryCategory string `json:"enquiryCategory"`
	ListingId       string `json:"listingId"`
	AgentId         string `json:"agentId"`
	AgencyID        string `json:"agencyId"`
}
