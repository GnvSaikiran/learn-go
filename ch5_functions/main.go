package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filesize, err := fileLen("main.go")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(filesize)

	helloPrefixer := prefixer("Hello")
	fmt.Println(helloPrefixer("Mark"))
	fmt.Println(helloPrefixer("Bob"))
}

func fileLen(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var fileSize int
	data := make([]byte, 2048)
	for {
		count, err := file.Read(data)
		if err != nil {
			if err != io.EOF {
				return 0, err
			}
			break
		}
		fileSize += count
	}

	return fileSize, nil
}

func prefixer(prefix string) func(string) string {
	return func(s string) string {
		return prefix + " " + s
	}
}
