package portage

import (
	"fmt"
	"strings"
)

// SlickStrings Returns an exportable function of string examples in Go.
func SlickStrings() {
	fmt.Println(strings.Title("Slick"))
	slickEscapes()
	slickRunes()
}

// slickEscapes Returns various examples of Escape sequences in Go.
func slickEscapes() {
	fmt.Println("Escape Sequences")
	fmt.Println("The \nDark \nSociety")
	fmt.Println("The \tDark \tSociety")
	fmt.Println("\"The\" \"Dark\" \"Society\"")
	fmt.Println("\\The \\Dark \\Society")
}

// slickRunes Returns examples of runes in Go.
func slickRunes() {
	fmt.Println('A')
	fmt.Println('B')
}