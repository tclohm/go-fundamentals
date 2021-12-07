package main

import (
	"fmt"
	"time"
)

func main() {
	mainMine := [5]string{"rock", "ore", "ore", "rock", "ore"}
	// channels block go routines in various situations
	// allows go routines to sync up with each other for a moment
	// two types of channels: unbuffered 
	oreChannel := make(chan string)
	minedOreChannel := make(chan string)
	// buffered
	// oreChannel := make(chan string, 3)

	// Finder goroutine
	go func(mine [5]string) {
		for _, item := range mine {
			if item == "ore" {
				oreChannel <- item // send
			}
		}
	}(mainMine)

	// Ore Breaker goroutine
	go func() {
		for i := 0 ; i < 3 ; i++ {
			foundOre := <-oreChannel // received
			fmt.Println("From Finder: ", foundOre)
			minedOreChannel <- "minedOre" // send to minedOreChannel
		}
	}()

	// Smelter
	go func() {
		for i := 0 ; i < 3 ; i++ {
			minedOre := <- minedOreChannel // read from minedOreChannel
			fmt.Println("From Miner: ", minedOre)
			fmt.Println("From Smelter: Ore is smelted")
		}
	}()

	<-time.After(time.Second * 5)
}