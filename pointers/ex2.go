package main

import "fmt"

type Coordinates struct {
	X, Y float64
}

var c = Coordinates{X: 10, Y: 20}

func main() {
	coordinatesMemoryAddress := c
	coordinatesMemoryAddress.X = 200
	fmt.Println(coordinatesMemoryAddress)
}