package main

import "fmt"

type Preparer interface {
	PrepareDish()
}

type Pasta string
type Salad string

func (p Pasta) PrepareDish() {
	fmt.Println("Boil Water")
	fmt.Println("Throw in pasta")
	fmt.Println("Drain water")
}

func (s Salad) PrepareDish() {
	fmt.Println("Chop salad")
	fmt.Println("Add Dressing")
}

func PrepareDishes(dishes []Preparer) {
	fmt.Println("Prepare dishes:")
	for i := 0 ; i < len(dishes) ; i++ {
		dish := dishes[i]
		fmt.Printf("--Dish: %v--\n", dish)
		dish.PrepareDish()
	}
	fmt.Println()
}

func main() {
	dishes := []Preparer{Pasta("Ravioli"), Salad("Caesar Salad")}
	PrepareDishes(dishes)
}