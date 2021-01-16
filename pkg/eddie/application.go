package eddie

// eddie MGT-COMMAND COMMAND [--options]
// eddie track sleep --nightly (make an entry of sleep, with that time, and record with others)
// eddie report sleep --daily (display a log of my average sleep a day)
// eddie ask hunger --weekly (ask me at a set time weekly, or when im active how my habit is)
// eddie ask fortune (easter egg of all my chinese fortune cookies) -> numbers and a fortune

// Application The default structure of Eddie or any CLI application.
type Application struct {
	Name        string   // name of the cli application (in this case Eddie)
	Version     string   // a semver string (not float) of the version of Eddie (or a cli program)
	Args		[]string // default is equivalent to os.Args
	Errors		[]error  // store some runtime errors
}

//      ______    __    ___
//     / ____/___/ /___/ (_)__
//    / __/ / __  / __  / / _ \
//   / /___/ /_/ / /_/ / /  __/
//  /_____/\__,_/\__,_/_/\___/
//