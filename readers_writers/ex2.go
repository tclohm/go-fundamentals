package main

import (
	"strings"
	"bufio"
	"io"
	"fmt"
)

func main() {
	source := strings.NewReader("SAMPLE")
	buffered := bufio.NewReader(source)
	// buffered.ReadBytes
	newString, err := buffered.ReadString('\n')
	if err == io.EOF {
		fmt.Println(newString)
	} else {
		fmt.Println("something went wrong...")
	}
}