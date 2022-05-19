package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"io"
	"strconv"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	sum := 0

	for {
		input, inputError := r.ReadString(' ')
		n := strings.TrimSpace(input)
		
		if n == "" {
			continue
		}

		number, err := strconv.Atoi(n)
		if err != nil {
			fmt.Println(err)
		} else {
			sum += number
		}
		// end of file : ctrl + d
		if inputError == io.EOF {
			break
		}

		if inputError != nil {
			fmt.Println("Error reading Stdin:", inputError)
		}
	}
	fmt.Printf("sum: %v\n", sum)
}