package main

import (
	"os"
	"strings"
	"fmt"
	"log"
)

func main() {
	files := listFiles("testdata")
	fmt.Println(strings.Join(files," "))

}

func listFiles(dirname string) []string {
	var dirs []string

	files, err := os.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		dirs = append(dirs, f.Name())
	}

	return dirs
}
