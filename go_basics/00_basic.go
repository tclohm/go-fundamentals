package main

import "fmt"

// once you declare it you better use it!
func main() {
	sentence := "Hello World, this is my first sentence in go!"

	for index, value := range sentence {
		if index % 2 == 0 {
			fmt.Println(string(value))
		}
	}
}