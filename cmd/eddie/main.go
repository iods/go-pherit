package main

import (
	"fmt"
	"io/ioutil"
	"log"
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
func scanDirectory(path string) error {
	fmt.Println("Scanning:", path)
	files, err := ioutil.ReadDir(path) // get a slice with the dir contents
	if err != nil {
		fmt.Printf("Returning error from scanDirectory(\"%s\") call \n", path)
		return err
	}

	for _, file := range files {

		filePath := filepath.Join(path, file.Name())

		if file.IsDir() {
			err := scanDirectory(filePath)
			if err != nil {
				fmt.Printf("Returning error from scanDirectory(\"%s\") call \n", path)
				return err
			}
		} else {
			fmt.Println(filePath)
		}

		//fmt.Println()
		//fmt.Printf("%s | is dir: %t | Size: %d | %s\n", file.Name(), file.IsDir(), file.Size(), file.Mode())
		//fmt.Printf("%s\n", file.ModTime())
	}

	return nil
}

func main() {
	err := scanDirectory("track")
	if err != nil {
		log.Fatal(err)
	}
	panic("oh no, sugar we are going down.")
}