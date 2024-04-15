package handler

import (
	"fmt"
	"os"
)

func Help(list []string) {
	getOut := callHandling(list, 0)
	if getOut { return }

	fmt.Println("add\t\t - adds text to a file\ncopy\t\t - copies a file to another\ndel\t\t - deletes a file or directory\nexist\t\t - checks the existence of a file or directory\ninfo\t\t - provides details about a file\nmove\t\t - moves a file from one directory to another\nread\t\t - reads a file\nshow\t\t - lists the specified directory\nstart\t\t - creates a file\nstartDir\t - creates a directory\nswitch\t\t - renames a file or directory\nthis\t\t - lists the current directory\n\nYour first time? try GOing to elsewhere, use 'goto relative/path'")

}

func Read(list []string) {
	getOut := callHandling(list, 1)
	if getOut { return }

	_, err := os.Stat(list[0])
	if os.IsNotExist(err) {
		fmt.Println("The file cannot be found, ", err)
		return
	} else {
		//fmt.Println("arquivo encontrado")
		data, err := os.ReadFile(list[0])
		if err != nil {
			fmt.Println("err: ", err)
			return
		}
		fmt.Print(string(data))
	}
}

func Show(list []string) {
	if len(list) == 0 {
		list = append(list, "./")
	}

	dirEntries, err := os.ReadDir(list[0])
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	for _, entry := range dirEntries {
		fmt.Println(entry.Name())
	}
}

func Info(list []string) {
	getOut := callHandling(list, 1)
	if getOut { return }

	fileInfo, err := os.Stat(list[0])
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Printf("%s\t%dbytes\t%s\t%s\n", fileInfo.Name(), fileInfo.Size(), fileInfo.Mode(), fileInfo.ModTime())
}

func Exist(list []string) {
	getOut := callHandling(list, 1)
	if getOut { return }

	_, err := os.Stat(list[0])
	if os.IsNotExist(err) {
		fmt.Printf("%s does not exist\n", list[0])
		return
	}
	fmt.Printf("%s exist\n", list[0])
}