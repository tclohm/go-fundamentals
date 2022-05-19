package main

import (
	"constraints"
	"fmt"
)

type Distance int32
type Velocity float64

type Number interface {
	constrai snts.Float | constraints.Integer
}

func clamp[T Number](value, min, max T) T {
	if value > max {
		return max
	} else if value < min {
		return min
	} else {
		return value
	}
}

