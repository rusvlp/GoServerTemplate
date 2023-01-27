package util

import (
	"fmt"
	"io"
	"os"
)

func ReadFile(path string) (res string, err error) {
	file, err := os.Open(path)
	defer file.Close()
	//fileLength, err := file.Read([]byte{})

	data := make([]byte, 64)

	for {
		n, errRead := file.Read(data)
		fmt.Println(n)
		if errRead == io.EOF {
			break
		}
		fmt.Println(string(data[:n]))
	}

	return "", nil
}

func ReadFileAsBytes(path string) ([]byte, error) {
	res, err := ReadFile(path)
	return []byte(res), err
}

func IsFileExits(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else {
		return false
	}
}
