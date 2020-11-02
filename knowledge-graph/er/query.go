package er

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

type QueryResponse struct {
	CaseSpec                 Case
	LocationSpec             Location
	ProposedSpecificUseClass SpecificUseClass
	ProposedGenericUseClass  GenericUseClass
	SpecificPropertyType     SpecificPropType
	GenericPropertyType      GenericPropType
}
