package main

import (
	"fmt"
	"common/constraints"
)

type MyArray[T constraints.Ordered] struct {
	inner []T
}

func (m *MyArray[T]) Max() T {
	max := m.inner[0]
	for i := 0 ; i < len(m.inner) ; i++ {
		if m.inner[i] > max {
			max = m.inner[i]
		}
	}
	return max
}

func main() {
	a := MyArray[int]{inner: []int{5, 4, 19, 13, 55, 0}}
	fmt.Println(arr.Max())
}