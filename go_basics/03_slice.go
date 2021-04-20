package main

import "fmt"

func main() {
	//var myArray [5]int
	// type, initialspace, maxspace
	// var mySlice []int = make([]int, 5, 10)

	// myArray[0] = 1
	// mySlice[0] = 1
	// // len(myArray), cap(myArray)
	// fmt.Println(myArray)
	// fmt.Println(mySlice)

	fruitArray := [5]string{"banana", "pear", "apple", "kumquat", "peach"}

	splicedFruit := fruitArray[1:3]
	fmt.Println(splicedFruit)
	fmt.Println(len(splicedFruit))
	fmt.Println(cap(splicedFruit))
}