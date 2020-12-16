/*
Print the size of a file.
If the error is not nil, pass it to log.Fatal
Print the os.Stat function that returns an os.FileInfo value
Size method on the FileInfo value to get the return value.
*/
package main

import (
	"fmt"
	"os"
)

func main()  {
	fileInformation, _ := os.Stat("filecheck.txt")
	fmt.Println(fileInformation.Size())
}
