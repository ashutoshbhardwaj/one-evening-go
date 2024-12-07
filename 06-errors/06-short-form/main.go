package main

import (
	"os"
	"fmt"

)


func main() {
	if _, err := os.Stat("analysis.xlsx"); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ok")
	
}
