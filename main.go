package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	inputDir := "./input"

	files, err := ioutil.ReadDir(inputDir)
	if err != nil {
		fmt.Println("Read file error:", err)
		return
	}

	for _, file := range files {
		infile, err := os.Open(filepath.Join(inputDir, file.Name()))
		if err != nil {
			fmt.Printf("Open file %s error：%s\n", file.Name(), err)
			return
		}
		defer infile.Close()

	}
}
