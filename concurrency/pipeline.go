package main

import "fmt"

// construct streaming data pipelines 

// receive values from upstream via inbounds channels
// perform some function on that data, producing new values
// send values downstream via outbound channels

// FIRST: source or producer, LAST: sink or consumer

// MARK: -- gen converst a list of intergers to a channel that emits the integers in the list
// start a goroutine that sends the integers on the channel and closes the channel when all values
// have been sent
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

func main() {
	// Set up the pipeline
	for n := range sq(sq(gen(2, 3))) {
		fmt.Println(n)
	}
	
}