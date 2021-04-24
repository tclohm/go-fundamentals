package main

import "fmt"

func main () {
	fmt.Printf("float: %f\n", 1.144)
	fmt.Printf("integer: %d\n", 23)
	fmt.Printf("string: %s\n", "hello world")
	fmt.Printf("boolean: %t\n", false)
	fmt.Printf("values: %v %v %v\n", 1.2, "\t", true)
	fmt.Printf("values: %#v %#v %#v\n", 1.2, "\t", false)
	fmt.Printf("types: %T %T %T\n", 1.2 "\t", true)
	fmt.Printf("percent sign: %%\n")
}