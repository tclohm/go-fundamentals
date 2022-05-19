package main

import "fmt"

func add(lhs, rhs int) int {
	return lhs + rhs
}

type OperatorFunc func(lhs, rhs int) int

func compute(
	lhs, rhs int,
	op OperatorFunc) int {
		fmt.Printf("Running a computation with %v & %v\n", lhs, rhs)
		return op(lhs, rhs)
}

func main() {
	fmt.Println("2 + 2 =", compute(2, 2, add))

	fmt.Println("20 - 15 =", compute(20, 15, func(lhs, rhs int) int {
		return lhs - rhs
	}))

	mul := func(lhs, rhs int) int {
		fmt.Printf("multipling %v * %v = ", lhs, rhs)
		return lhs * rhs
	}

	fmt.Println(compute(3, 3, mul))
}