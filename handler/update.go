package handler

import (
	"fmt"
	"os"
	"io"
	"strings"
)

func Add(list []string) {
	fileContent := strings.Join(list[0:len(list)-1], " ")
	fileContent = "\n" + fileContent
	//fmt.Println(fileContent)
	newList := make([]string, 0, 1)
	newList = append(newList, fileContent)
	newList = append(newList, list[len(list)-1])

	getOut := callHandling(newList, 2)
	if getOut { return }

	file, err := os.OpenFile(newList[1], os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(fileContent)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println("text inserted")
}

func Switch(list []string) {
	getOut := callHandling(list, 2)
	if getOut { return }

	err := os.Rename(list[0], list[1])
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println("name shitched")
}

func Move(list []string) {
	getOut := callHandling(list, 2)
	if getOut { return }

	srcFile, err := os.Open(list[0])
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	defer srcFile.Close()

	destFile, err := os.Create(list[1])
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	err = os.Remove(list[0])
	if err != nil {
		fmt.Println("error on removing original file")
	}
	fmt.Println("Move done")
}

func Goto(list []string) {
	getOut := callHandling(list, 1)
	if getOut { return }

	err := os.Chdir(list[0])
	if err != nil {
		fmt.Println("err: ", err)
	}
}