package main

import "fmt"

// activate and deactivate security tags on products

const (
	Active = true
	Inactive = false
)

type SecurityTag bool

type Item struct {
	name string
	tag SecurityTag
}

func activate(tag *SecurityTag) {
	*tag = Active
}

func deactivate(tag *SecurityTag) {
	*tag = Inactive
}

func checkout(items []Item) {
	fmt.Println("checking out...")
	for i := 0 ; i < len(items) ; i++ {
		fmt.Println("deactivating", items[i])
		deactivate(&items[i].tag)
	}
}

func main() {
	items := []Item{
		Item{name: "shirt", tag: true},
		Item{name: "watch", tag: true},
		Item{name: "apple", tag: true},
		Item{name: "pants", tag: false},
	}

	fmt.Println("original:", items)

	checkout(items)

	fmt.Println("deactivated:", items)
}