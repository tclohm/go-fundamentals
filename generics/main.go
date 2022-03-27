package main

import "fmt"
// MARK: -- finally!
func Print[T ~int|~string|~float64](t T) {
	fmt.Printf("%d\n", t)
}

func main() {
	Print("Hello World")
	Print(123)
	Print(123.4)
}