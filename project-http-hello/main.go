package main

import (
	"fmt"
	"net/http"
	"log"
)
var calls []string

var stats = map[string]int{}
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
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	calls = append(calls, name)
	stats[name]++
	fmt.Printf("calls: %#v\n", calls)
	fmt.Printf("stats: %#v\n", stats)
	fmt.Fprint(w, "Hello, ", name)

}