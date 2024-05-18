package converter

import (
	"encoding/json"
	"log"

	"gopkg.in/yaml.v3"
)

func JsonToYaml(jsonData []byte) ([]byte, error) {

	var data interface{}

	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Printf("error unmarshal json file: %s\n", err)
		return nil, err
	}

	yamlData, err := yaml.Marshal(data)

	if err != nil {
		log.Printf("error marshal yaml: %s\n", err)
		return nil, err
	}

	return yamlData, nil

}

func YamlToJson(yamlData []byte) ([]byte, error) {

	var data interface{}

	err := yaml.Unmarshal(yamlData, &data)
	if err != nil {
		log.Printf("error unmarshal yaml file: %s\n", err)
		return nil, err
	}

	jsonData, err := json.MarshalIndent(data, "", "   ")

	if err != nil {
		log.Printf("error marshalIndent json: %s\n", err)
		return nil, err
	}

	return jsonData, nil

}
