package models

type RetrieveReq struct {
	Key  string `json:"key"`
	Body Form   `json:"body"`
}
