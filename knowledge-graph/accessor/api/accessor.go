package api

type DBAccessor interface {
	GetCases() []Case
	UpsertCase(pastCase Case) int64
	RemoveCase(caseId int64)
	GetSimilarCases(query Query) []Case
	UpsertRelation(relation Relation) int64
}
