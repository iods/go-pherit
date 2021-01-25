package eddie

import (
	"flag"
	"os"
)

type Car struct {
	Model []string // model of the car (i.e. Ford Mustang)
}

var Flagname int
var help bool
var count []string





// var StrFlag = flag.String("text", "", "Description. (Required)")
// var BoolFlag = flag.Bool("bool", false, "Description of the flag.")

func Init() {
	// flag.StringVar(StrFlag, "t", "", "Description. (Required)")
	// flag.BoolVar(BoolFlag, "b", false, "Description of the flag.")

	flag.IntVar(&Flagname, "flagname", 1234, "Help message for this flagname.") // eddie -flagname
	flag.BoolVar(&help, "help", false, "Show this help message.")
	flag.BoolVar(&help, "h", false, "Show this help message.")
	flag.Parse()

	// some error handle if no args are provided
	if help {
		flag.PrintDefaults()
	}

	count = os.Args
	if len(count) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}
}