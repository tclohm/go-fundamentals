package main

import (
	"errors"
	"fmt"
)

type DivError struct {
	a, b int
}

func (d *DivError) Error() string {
	return fmt.Sprintf("Cannot divide by zero: %d / %d", d.a, d.b)
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivError{a, b}
	} else {
		return a / b, nil
	}
}

func divide(lhs, rhs int) (int, error) {
	if rhs == 0 {
		return 0, errors.New("cannot divide by zero")
	} else {
		return rhs / lhs, nil
	}
}

func main() {
	answer1, err := div(9,0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("answer is", answer1)
}