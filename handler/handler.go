package handler

import (
	"log"
	"path/filepath"

	"github.com/RedrikShuhartRed/json_yaml_converter/converter"
	"github.com/RedrikShuhartRed/json_yaml_converter/data"
)

func ConvJsonToYaml(link string, convType string) (string, error) {
	readFile, err := data.ReadFile(link)
	if err != nil {
		log.Printf("error ReadFIle in handler: %s", err)
		return "", err
	}

	converter, err := converter.JsonToYaml(readFile)
	if err != nil {
		log.Printf("error JsonToYaml in handler: %s", err)
		return "", err
	}
	createFile, path, err := data.CreateFile(link, convType)
	if err != nil {
		log.Printf("error CreateFile in handler: %s", err)
		return "", err
	}
	err = data.WriteResult(createFile, converter)
	if err != nil {
		log.Printf("error WriteResult in handler: %s", err)
		return "", err
	}
	err = createFile.Close()
	if err != nil {
		log.Printf("error close file in handler: %s", err)
		return "", err
	}
	resultPath, err := filepath.Abs(path)
	return resultPath, nil
}

func ConvYamlToJson(link string, convType string) (string, error) {

	readFile, err := data.ReadFile(link)
	if err != nil {
		log.Printf("error ReadFIle in handler: %s", err)
		return "", err
	}
	converter, err := converter.YamlToJson(readFile)
	if err != nil {
		log.Printf("error YamlToJson in handler: %s", err)
		return "", err
	}
	createFile, path, err := data.CreateFile(link, convType)
	if err != nil {
		log.Printf("error CreateFile in handler: %s", err)
		return "", err
	}
	err = data.WriteResult(createFile, converter)
	if err != nil {
		log.Printf("error WriteResult in handler: %s", err)
		return "", err
	}
	err = createFile.Close()
	if err != nil {
		log.Printf("error close file in handler: %s", err)
		return "", err
	}
	resultPath, err := filepath.Abs(path)
	return resultPath, nil
}
