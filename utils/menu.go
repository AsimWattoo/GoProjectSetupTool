package utils

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

type MenuOption struct {
	Name    string
	Handler func() bool
}

func ShowInteractiveMenu(options []MenuOption) int {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))

	if err != nil {
		panic(err)
	}

	defer term.Restore(int(os.Stdin.Fd()), oldState)

	selected := 0
	for {
		ClearScreen()
		fmt.Println("Please select an option (Use ↑ ↓ arrows, Enter to select): ")
		for i, option := range options {
			if i == selected {
				fmt.Printf("\033[32m> %s\033[0m\n", option.Name)
			} else {
				fmt.Printf("  %s\n", option.Name)
			}
		}
		var buf [3]byte
		os.Stdin.Read(buf[:])

		if buf[0] == 27 && buf[1] == 91 {
			switch buf[2] {
			case 65: // Up arrow
				if selected > 0 {
					selected--
				} else {
					selected = len(options) - 1
				}
			case 66: // Down arrow
				if selected < len(options)-1 {
					selected++
				} else {
					selected = 0
				}
			}
		} else if buf[0] == 13 {
			return selected
		} else if buf[0] == 3 {
			os.Exit(0)
		}
	}
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func MenuLoop(menuItems []MenuOption) {
	for {
		choice := ShowInteractiveMenu(menuItems)

		if choice != -1 {
			if !menuItems[choice].Handler() {
				break
			}
		}

		WaitForEnter()
	}
}

func WaitForEnter() {
	fmt.Println("Press Enter to continue...")
	for {
		var b [1]byte
		n, err := os.Stdin.Read(b[:])
		if err != nil || n == 0 || b[0] == '\n' {
			break
		}
	}
}
