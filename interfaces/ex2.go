package main

import "fmt"

// Approximation ~
type Integer32 interface {
	~int32 | ~uint32
}

func SumNumbers[T Integer32](arr []T) T {
	var sum T
	for _, integer := range arr {
		sum += integer
	}
	return sum
}

func main() {
	nums := []int32{1, 2, 3}
	nums2 := []uint32{1, 2, 3}
	total := SumNumbers(nums)
	total2 := SumNumbers(nums2)

	fmt.Println(total, total2)
}