package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"GoHandleSomeFiles/fileHandler/handler"
)

type commandFunc func([]string)

var commandMap = map[string]commandFunc {
	//create.go
	"start": handler.Start,
	"startDir": handler.SDir,
	"copy": handler.Copy,

	//read.go
	"read": handler.Read,
	"show": handler.Show,
	"this": handler.Show,
	"info": handler.Info,
	"exist": handler.Exist,
	"help": handler.Help,

	//update.go
	"add": handler.Add,
	"switch": handler.Switch,
	"move": handler.Move,
	"goto": handler.Goto,

	//delete.go
	"del": handler.Del,
}


func main() {
	reader := bufio.NewReader(os.Stdin)
	for  {
		pwd, _ := os.Getwd()
		fmt.Print(pwd, " >> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		inputList := strings.Fields(input)

		if cmd, ok := commandMap[inputList[0]]; ok {
			cmd(inputList[1:])
		} else {
			fmt.Println("Pass a valid command")
		}
		fmt.Println("")
	}
}