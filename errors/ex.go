package main

import (
	"errors"
	"fmt"
	"os"
)

func isGreaterThanTen(num int) error {
	if num < 10 {
		return errors.New("something bad")
	}
	return nil
}

func openFile() error {
	f, err := os.Open("missingFile.txt")
	if err != nil {
		return err
	}
	// LIFO
	defer f.Close()
	return nil
}

func main() {
	num := 9
	if err := isGreaterThanTen(num); err != nil {
		// panic kills the program
		//panic(err)
		fmt.Println(fmt.Errorf("%d is not greater than 10", num))
		// log the fatal errors
		// log.Fatalln(err)
	}

	// err is within that scope
	if err:= openFile(); err != nil {
		fmt.Println(fmt.Errorf("%v", err))
	}
}