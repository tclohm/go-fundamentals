package main

// channels -- one way communciation
// sending and receiving

import "fmt"

func main() {
	channel := make(chan int)
	go func() { channel <- 1 }()
	go func() { channel <- 2 }()
	go func() { channel <- 3 }()

	first := <-channel
	second := <-channel
	third := <-channel

	fmt.Println(first, second, third)
}