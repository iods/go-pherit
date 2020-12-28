package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func cursed() {
	fmt.Println("Oh shit, I am stuck in an infinite loop.")
	cursed()
}

func main() {
	files, err := ioutil.ReadDir("track")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Println("Directory:", file.Name())
		} else {
			fmt.Println("Filename:", file.Name())
		}
	}

	cursed()
}