package main

import (
	"fmt"
	"os"
)

func main() {
	doCommand(os.Args[1:])
}

func doCommand(args []string) {
	if len(args) == 0 || args[0] == "-h" {
		showHelp()
		return
	}
}

func showHelp() {
	fmt.Println("this is help")
}
