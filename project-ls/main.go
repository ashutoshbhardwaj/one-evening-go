package main

import (
	"os"
	"fmt"
	"log"
)

func main() {
	files := listFiles("testdata")
	for _, file := range files {
		fmt.Println(file)
	}

}

func listFiles(dirname string) []string {
	dirs := []string{}

	files, err := os.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		dirs = append(dirs, f.Name())
	}
	// fmt.Println("--")
	// fmt.Println(dirs)
	// fmt.Println("--")
	return dirs
}
