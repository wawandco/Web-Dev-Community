package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Server listening at port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
