package main

import (
	"fmt"
	"os"
)

func one() {
	fmt.Println("1")
}
func two() {
	fmt.Println("2")
}
func three() {
	fmt.Println("3")
}

func sample() {
	fmt.Println("Begin")

	// STACK : 3 , 2 , 1
	defer one()
	defer two()
	defer three()

	fmt.Println("End")
}

func example() {
	file, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	buffer := make([]byte, 0, 30)
	bytes, err := file.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%c\n", bytes)
}

func main() {
	sample()
}