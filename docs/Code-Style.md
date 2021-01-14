Code style and code standards of Go.

## File Layout
```go
// Package Clause
package main

// Import Statements
import "fmt"

// Function References
func main() {
    fmt.Println("Go Code Style")
}
```

### Package Clause

A package is a grouping of code that all does similar things. The **package clause** gives the name of the package that this files code becomes part of. The package `main` is required if the code is to be run directly, i.e. an executable, and from the terminal.

* Every file must begin with a Package Clause
* A package is a group of related functions and other code
* Before you can use a package's functions, you have to import that package
* A package is essentially source level coupling

### Import Statements

Each Go file must import other packages in order to use their resources. Rather than loading all packages, the import statement allows for explicit declaration of packages for your Go file.

* Every Go file must import each Package it references
* Only import Packages that your Go file references

### Functions

A function is one or more lines of code you call or run from various places of a program. When calling a function, using arguments to provide the function with data will change it's output.

With Go, the function `main` is executed first.

* Go looks for the function `main` to run first


## Naming Conventions

There are very few, and minimal rules when it comes to naming in Go. They are:
* Variables, functions, and types must start with a letter
* They should be written in camelCase
* lowercase functions are not **exported**
* Functions that start with Uppercase letters **ARE** exported

Only variables, functions, and types whose names begin with a capital letter are considered to be exported, meaning accessible outside of the current package.

## FAQs

**What about semi-colons in Go?**

You _can_ use them to separate statements in Go; however, not only is it not required, it is generally frowned upon.

**What is this `go fmt` thing?**

The Go Compiler comes with a standard formatting tool. This tool can be run locally and will automatically fix all formatting.