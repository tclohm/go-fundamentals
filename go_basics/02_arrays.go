package main

import "fmt"

func main() {
	// var scores [5]float64 = [5]float64{9, 1.5, 4.5, 7, 8}
	// scores := [5]float64{9, 1.5, 4.5, 7, 8}
	scores := [...]float64{9, 1.5, 4.5, 7, 8}

	for _, value := range scores {
		fmt.Println(value)
	}

	// fmt.Println(scores)
}

