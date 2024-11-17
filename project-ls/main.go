package main

import (
	"os"
	"strings"
	"fmt"
)

func main() {
	files := listFiles("testdata")
	fmt.Println(strings.Join(files," "))

}

func listFiles(dirname string) []string {
	var dirs []string

	files, _ := os.ReadDir(dirname)
	for _, f := range files {
		dirs = append(dirs, f.Name())
	}

	return dirs
}
