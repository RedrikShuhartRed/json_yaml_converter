package main

import (
	"fmt"
	"log"

	"github.com/RedrikShuhartRed/json_yaml_converter/handler"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	var link string
	for {
		fmt.Print("Enter the path to the file in the format:D:/somefolder/somefolder/yourfile.txt\n")
		_, err := fmt.Scanln(&link)
		if err != nil {
			log.Printf("error Scan link: %s\n", err)
			return
		}

		var convType string

		fmt.Print("Select conversion format: 1-Json To Yaml, 2-Yaml To Json, 3-Exit\n")
		_, err = fmt.Scanln(&convType)
		if err != nil {
			log.Printf("error conversation type: %s\n", err)
			return
		}

		switch convType {
		case "1":
			fmt.Println("JtoY conversion")
			path, _ := handler.ConvJsonToYaml(link, convType)
			fmt.Println("Your Yaml file:", path)

		case "2":
			fmt.Println("YtoJ conversion")
			path, _ := handler.ConvYamlToJson(link, convType)
			fmt.Println("Yot Json file:", path)

		case "3":
			return
		default:
			fmt.Println("Incorrect choice. Please select 1, 2 or 3.")
		}
	}
}
