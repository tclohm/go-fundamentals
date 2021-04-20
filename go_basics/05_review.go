package main

import "fmt"

// func average(numbers[5]float64) float64 {
// 	total := 0.0
// 	for _, number := range numbers {
// 		total += number
// 	}
// 	return total / float64(len(numbers))
// }
// var animals map[string]string = map[string]string{
// 	"fido": "dog",
// 	"jerry": "cat",
// }

// func doesItExist(pet string) bool {
// 	_, ok := animals[pet]
// 	return ok
// }

var initialGroceries = []string{"beans", "lemons", "chicken", "fruit"}

func appendToList(items ...string) []string {
	foods := initialGroceries
	for _, item := range items {
		foods = append(foods, item)
	}

	return foods
}

func main() {
	// scores := [5]float64{1.0, 2.0, 3.3, 4.2, 5.9}
	// fmt.Println(average(scores))
	// fido := "fido"
	// fmt.Println(doesItExist(fido))

	fmt.Println(appendToList("cherries", "water"))
}
