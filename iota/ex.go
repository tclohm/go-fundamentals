package main

import "fmt"

func main() {
	const (
		s0 = iota
		s1
		s2
		s3
		_
		s5
		s6
	)

	fmt.Println(s1, s2, s3, s5, s6)
}