package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

/*
cursed Returns an example of recursion.
*/
func cursed(start int, end int) {
	fmt.Printf("cursed(%d, %d) called.\n", start, end)
	fmt.Println("Starting at:", start)
	if start < end {
		cursed(start + 1, end)
	}
	fmt.Printf("Returning from cursed(%d, %d) called.\n", start, end)
}

/*
scanDirectory Scans and prints the contents of the directory path provided.
*/
func scanDirectory(path string) {
	fmt.Println("Scanning:", path)
	files, err := ioutil.ReadDir(path) // get a slice with the dir contents
	if err != nil {
		panic(err)
	}

	for _, file := range files {

		filePath := filepath.Join(path, file.Name())

		if file.IsDir() {
			scanDirectory(filePath)
		} else {
			fmt.Println(filePath)
		}
	}
}

func main() {
	scanDirectory("track")
}