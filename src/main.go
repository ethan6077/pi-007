package main

import (
	"net/http"
)

const PORT_NUM string = ":8080"

func main() {
	println("Hi, I am Bond, James Bond.")

	println("Listening on port", PORT_NUM)

	http.HandleFunc("/pi-007/getMyKey", GetPiKey)

	http.HandleFunc("/pi-007/process", ProcessPi)

	http.HandleFunc("/pi-007/retrieve", RetrievePi)

	err := http.ListenAndServe(PORT_NUM, nil)
	if err != nil {
		panic(err)
	}

}
