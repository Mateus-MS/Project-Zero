package main

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func main() {
	options := []string{"Option 1", "Option 2", "Option 3", "Quit"}
	selected := 0

	// Put terminal in raw mode
	oldState, _ := term.MakeRaw(int(os.Stdin.Fd()))
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	printMenu := func() {
		fmt.Print("\033[H\033[2J") // clear screen
		for i, opt := range options {
			if i == selected {
				fmt.Printf("> %s\n", opt)
			} else {
				fmt.Printf("  %s\n", opt)
			}
		}
	}

	buf := make([]byte, 3)
	for {
		printMenu()
		n, _ := os.Stdin.Read(buf)

		if n == 1 && buf[0] == '\n' { // Enter
			fmt.Printf("You selected: %s\n", options[selected])
			break
		}

		if n == 3 { // Arrow keys
			if buf[0] == 27 && buf[1] == 91 {
				switch buf[2] {
				case 65: // Up
					if selected > 0 {
						selected--
					}
				case 66: // Down
					if selected < len(options)-1 {
						selected++
					}
				}
			}
		}
	}
}
