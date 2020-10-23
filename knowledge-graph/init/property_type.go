package main

type SpecificPropType string

const (
	SpecificPropType_Business1 SpecificPropType = "Business1"
)

var SpecificPropTypes = []SpecificPropType{
	SpecificPropType_Business1,
}

type GenericPropType string

const (
	GenericPropType_Industrial GenericPropType = "Industrial"
)

var GenericPropTypes = []GenericPropType{
	GenericPropType_Industrial,
}
