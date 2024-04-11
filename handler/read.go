package handler

import (
	"fmt"
	"log"
	"os"
)

func ReadFile(file []string) {
	_, err := os.Stat(file[0])
	if os.IsNotExist(err) {
		fmt.Println("esse arquivo não existe, ", err)
		return
	} else {
		fmt.Println("arquivo encontrado")
		data, err := os.ReadFile(file[0])
		if err != nil {
			log.Fatal(err) 
			fmt.Println("conteúdo do arquivo não pôde ser lido")
		} else {
			fmt.Printf("Conteúdo do arquivo:\n %s", string(data))
			return
		}
	}
}
