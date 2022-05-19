package main

// con -- multi line to be execute
// threaded -- parallel based on # of cpu cores
// async 	-- pause and resume execution -- main thread, other threads

// single threaded -- deterministic
// concurrent -- non-deterministic

import (
	"time"
	"fmt"
)

func count(amount int) {
	for i := 1 ; i <= amount ; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i)
	}
}

func main() {
	go count(5)
	fmt.Println("wait for goroutine")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("end program")
}