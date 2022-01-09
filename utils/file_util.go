package utils

import (
	"errors"
	"io/ioutil"
	"log"
)

func ReadFile(filePath string) ([]byte, error) {

	if IsEmpty(filePath) {
		return nil, errors.New("File path can not be null")
	}
	log.Printf("File reading from path. filepath is -> %s", filePath)
	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return data, nil

}
