/*
Go-Pherit: A Learning Experience and personal SDK in Golang
Copyright (c) 2020 Rye Miller
https://darksociety.io
*/

package main

/*
Exercise: Declare variables (w/ bits)

1. Declare a few variables using the following types:
  * int
  * int8
  * int16
  * int32
  * int64
  * float32
  * float64
  * complex64
  * complex128
  * bool
  * string
  * rune
  * byte

Output:
1. 0 0 0 0 0 0 0 (0+0i) (0+0i) false 0 0

Notes:

*/

import "fmt"

func main() {

	// integer types
	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64

	// float types
	var f32 float32
	var f64 float64

	// complex types
	var c64 complex64
	var c128 complex128

	// bool type
	var b bool

	// string types
	var s string
	var r rune
	var by byte

	fmt.Println(
		i, i8, i16, i32, i64,
		f32, f64,
		c64, c128,
		b, r, by,
	)

	fmt.Printf("%q\n", s)
}