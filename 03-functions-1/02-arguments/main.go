package main

import "fmt"

func main() {
	fmt.Println(Greet("Alice"))
	fmt.Println(Greet("Bob"))
}

func Greet(name string) { 
	return "Hello, " + name 
}
