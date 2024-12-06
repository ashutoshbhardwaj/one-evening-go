package main

func main() {
	_ = Sum(1, 2, 3, 4, 5)
}

func Sum(n ...int) int {
	sum := 0
	for _, element := range n {
		sum = sum + element
	}
	return sum
}