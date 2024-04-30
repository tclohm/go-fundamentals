package main

// wait group -- synchronizing
// channels -- communication

import (
	"fmt"
	"time"
	"sync"
)

func doWork(d time.Duration, wg *sync.WaitGroup) {
	fmt.Println("doing work....")
	time.Sleep(d)
	fmt.Println("done")
	wg.Done()
}


func main() {
	start := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(2)
	go doWork(time.Second * 2, &wg)
	go doWork(time.Second * 4, &wg)
	wg.Wait()
	fmt.Printf("work took %v seconds\n", time.Since(start))
}