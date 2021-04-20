package main

import (
	"fmt"
	"net/http"
	"log"
)

func home(writer http.ResponseWriter, req *http.Request) {
	fmt.Println("Home!")
}

func main() {
	http.HandleFunc("/", home)
	fmt.Println("Server is running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}