package eddie

import (
	"flag"
	"fmt"
)

// App Holds the flags for the application
type App struct {
	subject      string
	isAwesome    bool
	howAwesome   int
	countTheWays CountTheWays
}

func (a *App) Setup() {

	// we add a flag directly like so
	// var somevar = flag.String("flag_name", "default_value", "description")
	// but it is better to run it this way
	flag.StringVar(&a.subject,
		"subject",
		"",
		"subject is a string, it defaults to empty")
	flag.StringVar(&a.subject,
		"s",
		"",
		"subject is a string, it defaults to empty (shorthand)")
	flag.BoolVar(&a.isAwesome,
		"isawesome",
		false,
		"is it awesome or what?")
	flag.IntVar(&a.howAwesome,
		"howawesome",
		10,
		"how awesome is it out of 10?")

	flag.Var(&a.countTheWays,
		"c",
		"comma separated list of integers")
}

func (a *App) GetMessage() string {
	m := a.subject
	if a.isAwesome {
		m += " is awesome"
	} else {
		m += " is NOT awesome"
	}

	m = fmt.Sprintf("%s with a certainty of %d out of 10. Let me count the ways %s",
		m, a.howAwesome, a.countTheWays.String())
	return m
}