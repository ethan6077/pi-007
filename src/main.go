package main

import "net/http"

const PORT_NUM string = ":8080"

func main() {
	println("Hi, I am Bond, James Bond.")

	println("Listening on port", PORT_NUM)

	http.HandleFunc("/pi/process", processPi)

	err := http.ListenAndServe(PORT_NUM, nil)
	if err != nil {
		panic(err)
	}

}
