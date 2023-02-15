package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	inputDir := "./input"
	outputDir := "./output"

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

		outFileName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name())) + ".jpg"
		outfile, err := os.Create(filepath.Join(outputDir, outFileName))
		if err != nil {
			fmt.Printf("Create file %s error：%s\n", outFileName, err)
			return
		}
		defer outfile.Close()
	}
}
