package main

import (
	"fmt"
	"time"
	//"sync"
)

func main() {
	//var wg sync.WaitGroup
	//wg.Add(1)
	//go count("sheep") // goroutine -- do it in the background, concurrency! Constrains by cores
	//go count("cow")

	// go func() {
	// 	count("sheep")
	// 	wg.Done()
	// }()

	// wg.Wait() // wait until we are done

	//c := make(chan string)
	//go count("sheep", c)

	// for {
	// 	msg, open := <- c

	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(msg)
	// }

	// for msg := range c {
	// 	fmt.Println(msg)
	// }

	// c := make(chan string, 2)
	// c <- "hello"
	// c <- "world"

	// msg := <- c
	// fmt.Println(msg)

	// msg = <- c
	// fmt.Println(msg)

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every two second"
			time.Sleep(time.Second * 2)
		}
	}()

	// for {
	// 	fmt.Println(<-c1)
	// 	fmt.Println(<-c2)
	// }

	for {
		select {
		case msg1 := <- c1:
			fmt.Println(msg1)
		case msg2 := <- c2:
			fmt.Println(msg2)
		}
	}
}

func count(thing string, c chan string) {
	for i := 1 ; i <= 5 ; i++ {
		//fmt.Println(i, thing)
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}