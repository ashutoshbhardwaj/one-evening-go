package main

import "fmt"

func Alphabet(length int) []string {
	var s []string
	for i:=0; i < length; i++ {
		s = append(s,characterByIndex(i))
	}
	return s
}

func main() {
	alphabet := Alphabet(26)
	fmt.Println(alphabet)
}

func characterByIndex(i int) string {
	return string(rune('a' + i))
}
