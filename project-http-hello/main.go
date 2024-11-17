package main

import (
	"fmt"
	"net/http"
	"log"
)
func main() {
	http.HandleFunc("/hello", handleEnpoint)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
// This function will get 2 parameters as context - 
// 1. http.ResponseWriter 
// 2. *http.Request
func handleEnpoint (w http.ResponseWriter, r *http.Request) {
	// do what you want to do 
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, "Hello, ", name)
}