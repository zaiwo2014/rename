package main

import (
	"fmt"
	"os"
	"path"
	"strings"
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

	//当新文件名以 . 开头时, 只改变后缀名
	if to[0] == '.' {
		doRenameExt(from, to)
	} else if to[0] == ':' {
		doTransform(from, to)
	} else {
		doRename(from, to)
	}

}

func doRenameExt(from string, to string) {
	//源文件目录
	dir := path.Dir(from)

	part := strings.Split(path.Base(from), ".")[0]

	newName := dir + "/" + part + to

	err := os.Rename(from, newName)

	if err != nil {
		fmt.Println("Err:", err)
		os.Exit(1)
	}
}

func doRename(from string, to string) {
	//源文件目录
	dir := path.Dir(from)

	newName := dir + "/" + to

	err := os.Rename(from, newName)

	if err != nil {
		fmt.Println("Err:", err)
		os.Exit(2)
	}
}

func doTransform(from string, to string) {
	options := strings.Split(to, ":")[1]
	var targetName string
	switch options {
	case "upper":
		targetName = strings.ToUpper(strings.Split(path.Base(from), ".")[0])
	case "lower":
		targetName = strings.ToLower(strings.Split(path.Base(from), ".")[0])
	default:
		fmt.Println("Err:", "Command does not exist in options")
		os.Exit(1)
	}

	newName := path.Dir(from) + "/" + targetName + path.Ext(from)

	err := os.Rename(from, newName)

	if err != nil {
		fmt.Println("Err:", err)
		os.Exit(1)
	}

}

func showHelp() {
	fmt.Println("Usage: $ rename [from] [to] [options]")
}
