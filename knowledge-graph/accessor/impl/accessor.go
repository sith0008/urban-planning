package impl

import (
	"fmt"
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

// TODO: insert is done, handle updates
func (accessor *DBAccessorImpl) UpsertCase(c Case) (int64, error) {
	tx, err := accessor.DBSession.BeginTransaction()
	if err != nil {
		log.Println("[ERROR]: failed to start transaction")
		log.Fatal(err)
		return 0, err
	}

	insertCaseResult, err := tx.Run("CREATE (c:Case) SET c.case_id = %caseId, c.proposed_use = $proposedUse, c.GFA = $GFA, c.decision = $decision, c.evaluation = $evaluation RETURN id(c)",
		map[string]interface{}{
			"caseId":      c.Id,
			"proposedUse": c.ProposedUseDesc,
			"GFA":         c.GFA,
			"decision":    c.Decision,
			"evaluation":  c.Evaluation,
		})
	if err != nil {
		log.Println("[ERROR]: failed to insert case")
		log.Println(err)
		return 0, err
	}
	var caseId int64
	if insertCaseResult.Next() {
		caseId = insertCaseResult.Record().GetByIndex(0).(int64)
		log.Printf("[INFO] Inserted case: %d", caseId)
	}
	tx.Commit()
	return caseId, nil
}

// TODO: insert is done, handle updates
func (accessor *DBAccessorImpl) UpsertLocation(location Location) (int64, error) {
	tx, err := accessor.DBSession.BeginTransaction()
	if err != nil {
		log.Println("[ERROR]: failed to start transaction")
		log.Fatal(err)
		return 0, err
	}

	insertLocationResult, err := tx.Run("CREATE (l:Location) SET l.postalCode = $postalCode RETURN id(l)",
		map[string]interface{}{
			"postalCode": location.PostalCode,
		})
	if err != nil {
		log.Println("[ERROR]: failed to insert location")
		log.Println(err)
		return 0, err
	}
	var locationId int64
	if insertLocationResult.Next() {
		locationId = insertLocationResult.Record().GetByIndex(0).(int64)
		log.Printf("[INFO] Inserted location: %d", locationId)
	}
	tx.Commit()
	return locationId, nil
}

// TODO: insert is done, handle updates
func (accessor *DBAccessorImpl) UpsertCaseLocRelation(caseId int64, locationId int64) (int64, error) {
	tx, err := accessor.DBSession.BeginTransaction()
	if err != nil {
		log.Println("[ERROR]: failed to start transaction")
		log.Fatal(err)
		return 0, err
	}
	insertRelationResult, err := tx.Run("MATCH (c:Case),(l:Location) WHERE id(c) = $caseId AND id(l) = $locationId CREATE (c)-[r:LOCATED_IN]->(l) RETURN id(r)",
		map[string]interface{}{
			"caseId":     caseId,
			"locationId": locationId,
		})
	if err != nil {
		log.Println("[ERROR]: failed to insert LOCATED_IN relation")
		log.Println(err)
		return 0, err
	}
	var relationId int64
	if insertRelationResult.Next() {
		relationId = insertRelationResult.Record().GetByIndex(0).(int64)
		log.Printf("[INFO] Inserted LOCATED_IN relation: %d", relationId)

	} else {
		log.Println("[WARN] Cannot find match")
		tx.Rollback()
		return 0, err
	}
	tx.Commit()
	return relationId, nil
}

func (accessor *DBAccessorImpl) RemoveCase(id int64) error {
	tx, err := accessor.DBSession.BeginTransaction()
	if err != nil {
		log.Println("[ERROR]: failed to start transaction")
		log.Fatal(err)
		return err
	}
	log.Printf("[INFO]: Removing case id %d", id)
	removeCaseResult, err := tx.Run("MATCH (c:Case) WHERE id(c) = $caseId DETACH DELETE c",
		map[string]interface{}{
			"caseId": id,
		})
	if removeCaseResult.Next() {
		fmt.Println(removeCaseResult.Record().GetByIndex(0).(string))
	}
	tx.Commit()
	return nil
}
func (accessor *DBAccessorImpl) RemoveLocation(id int64) error {
	tx, err := accessor.DBSession.BeginTransaction()
	if err != nil {
		log.Println("[ERROR]: failed to start transaction")
		log.Fatal(err)
		return err
	}
	log.Printf("[INFO]: Removing location id %d", id)
	removeCaseResult, err := tx.Run("MATCH (l:Location) WHERE id(l) = $locationId DETACH DELETE l",
		map[string]interface{}{
			"locationId": id,
		})
	if removeCaseResult.Next() {
		fmt.Println(removeCaseResult.Record().GetByIndex(0).(string))
	}
	tx.Commit()
	return nil
}

func (accessor *DBAccessorImpl) ClearDatabase() error {
	tx, err := accessor.DBSession.BeginTransaction()
	if err != nil {
		log.Println("[ERROR]: failed to start transaction")
		log.Fatal(err)
		return err
	}
	_, err = tx.Run("MATCH (n) DETACH DELETE n", map[string]interface{}{})
	if err != nil {
		log.Fatal(err)
		return err
	}
	tx.Commit()
	return nil
}

func (accessor *DBAccessorImpl) UpsertCaseUseClassRelation(caseId int64, specifcUseClass SpecificUseClass) (int64, error) {
	tx, err := accessor.DBSession.BeginTransaction()
	if err != nil {
		log.Println("[ERROR]: failed to start transaction")
		log.Fatal(err)
		return 0, err
	}
	insertCaseUseClassRelationResult, err := tx.Run("MATCH (c:Case), (s:SpecificUseClass) WHERE id(c) = $caseId AND s.name = $proposedUseClass CREATE (c)-[r:HAS_USE_CLASS]->(s) RETURN id(r)",
		map[string]interface{}{
			"caseId":           caseId,
			"proposedUseClass": specifcUseClass,
		})
	if err != nil {
		log.Println("[ERROR]: failed to insert HAS_USE_CLASS relation")
		log.Println(err)
		return 0, err
	}
	var caseUseClassRelationId int64
	if insertCaseUseClassRelationResult.Next() {
		caseUseClassRelationId = insertCaseUseClassRelationResult.Record().GetByIndex(0).(int64)
		log.Printf("[INFO] Inserted case-uc relation: %d", caseUseClassRelationId)
	}
	tx.Commit()
	return caseUseClassRelationId, nil
}

func (accessor *DBAccessorImpl) UpsertLocationPropTypeRelation(locationId int64, specificPropType SpecificPropType) (int64, error) {
	tx, err := accessor.DBSession.BeginTransaction()
	if err != nil {
		log.Println("[ERROR]: failed to start transaction")
		log.Fatal(err)
		return 0, err
	}
	insertLocationPropTypeRelationResult, err := tx.Run("MATCH (l:Location), (s:SpecificPropType) WHERE id(l) = $locationId AND s.name = $propType CREATE (l)-[r:HAS_PROP_TYPE]->(s) RETURN id(r)",
		map[string]interface{}{
			"locationId": locationId,
			"propType":   specificPropType,
		})
	if err != nil {
		log.Println("[ERROR]: failed to insert HAS_PROP_TYPE relation")
		log.Println(err)
		return 0, err
	}
	var locationPropTypeRelationId int64
	if insertLocationPropTypeRelationResult.Next() {
		locationPropTypeRelationId = insertLocationPropTypeRelationResult.Record().GetByIndex(0).(int64)
		log.Printf("[INFO] Inserted location-pt relation: %d", locationPropTypeRelationId)
	}
	tx.Commit()
	return locationPropTypeRelationId, nil
}

// GetCases() []Case
// UpsertCase(pastCase Case) int64, error
// RemoveCase(caseId int64)
// UpsertLocation(location Location) int64
// GetSimilarCases(query Query) []Case
// UpsertRelation(relation Relation) int64
