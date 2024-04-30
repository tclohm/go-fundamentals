package main

// wait group -- synchronizing
// channels -- communication

import (
	"fmt"
	"time"
)

func doWork(d time.Duration) {
	fmt.Println("doing work....")
	time.Sleep(d)
	fmt.Println("done")
}

func main() {
	start := time.Now()
	go doWork(time.Second * 2)
	go doWork(time.Second * 4)
	fmt.Printf("work took %v seconds\n", time.Since(start))
}