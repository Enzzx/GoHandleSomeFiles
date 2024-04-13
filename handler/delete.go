package handler

import (
	"fmt"
	"os"
)

func Del(file []string) {
	getOut := callHandling(file, 1)
	if getOut { return }

	err := os.RemoveAll(file[0])
	if err != nil {
		fmt.Printf("Cannot delete %s: %s\n", file[0], err)
	} else {
		fmt.Println(file[0], " deleted")
	}
}