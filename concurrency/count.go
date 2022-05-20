package main

import (
	"fmt"
	"sync"
	"strings"
	"unicode"
	"bufio"
	"os"
)

type Count struct {
	count int
	sync.Mutex
}

func getWords(line string) []string {
	return strings.Split(line, " ")
}

func countLetter(word string) int {
	letters := 0
	for _, r := range word {
		if unicode.IsLetter(r) {
			letters++
		}
	}
	return letters
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	totalLetters := Count{}

	var wg sync.WaitGroup

	for {
		if scanner.Scan() {
			line := scanner.Text()
			words := getWords(line)

			for _, word := range words {
				copy := word
				wg.Add(1)
				go func() {
					totalLetters.Lock()
					defer totalLetters.Unlock()
					defer wg.Done()
					sum := countLetter(copy)
					totalLetters.count += sum
				}()
			}
		} else {
			break
		}
	}

	wg.Wait()
	totalLetters.Lock()
	sum := totalLetters.count
	totalLetters.Unlock()

	fmt.Println("total letters =", sum)
}