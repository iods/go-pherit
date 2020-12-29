/*
Exercise:

1. Print the count of the command-line arguments.
2. Print the path of the running program (get it from os.Args)
  * use go build then run it from the executable
3. Print your name
  * get it from the first argument
  * it should return "hello rye, how are you?"
4. Greet three people
5. Greet five people

Input:
bilbo balbo bungg
Rye
bilbo balbo bungg
bilbo balbo bungo gandalf legolas

Output:
There are 3 names.
myprogram
Hi, Rye!
How are you?
There are 3 people!
Hello great bilbo !
Hello great balbo !
Hello great bungo !
Nice to meet you all.
There are 5 people!
Hello great bilbo !
Hello great balbo !
Hello great bungo !
Hello great gandalf !
Hello great legolas !
Nice to meet you all.

Notes:
  * Use the fmt.Printf function

Bonus #1:
  * Observe the error if you pass less then 3 arguments.
*/
package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println(len(os.Args[1:]))

	fmt.Println(os.Args[0])

	name := os.Args[1]
	fmt.Printf("Hey, %s!\n", name)
	fmt.Println("How are you?")

	count := len(os.Args[1:])
	fmt.Printf("There are %d people.\n", count)
	fmt.Println("Hello great", os.Args[1], "!")
	fmt.Println("Hello great", os.Args[2], "!")
	fmt.Println("Hello great", os.Args[3], "!")
	fmt.Println("Nice to meet you all.")

	counts := os.Args[1:]
	fmt.Printf("There are %d people.\n", count)
	for _, num := range counts {
		fmt.Printf("Hello great %s !\n", num)
	}
	fmt.Println("Nice to meet you all.")
}