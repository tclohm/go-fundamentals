package main

import (
	"strings"
	"io"
	"fmt"
)

func main() {
	reader := strings.NewReader("Sample")

	var newString strings.Builder
	buffer := make([]byte, 4)
	for {
		numBytes, err := reader.Read(buffer)
		chunk := buffer[:numBytes]
		newString.Write(chunk)
		fmt.Printf("Read %v bytes: %c\n", numBytes, chunk)
		fmt.Printf("%v\n", newString.String())
		if err == io.EOF {
			break
		}
	}
	fmt.Printf("DONE: %v\n", newString.String())
}