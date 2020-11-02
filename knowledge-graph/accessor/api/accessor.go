package api

import (
	. "github.com/sith0008/urban-planning/knowledge-graph/er"
)

type DBAccessor interface {
	GetCases() []Case
	UpsertCase(c Case) (int64, error)
	UpsertLocation(location Location) (int64, error)
	UpsertCaseLocRelation(caseId int64, locationId int64) (int64, error)
	RemoveCase(caseId int64) error
	RemoveLocation(locationId int64) error
	ClearDatabase() error
	// GetSimilarCases(query Query) []Case
	// UpsertRelation(nodeOne int64, nodeTwo int64, relationType RelationType)
	// RemoveRelation(nodeOne int64, nodeTwo int64, relationType RelationType)
}
