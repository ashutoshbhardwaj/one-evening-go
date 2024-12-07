package main

import (
	"fmt"
	"os"
)

func CheckFile(inp string) bool {
	_, err := os.ReadFile(inp)
	if err != nil {
		return false
	}
	return true
}
func main() {
	ok := CheckFile("input.csv")
	if ok {
		fmt.Println("File correctly read")
	} else {
		fmt.Println("Failed to read file")
	}
}
