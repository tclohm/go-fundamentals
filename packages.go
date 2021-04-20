package main

import (
	"fmt"
	"go/utils"
)

func calcuateImportantData() int {
	totalValue := utils.Add(5, 4)
	return totalValue
}

func main() {
	fmt.Println("Packages!")
	total := calcuateImportantData()
	fmt.Println(total)
}