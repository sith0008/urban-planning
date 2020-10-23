package api

import (
	. "github.com/sith0008/urban-planning/knowledge-graph/er"
)

type DBAccessor interface {
	GetCases() []Case
	UpsertCase(pastCase Case) (string, error)
	RemoveCase(caseId int64)
	GetSimilarCases(query Query) []Case
	UpsertRelation(nodeOne int64, nodeTwo int64, relationType RelationType)
}
