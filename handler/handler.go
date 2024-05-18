package handler

import (
	"github.com/RedrikShuhartRed/json_yaml_converter/converter"
)

func convJsonToYaml(link string, convType string) {

	readFile, _ := file.ReadFile(link)
	converter, _ := converter.JsonToYaml(readFile)
	createFIle, _ := file.CreateFile(convType)
	file.WriteResult(createFIle, converter)
}

func convYamlToJson(link string, convType string) {

	readFile, _ := file.ReadFile(link)
	converter, _ := converter.YamlToJson(readFile)
	createFIle, _ := file.CreateFile(convType)
	file.WriteResult(createFIle, converter)
}
