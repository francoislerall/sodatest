package main

type Profile struct {
    Id string
    Name   string
    Modules []Module
	Indicators []Indicator
}

type Module struct {
	Index int
	Name string
}

type Indicator struct {
	Name string
	Type string
	Group string
	Code string
	Unit string
	UUID string
	UnitGroupUUID string
}
