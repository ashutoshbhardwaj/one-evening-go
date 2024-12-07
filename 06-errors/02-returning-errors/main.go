package main

import (
	"fmt"
	"errors"
)

func main() {
	result, err := Divide(100, 50)
	fmt.Println("Result:", result, "Error:", err)
}

// func Divide(x float64, y float64) (float64, error){
// 	if y != 0 {
// 		return x/y, nil 
// 	} else {
// 		return 0, errors.New("Divs")
// 	}
// }

func Divide(x float64, y float64) (float64, error){
		if y == 0 {
			return 0, errors.New("Divs")
		} 
	
		return x / y, nil 
}