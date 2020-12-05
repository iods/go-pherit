package portage

import (
	"fmt"
	"runtime"
)

// Slick Returns the name of the boat we had in Portage.
func Slick() {
	fmt.Println("Slick was the name of our boat.")
}

// Version Returns the version of Go on our OS.
func Version() string {
	return runtime.Version()
}

// Cpus Returns the amount of CPUs running on your system.
func Cpus() int{
	return runtime.NumCPU()
}