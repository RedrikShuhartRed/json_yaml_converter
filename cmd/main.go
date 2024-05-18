package main

import (
	"fmt"
	"log"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	for {
		var convType string
		fmt.Print("Выберите формат конвертакции: 1-Json To Yaml, 2-Yaml To Json, 3-Exit\n")
		fmt.Scanln(&convType)
		switch convType {
		case "1":
			fmt.Println("конвертация JtoY")

		case "2":
			fmt.Println("Конвертация YtoJ")
		case "3":
			return
		default:
			fmt.Println("Некорректный выбор. Пожалуйста, выберите 1, 2 или 3.")
		}
	}
}
