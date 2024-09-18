package internal

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

func SmartPrint(text string) {
	if terminal.IsTerminal(int(os.Stdout.Fd())) {
		fmt.Println(text)
	} else {
		fmt.Print(text)
	}
}
