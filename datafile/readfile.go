package datafile

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := "data.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err) // if there is an error opening the file, log it
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // read a line from the file
		fmt.Println(scanner.Text()) // print the line
	}

	err = file.Close() // close the file to free resources

	if err != nil {
		log.Fatal(err) // error closing the file, log and report it
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err()) // if error scanning file, report and exit
	}
}