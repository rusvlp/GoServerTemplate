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
	tmpData := make([]byte, 64)
	resLength := 0
	for i := 1; ; i++ {
		n, errRead := file.Read(tmpData)
		resLength += n
		if errRead == io.EOF {
			break
		}
		data = append(data, tmpData[:n]...)
	}

	return data, nil
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
