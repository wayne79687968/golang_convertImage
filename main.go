package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	inputDir := "./input"

	files, err := ioutil.ReadDir(inputDir)
	if err != nil {
		fmt.Println("Read file error:", err)
		return
	}
}
