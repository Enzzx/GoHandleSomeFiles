package handler

import (
	"fmt"
	"os"
)

func Read(file []string) {
	getOut := callHandling(file, 1)
	if getOut { return }

	_, err := os.Stat(file[0])
	if os.IsNotExist(err) {
		fmt.Println("The file cannot be found, ", err)
		return
	} else {
		//fmt.Println("arquivo encontrado")
		data, err := os.ReadFile(file[0])
		if err != nil {
			fmt.Println("The file content connot be read")
		} else {
			fmt.Print(string(data))
			return
		}
	}
}

func Show(file []string) {
	if len(file) == 0 {
		file = append(file, "./")
	}

	dirEntries, err := os.ReadDir(file[0])
	if err != nil {
		fmt.Println("The directory cannot be found, ", err)
	} else {
		for _, entry := range dirEntries {
			fmt.Println(entry.Name())
		}
		return
	}
}

func Info(file []string) {
	getOut := callHandling(file, 1)
	if getOut { return }

	fileInfo, err := os.Stat(file[0])
	if err != nil {
		fmt.Println("The directory cannot be found, ", err)
	} else {
		fmt.Printf("%s\t%dbytes\t%s\t%s\n", fileInfo.Name(), fileInfo.Size(), fileInfo.Mode(), fileInfo.ModTime())
	}
}

func Have(file []string) {
	getOut := callHandling(file, 1)
	if getOut { return }

	_, err := os.Stat(file[0])
	if os.IsNotExist(err) {
		fmt.Printf("%s does not exist\n", file[0])
	} else {
		fmt.Printf("%s exist\n", file[0])
		return
	}
}