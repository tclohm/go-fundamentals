package main

import (
	"regexp"
	"fmt"
)

var EmailExpr *regexp.Regexp

func init() {
	compiled, ok := regex.Compile(`.+@.+\..+`)
	if ok != nil {
		panic("failed to compile regular expression")
	}
	EmailExpr = compiled
	fmt.Println("reg exp compiled successfully")
}

func isValidEmail(add string) bool {
	return EmailExpr.Match([]byte(addr))
}

func main() {
	fmt.Println(isValidEmail("invalid"))
	fmt.Println(isValidEmail("valid@example.com"))
	fmt.Println(isValidEmail("invalid@example"))
}