package util

import (
	"errors"
	"io"
	"os"
)

func ReadFile(path string) (res []byte, err error) {
	if !IsFileExits(path) {
		return nil, errors.New("File Not Exist")
	}
	file, err := os.Open(path)
	defer file.Close()
	//fileLength, err := file.Read([]byte{})

	data := make([]byte, 64)
	tmpData := make([]byte, 256)
	for {
		n, errRead := file.Read(tmpData)
		if errRead == io.EOF {
			break
		}
		tmpData = append(tmpData[:n], data...)
	}

	return tmpData, nil
}

func ReadFileAsString(path string) (string, error) {
	res, err := ReadFile(path)
	return string(res), err
}

func IsFileExits(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else {
		return false
	}
}
