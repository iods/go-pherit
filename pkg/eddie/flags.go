package eddie

import (
	"flag"
	"os"
)


type Car struct {
	Model []string // model of the car (i.e. Ford Mustang)
}


var (
	boolflag bool
	count    []string
	Flagname int
	help     bool
	Ip		 int
	Strflag  string
	Wordptr  string
)

func Init() {
	flag.StringVar(&Strflag, "t", "", "Description. (Required)")				  // eddie -t arg
	flag.BoolVar(&boolflag, "b", false, "Description of the flag.")			  // eddie -b description is big
	flag.IntVar(&Flagname, "flagname", 1234, "Help message for this flagname.")  // eddie -flagname
	flag.IntVar(&Ip, "ip", 127001, "Help message for this flagname.")			  // eddie -ip

	flag.StringVar(&Wordptr, "word", "foo", "Another word in the mix.")          // eddie -word hi

	flag.BoolVar(&help, "help", false, "Show this help message.")			      // eddie -help
	flag.BoolVar(&help, "h", false, "Show this help message.")					  // eddie -h


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