package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	doCommand(os.Args[1:])
}

func doCommand(args []string) {
	if len(args) < 2 || args[0] == "-h" {
		showHelp()
		return
	}

	//源文件名 -> 可能会包含路径
	from := args[0]

	//目标名
	to := args[1]

	//源文件目录
	dir := path.Dir(from)

	newName := dir + "/" + to

	err := os.Rename(from, newName)

	if err != nil {
		fmt.Println("Err:", err)
		os.Exit(1)
	}
}

func showHelp() {
	fmt.Println("Usage: $ rename [from] [to] [options]")
}
