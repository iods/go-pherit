package datatypes

import (
	"fmt"
	"strings"
)

// SlickTypes Returns an exportable function of string examples in Go.
func SlickTypes() {
	title := "Slick: String Types in Go."
	fmt.Println(len(title))
	fmt.Println(strings.Title(title))
	slickRawStringLiterals()
	slickConcats()
	slickEscapes()
	slickRunes()
}

// slickRawStringLiterals Returns examples of raw string literals in Go.
func slickRawStringLiterals() {
	// Strings and raw string literals are the same.
	// Both can be used as a string value.
	var s string

	s = "Slick was the boat we had in Michigan."
	s = `Slick was the boat we had in Michigan.`
	fmt.Println(s)

	s = "<html>\n\t<body>\n\t\t<h1>\"Slick\"</h2>\n\t</body>\n</html>"
	fmt.Println(s)

	s = `
<html>
	<body>
		<h1>"Slick"</h1>
	</body>
</html>`
	fmt.Println(s)

	fmt.Println("C:\\path\\to\\my\\file.txt")
	fmt.Println(`C:\path\to\my\file.txt`)
}

// slickConcats Returns examples of concatenating strings and non-strings in Go.
func slickConcats() {
	name, color := "Slick", "Blue"
	fmt.Println("The " + name + " was " + color)

	// adding a period using string concat w/ assignment operation
	// color = color + "."
	color += "."
	fmt.Println("The " + name + " was " + color)

	// works with raw string literals too
	fmt.Println(
		`The ` + name + ` was ` + color,
	)
}

// slickEscapes Returns various examples of Escape sequences in Go.
func slickEscapes() {
	fmt.Println("Escape Sequences")
	fmt.Println("The \nDark \nSociety")
	fmt.Println("The \tDark \tSociety")
	fmt.Println("\"The\" \"Dark\" \"Society\"")
	fmt.Println("C:\\path\\to\\my\\file.txt")
}

// slickRunes Returns examples of rune literals in Go.
func slickRunes() {
	fmt.Println('A')
	fmt.Println('B')
}
