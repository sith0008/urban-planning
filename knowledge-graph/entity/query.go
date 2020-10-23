package entity

type Query struct {
	ProposedUseClass SpecificUseClass
	ProposedUseDesc  string
	GFA              float64
	Location         LocationSpec
}

type LocationSpec struct {
	PostalCode int64
	LotNumber  string
	Floor      int64
	Unit       int64
}
