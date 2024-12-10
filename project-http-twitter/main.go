// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"log"
// 	"encoding/json"
// 	"io"
// )

// type userPayload struct {
// 	Message  string `json:"message"`
// 	Location   string    `json:"location"`
// }

// func main() {
// 	http.HandleFunc("/tweets", tweetHandle)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// func tweetHandle(w http.ResponseWriter, r *http.Request) {
// 	body, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		log.Println("Failed to read body: %s", err)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	defer r.Body.Close() 

// 	u := userPayload{}

// 	if err := json.Unmarshal(body, &u); err != nil {
// 		log.Println("Failed to unmarshal payload: %s", err)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	if u.Message == "" {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	fmt.Printf("Tweet: `%s` from %s\n", u.Message, u.Location)
// }

//
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/tweets", addTweet)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type tweet struct {
	Message  string `json:"message"`
	Location string `json:"location"`
}

type response struct {
	ID int `json:"ID"`
}
var count = 0

func addTweet(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Failed to read body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	tw := tweet{}

	if err := json.Unmarshal(body, &tw); err != nil {
		log.Printf("Failed to unmarshal payload: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if tw.Message == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Printf("Tweet: `%s` from %s\n", tw.Message, tw.Location)
	count ++ 
	resp := response{
		ID: count,
	}
	responsePayload, err := json.Marshal(resp)
	if err != nil {
		log.Println("Failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(responsePayload)
}
