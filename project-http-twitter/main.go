package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	s := server{
		TweetRepository: &TweetMemoryRepositry{},
	}
	http.HandleFunc("/tweets", s.addTweet)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type tweet struct {
	Message  string `json:"message"`
	Location string `json:"location"`
}

type response struct {
	ID int `json:"ID"`
}

type TweetRepository interface {
	AddTweet(tw tweet) (int) 
}

type TweetMemoryRepositry struct {
	Tweets []tweet
}

func (tmr *TweetMemoryRepositry) AddTweet(tw tweet) (int) {
	tmr.Tweets = append(tmr.Tweets, tw)
	return len(tmr.Tweets)

}
// var count = 0

type server struct {
	TweetRepository TweetRepository
}

func (s *server) addTweet(w http.ResponseWriter, r *http.Request) {
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
	// count ++ 
	id := s.TweetRepository.AddTweet(tw)
	
	resp := response{
		ID: id,
	}
	responsePayload, err := json.Marshal(resp)
	if err != nil {
		log.Println("Failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(responsePayload)
}
