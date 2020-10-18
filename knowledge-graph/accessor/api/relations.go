package api

type Relation interface {
	GetType() RelationType
}

type RelationType int8

const (
	RelationType_HasA   RelationType = 1
	RelationType_PartOf RelationType = 2
)

type HasARelation struct {
}

func (r HasARelation) GetType() RelationType {
	return RelationType_HasA
}

type PartOfRelation struct {
}

func (r PartOfRelation) GetType() RelationType {
	return RelationType_PartOf
}
