package main

import "fmt"

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	// dot notation auto derefences pointers
	counter.hits += 1
	fmt.Println("Counter", counter)
}

func replace(old *string, new string, counter *Counter) {
	// derefence old and equall to new
	*old = new
	increment(counter)
}

func main() {
	c := Counter{}

	hello := "Hello"
	world := "World!"
	fmt.Println(hello, world)

	replace(&hello, "Hi", &c)

	fmt.Println(hello, world)

	phrase := []string{hello, world}
	fmt.Println(phrase)

	replace(&phrase[1], "Go!", &c)

	fmt.Println(phrase)
}