package main

import (
	"fmt"
	"os"
)

func ReadFile(fileName string) (string, error) {
	//
	file, error := os.Open(fileName)
	if error != nil {
		fmt.Println("")
		return "", error
	}
	//
	buf := make([]byte, 1024)
	defer file.Close()
	var content string
	for {
		n, _ := file.Read(buf)
		if n == 0 {
			break
		}
		content += string(buf[0:n])
	}
	return content, nil
}

func main() {
	//
	fileTxt, error := ReadFile("main.go")
	if error != nil {
		fmt.Println(error.Error())
	}

	fmt.Println(fileTxt)
}
