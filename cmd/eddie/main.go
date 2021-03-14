package main

import (
	"flag"
	"os"

	"github.com/iods/go-pherit/internal/format"
)

type Application struct {

}

var (
	flagword string
)


func main() {
	// custom flagset (subcommands)

	//flag.StringVar(&flagword, "word", "default-word", "This is the description of the flag. (Required)")
	//flag.String("", "default-f1", "This is the description of the flag.")

	// flag.Parse()

	//format.RenderFlagSet(flag.CommandLine)

	custom := flag.NewFlagSet("track", flag.ExitOnError)
	custom.StringVar(&flagword,"sleep", "nightly", "Track sleep nightly.")

	custom.Parse(os.Args[1:])

	flag.Parse()

	format.RenderFlagSet(custom)
}
