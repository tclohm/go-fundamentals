package main

import "fmt"

func sum(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func main() {
	a := []int{1,2,3,4}
	b := []int{5,6,7}

	all := append(a, b...)

	fmt.Println(sum(all...))

	fmt.Println(sum(2, 4, 6, 8, 10, 12))
}