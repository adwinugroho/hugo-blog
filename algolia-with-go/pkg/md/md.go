package md

import (
	"fmt"
	"io/ioutil"
	"log"
)

func ReadAllFilesInDirectory(dir string) ([]string, error) {
	filesInDir, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Printf("error cause:%+v\n", err)
		return nil, err
	}
	//var files = []string{}
	var arrayFileName []string
	for _, file := range filesInDir {
		arrayFileName = append(arrayFileName, file.Name())
	}
	return arrayFileName, nil
}

func FileToString(file string) (string, error) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("error cause:%+v\n", err)
		return "", err
	}
	var stringBytes = string(bytes)
	return stringBytes, nil
}

func GetContent(mdString string) {

}
