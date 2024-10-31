package main

import "net/http"

var PI_KEY string = "my32digitkey12345678901234567890"

func GetPiKey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(PI_KEY))
}
