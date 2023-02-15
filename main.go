package main

import (
	"fmt"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
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
		if !strings.HasSuffix(file.Name(), ".png") &&
			!strings.HasSuffix(file.Name(), ".gif") {
			fmt.Printf("Skip file %s, unsupported format\n", file.Name())
			continue
		}

		infile, err := os.Open(filepath.Join(inputDir, file.Name()))
		if err != nil {
			fmt.Printf("Open file %s error：%s\n", file.Name(), err)
			return
		}
		defer infile.Close()

		img, _, err := image.Decode(infile)
		if err != nil {
			fmt.Printf("Decode image %s error：%s\n", file.Name(), err)
			return
		}

		outFileName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name())) + ".jpg"
		outfile, err := os.Create(filepath.Join(outputDir, outFileName))
		if err != nil {
			fmt.Printf("Create file %s error：%s\n", outFileName, err)
			return
		}
		defer outfile.Close()

		err = jpeg.Encode(outfile, img, &jpeg.Options{Quality: 100})
		if err != nil {
			fmt.Printf("Encode file %s error：%s\n", outFileName, err)
			return
		}

		fmt.Printf("Converted successfully %s\n", outFileName)
	}
}
