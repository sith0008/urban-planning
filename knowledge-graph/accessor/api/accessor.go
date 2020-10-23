package api

import (
	. "github.com/sith0008/urban-planning/knowledge-graph/entity"
)

type DBAccessor interface {
	GetCases() []Case
	UpsertCase(pastCase Case) (string, error)
	RemoveCase(caseId int64)
	GetSimilarCases(query Query) []Case
}
