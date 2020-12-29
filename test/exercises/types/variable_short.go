/*
Exercise: Short declarations

1. Declare and print four variables using the short declaration statement.
2. Declare two variables using multiple short declaration
3. Declare two variables using multiple short declaration
  * the first variable (c) should have a value of 42
  * the second variable (d) should have a value of "good"
4. Declare a variable named `sum` using short declaration
  * initialize it with an expression by adding 27 and 3.5
5. Declare two variables using multiple short declaration
  * init both variables to true
  * discard the 2nd variable's value using a blank-identifier
  * print the first variable only

Output:
1. i: 314 f: 3.14 s: Hello b: true
2. 14 true
3. 42 good
4. 30.5
5. true

Notes:

*/
package main

import "fmt"

func main() {
	a, b := 14, true
	c, d := 42, "good"
	i := 314
	f := 3.14
	s := "Hello"

	sum := 27 + 3.5

	on, _ := true, true

	fmt.Println("i:", i, "f:", f, "s:", s, "b:", b)
	fmt.Println(a, b)
	fmt.Println(c, d)
	fmt.Println(sum)
	fmt.Println(on)
}
