package main

import (
	"strings"
	"fmt"
)

func custom(fn func(msg string), msg string) {
	msg = strings.ToUpper(msg)
	fn(msg)
}

func surround() func(msg string) {
	return func(msg string) {
		fmt.Printf("%.*s\n", len(msg), "--------")
		fmt.Println(msg)
		fmt.Printf("%.*s\n", len(msg), "--------")
	}
}

type DiscountFunc func(subTotal float64) float64

func calculatePrice(
	subTotal float64, 
	discountFn DiscountFunc) float64 {
		return subTotal - (subTotal * discountFn(subTotal))
	}

func main() {
	custom(surround(), "good day")
}