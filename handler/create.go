package handler

import (
	"fmt"
	"os"
	"io"
)

func Start(list []string) {
	getOut := callHandling(list, 1)
	if getOut { return }

	file, err := os.Create(list[0])
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Printf("file %s created sucefully\n", list[0])
	}
	defer file.Close()
}

func SDir(list []string) {
	getOut := callHandling(list, 1)
	if getOut { return }

	err := os.Mkdir(list[0], 0755)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println("directory created")
}

func Copy(list []string) {
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
	fmt.Printf("Copy of %s in %s done\n", list[0], list[1])
}