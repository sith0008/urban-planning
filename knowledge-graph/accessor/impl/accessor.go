package impl

import (
	"log"

	"github.com/neo4j/neo4j-go-driver/neo4j"
	. "github.com/sith0008/urban-planning/knowledge-graph/er"
)

type DBAccessorImpl struct {
	DBSession neo4j.Session
}

func NewDBAccessorImpl(dbSession neo4j.Session) *DBAccessorImpl {
	return &DBAccessorImpl{DBSession: dbSession}

}

func (accessor *DBAccessorImpl) GetCases() []Case {
	//tx, err := accessor.DBSession.BeginTransaction()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//tx.Run()
	return []Case{}
}

func (accessor *DBAccessorImpl) UpsertCase(pastCase Case) (string, error) {
	tx, err := accessor.DBSession.BeginTransaction()
	if err != nil {
		log.Println("[ERROR]: failed to start transaction")
		log.Fatal(err)
	}
	result, err := tx.Run("CREATE (c:Case) SET c.use = $use RETURN c.use + ', from node ' + id(c)",
		map[string]interface{}{
			"use": pastCase.ProposedUseDesc,
		})
	if err != nil {
		log.Println(err)
		return "", err
	}
	tx.Commit()
	if result.Next() {
		return result.Record().GetByIndex(0).(string), nil
	}

	return "", result.Err()
}

// GetCases() []Case
// UpsertCase(pastCase Case) int64
// RemoveCase(caseId int64)
// GetSimilarCases(query Query) []Case
// UpsertRelation(relation Relation) int64
