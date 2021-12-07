package main

import (
	"fmt"
	"time"
)


func main() {
	bufferedChannel := make(chan string, 3)

	go func() {
		bufferedChannel <-"first"
		fmt.Println("Sent 1st")
		bufferedChannel <-"second"
		fmt.Println("Sent 2nd")
		bufferedChannel <-"third"
		fmt.Println("Sent 3rd")
	}()

	<-time.After(time.Second * 2)

	go func() {
		firstRead := <-bufferedChannel
		fmt.Println("Receiving..")
		fmt.Println(firstRead)
		secondRead := <-bufferedChannel
		fmt.Println(secondRead)
		thirdRead := <-bufferedChannel
		fmt.Println(thirdRead)
	}()

	fmt.Println("channels are weird...")
}