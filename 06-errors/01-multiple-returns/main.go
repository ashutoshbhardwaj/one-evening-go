package main

import "fmt"

func main() {
	x := "World!"
	y := "Hello,"

	x, y = Swap(x, y)

	fmt.Println(x, y)
}

func Swap(x string, y string) (string, string) {
	return y, x
}