package main

import (
	"errors"
)

var message = ""
func MustStoreMessage(m string) {
	err := StoreMessage(m)
	if err != nil {
		panic("StoreMessage returned error")
	}
}
func StoreMessage(m string) error {
	if m == "" {
		return errors.New("no message")
	}

	message = m

	return nil
}

func main() {
	MustStoreMessage("Hello!")
}
