package main

import "fmt"

var eatenCount int
var originalCount int

func main() {
	originalCount = 10
	eatenCount = 4

	fmt.Println("I started with", originalCount, "apples.")

	fmt.Println("Some jerk ate", eatenCount, "apples.")

	fmt.Println("There are", originalCount - eatenCount, "apples.")
}