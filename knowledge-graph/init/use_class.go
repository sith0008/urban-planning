package main

type SpecificUseClass string

const (
	SpecificUseClass_Restaurant SpecificUseClass = "Restaurant"
)

var SpecificUseClasses = []SpecificUseClass{
	SpecificUseClass_Restaurant,
}

type GenericUseClass string

const (
	GenericUseClass_FoodBev GenericUseClass = "FoodBEv"
)

var GenericUseClasses = []GenericUseClass{
	GenericUseClass_FoodBev,
}
