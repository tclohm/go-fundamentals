package main
import "fmt"

var myInt int = 42
var myIntPointer = &myInt

func createPointer() *float64 {
	var myFloat = 98.5
	return &myFloat
}

func double(number *int) {
	*number *= 2
}

func main() {
	fmt.Printf("address %v\n", myIntPointer)
	fmt.Printf("value %d\n", *myIntPointer)

	var myFloatPointer *float64 = createPointer()
	fmt.Println(*myFloatPointer)

	fmt.Printf("\n")

	amount := 6

	fmt.Println(amount)
	double(&amount)
	fmt.Println(amount)


}