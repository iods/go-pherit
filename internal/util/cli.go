package util

import(
	"fmt"
	"os"
)

func main() {
	args := len(os.Args[1:]) // remove the project name
	fmt.Println(args)
}

