package main

import (
	"fmt"
	"log"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	var link string
	fmt.Print("Введите путь до файла в формате:D:/somefolder/somefolder/yourfile.txt\n")
	fmt.Scanln(&link)
	for {
		var convType string

		fmt.Print("Выберите формат конвертакции: 1-Json To Yaml, 2-Yaml To Json, 3-Exit\n")
		fmt.Scanln(&convType)

		switch convType {
		case "1":
			fmt.Println("конвертация JtoY")
			handler.convJsonToYaml(link, convType)

		case "2":
			fmt.Println("Конвертация YtoJ")
			handler.convYamlToJson(link, convType)
		case "3":
			return
		default:
			fmt.Println("Некорректный выбор. Пожалуйста, выберите 1, 2 или 3.")
		}
	}
}
