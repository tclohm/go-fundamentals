package main

import (
	"fmt"
	"sync"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// mult functions read the same channel until that channel is closed; fan-out
// way to distribute work amonst a group of workers to parallelize CPU use and I/O

// func can read from multiple inputs and proceed until all are closed by multiplexing the input
// channels onto a single channel that's closed when all the inputs are closed; fan-in

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	// Start an output goroutine for each input channel in cs. output
	// copies values from c to out until c is closed, then calls wg.Done
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// start a goroutine to close out once all the output goroutine are done
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	in := gen(2, 3)
	c1 := sq(in)
	c2 := sq(in)

	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}