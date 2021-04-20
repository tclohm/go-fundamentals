package main

import (
	"fmt"
	"strings"
)

func changeName(n *string) {
	*n = strings.ToUpper(*n)
}

func main() {
	var name string
	var namePointer *string

	name = "Lizzo"
	namePointer = &name // address
	var nameValue = *namePointer // gives value

	fmt.Println("name: ", name)
	fmt.Println("name: *", namePointer)
	fmt.Println("name value:", nameValue)

	changeName(&nameValue)

	fmt.Println(nameValue)

}