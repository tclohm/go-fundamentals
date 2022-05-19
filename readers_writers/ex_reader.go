package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

const (
	CmdHello = "hello"
	CmdGoodbye = "bye"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	numLines := 0
	numCommands := 0
	for scanner.Scan() {
		if strings.ToUpper(scanner.Text()) == "Q" {
			break
		} else {
			text := strings.TrimSpace(scanner.Text())

			switch text {
			case CmdHello:
				numCommands += 1
				fmt.Println("command response: hi")
			case CmdGoodbye:
				numCommands += 1
				fmt.Println("command response: bye")
			}
			if text != "" {
				numLines += 1
			}
		}
	}

	fmt.Printf("Youve entered %v lines\n", numLines)
	fmt.Printf("Youve entered %v commands\n", numCommands)
}