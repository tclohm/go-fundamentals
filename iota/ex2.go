package main

import "fmt"

const (
	Add = iota
	Subtract
	Multiple
	Divide
)

type Operation int

func (op Operation) calc(lhs, rhs int) int {
	switch op {
	case Add:
		return lhs + rhs
	case Subtract:
		return lhs - rhs
	case Multiple:
		return lhs * rhs
	case Divide:
		return lhs / rhs
	}
	panic("unhandled operation")
}

func main() {
	add := Operation(Add)
	sub := Operation(Subtract)
	mul := Operation(Multiple)
	div := Operation(Divide)

	fmt.Println(add.calc(2,2))

	fmt.Println(sub.calc(10, 3))

	fmt.Println(mul.calc(3, 3))

	fmt.Println(div.calc(100, 2))
}