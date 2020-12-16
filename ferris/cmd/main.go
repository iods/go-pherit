package main

/*
Package ferris includes the logic to report on whether
  or not a grade is passing or failing.
*/
import "github.com/thedarksociety/go-pherit/ferris"



func main()  {
	// Everything to the end of the line is a comment
	// Additional lines require another `//` marker

	/* Everything until closing marker is comment
	   This includes new lines

	   And blank lines */

	/* Multiline comments can be single lined too */
	ferris.PassFail() // Comments can start anywhere on a line
}