package main

import "fmt"

var count int

func AllocateBuffer() *string {
	count++
	if count <= 3 {
		return new(string)
	}
	return nil

}

func main() {
	var buffers []*string

	for {
		b := AllocateBuffer()
		if b == nil {
			break
		}

		buffers = append(buffers, b)
	}

	fmt.Println("Allocated", len(buffers), "buffers")
}
