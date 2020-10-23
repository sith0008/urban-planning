package api

type Case struct {
	Id              int64
	ProposedUseDesc string
	GFA             float64
	Decision        string
	Evaluation      string
}

type Location struct {
	PostalCode int64
	LotNumber  string
	Floor      int64
	Unit       int64
}
type Query struct {
	ProposedUseClass ProposedUseClass
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

type ProposedUseClass int16

const (
	ProposedUseClass_Unknown ProposedUseClass = iota
	ProposedUseClass_Shop    ProposedUseClass = 1
	// TODO: add on to the list
)
