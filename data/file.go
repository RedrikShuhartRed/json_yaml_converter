package file

import (
	"io"
	"log"
	"os"
)

func ReadFile(link string) ([]byte, error) {
	openFile, err := os.Open(link)
	if err != nil {
		log.Printf("error %s", err)
	}
	readFile, err := io.ReadAll(openFile)
	if err != nil {
		panic(err)
	}

	return readFile, nil
}

func CreateFile(convType string) (*os.File, error) {
	if convType == "1" {
		CreateFile, err := os.Create("resultYaml")
		if err != nil {
			panic(err)
		}
		return CreateFile, nil
	}

	CreateFile, err := os.Create("resultJson")
	if err != nil {
		panic(err)
	}
	return CreateFile, nil

}

func WriteResult(Createfile *os.File, readFile []byte) {

	_, err := Createfile.Write(readFile)
	if err != nil {
		panic(err)
	}
}
