package main

// wait group -- synchronizing
// channels -- communication

import (
	"fmt"
	"time"
	"sync"
	"math/rand"
)

func doWork(d time.Duration, wg *sync.WaitGroup) {
	fmt.Println("doing work....")
	time.Sleep(d)
	fmt.Println("done")
	wg.Done()
}

func tell(d time.Duration, resChan chan string) {
	fmt.Println("Hey, I'm doing work....")
	time.Sleep(d)
	resChan <- fmt.Sprintf("I'm done with job ID: %v", rand.Intn(100))
}


func main() {
	start := time.Now()
	//wg := sync.WaitGroup{}
	//wg.Add(2)
	//go doWork(time.Second * 2, &wg)
	//go doWork(time.Second * 4, &wg)
	//wg.Wait()

	resultChan := make(chan string)
	go tell(time.Second * 2, resultChan)
	go tell(time.Second * 4, resultChan)

	res1 := <- resultChan
	fmt.Println(res1)
	res2 := <- resultChan
	fmt.Println(res2)
	fmt.Printf("work took %v seconds\n", time.Since(start))
}