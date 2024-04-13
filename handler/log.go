package handler

import (
	"fmt"

)

func callHandling(file []string, size int) bool {
	value := false
	if len(file) < size {
		fmt.Println("Missing arguments")
		value = true
	}

	if len(file) > size {
		fmt.Println("Arguments out of range")
		value = true
	}

	return value
}