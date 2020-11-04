package er

type QueryRequest struct {
	ProposedUseClass SpecificUseClass
	ProposedUseDesc  string
	GFA              float64
	PostalCode       string
	LotNumber        string
	Floor            int64
	Unit             int64
}

type QueryResponse struct {
	CaseSpec                 Case
	LocationSpec             Location
	ProposedSpecificUseClass SpecificUseClass
	ProposedGenericUseClass  GenericUseClass
	SpecificPropertyType     SpecificPropType
	GenericPropertyType      GenericPropType
}
