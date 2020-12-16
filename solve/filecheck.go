/*
Print the size of a file.

Print the os.Stat function that returns an os.FileInfo value
Size method on the FileInfo value to get the return value.
*/
package main

import (
	"fmt"
	"log"
	"os"
)

func main()  {
	fileInformation, err := os.Stat("filecheck.txt")
	// If the error is not nil, pass it to log.Fatal
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInformation.Size())
}
