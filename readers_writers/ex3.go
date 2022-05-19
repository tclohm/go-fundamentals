package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0, 5)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	}

	fmt.Printf("Line count: %v\n", len(lines))
	for _, line := range lines {
		fmt.Printf("Line: %v\n", line)
	}
}