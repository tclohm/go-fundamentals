package main

import (
	"fmt"
	"unicode"
)

type LineCallback func(line string)

func lineIterator(lines []string, callback LineCallback) {
	for _, line := range lines {
		callback(string(line))
	}
}

func main() {
	lines := []string{
		"Somewhere over",
		"the rainbow",
		"way up high.",
		"Average person",
		"for an",
		"Average person",
		"woah!",
		"Three people are in",
		"on it!",
		"Why can't I find my 14 private keys",
	}

	letters := 0
	numbers := 0
	punctuation := 0
	spaces := 0

	lineFunc := func(line string) {
		for _, rune := range line {
			if unicode.IsLetter(rune) {
				letters++
			}
			if unicode.IsDigit(rune) {
				numbers++
			}
			if unicode.IsPunct(rune) {
				punctuation++
			}
			if unicode.IsSpace(rune) {
				spaces++
			}
		}
	}

	lineIterator(lines, lineFunc)

	fmt.Println(letters, "letters")
	fmt.Println(numbers, "numbers")
	fmt.Println(punctuation, "punctuation")
	fmt.Println(spaces, "spaces")
}