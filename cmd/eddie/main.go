package main

import (
	"flag"
	"github.com/iods/go-pherit/pkg/format"
)

type Application struct {

}

var (
	flagword string
)


func main() {

	flag.StringVar(&flagword, "word", "default-word", "This is the description of the flag. (Required)")
	flag.String("", "default-f1", "This is the description of the flag.")

	flag.Parse()

	format.RenderFlagSet(flag.CommandLine)
}
