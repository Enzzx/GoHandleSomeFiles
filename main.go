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
	"readFile": handler.ReadFile,
}


func main() {
	reader := bufio.NewReader(os.Stdin)
	for  {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		inputList := strings.Fields(input)

		if cmd, ok := commandMap[inputList[0]]; ok {
			cmd(inputList[1:])
		} else {
			fmt.Println("Pass a valid command")
		}
	}
}
