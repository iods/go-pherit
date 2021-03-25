package main

import "time"

type Behavior struct {
	//Prepare
	///Introduce
	//Practice
	//Maintain

	//Resolve, Rehearse, Repeat
}

type Reward struct {

}
type Reminder struct {

	Name string
	Description string
	Time time.Time
	RoutineID uint
	Routine Routine
}

type Routines struct {

	Name string
	Description string
	Type string
}

type Report struct {
	Name string
	Description string
	RoutineID uint
	Routine Routine
}