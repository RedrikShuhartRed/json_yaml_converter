package data

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

func ReadFile(link string) ([]byte, error) {
	openFile, err := os.Open(link)
	if err != nil {
		log.Printf("error open file, chek link: %s\n", err)
		return nil, err
	}
	readFile, err := io.ReadAll(openFile)
	if err != nil {
		log.Printf("error read file: %s\n", err)
		return nil, err
	}

	return readFile, nil
}

func CreateFile(link string, convType string) (*os.File, string, error) {

	base := filepath.Base(link)
	path := base[:len(base)-len(filepath.Ext(base))]

	if convType == "1" {
		path += "Yaml"
		CreateFile, err := os.Create(path)
		if err != nil {
			log.Printf("error create result file: %s\n", err)
			return nil, "", err
		}
		return CreateFile, path, nil
	}
	path += "Json"
	CreateFile, err := os.Create(path)
	if err != nil {
		log.Printf("error create result file: %s\n", err)
		return nil, "", err
	}
	return CreateFile, path, nil

}

func WriteResult(Createfile *os.File, readFile []byte) error {

	_, err := Createfile.Write(readFile)
	if err != nil {
		log.Printf("error write result in file: %s\n", err)
		return err
	}

	return nil
}
