package main

import "fmt"

func main() {
	result := Sum(1, 2, 3, 4, 5)
	fmt.Println(result)
}

func Sum(first int, second int, third int, fourth int, fifth int) int {
	return first + second + third + fourth + fifth
}