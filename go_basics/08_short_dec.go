package main

import "fmt"

func main() {
	quantity := 4
	length, width := 1.2, 4.4
	customer := "Jeremy Forrest"

	fmt.Println(customer)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with with an area of")
	fmt.Println(length*width, "square meters")	
}