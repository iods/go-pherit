package main


type Record struct {
	Name string `bson:"name"` // e.g. Sleep
	Type string `bson:"type"` // e.g.
}

type Routine struct {
	Records []Record `bson:"records"`
}