package impl

import (
	"errors"
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
	insertCaseResult, err := tx.Run("CREATE (c:Case) SET c.case_id = $caseId, c.proposed_use = $proposedUse, c.GFA = $GFA, c.decision = $decision, c.evaluation = $evaluation RETURN id(c)",
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
		tx.Close()
		return 0, err
	}
	var caseId int64
	if insertCaseResult.Next() {
		caseId = insertCaseResult.Record().GetByIndex(0).(int64)
		log.Printf("[INFO] Inserted case: %d", caseId)
	}
	tx.Commit()
	tx.Close()
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

	insertLocationResult, err := tx.Run("CREATE (l:Location) SET l.postalCode = $postalCode, l.floor = $floor, l.unit = $unit RETURN id(l)",
		map[string]interface{}{
			"postalCode": location.PostalCode,
			"floor":      location.Floor,
			"unit":       location.Unit,
		})
	if err != nil {
		log.Println("[ERROR]: failed to insert location")
		log.Println(err)
		tx.Close()
		return 0, err
	}
	var locationId int64
	if insertLocationResult.Next() {
		locationId = insertLocationResult.Record().GetByIndex(0).(int64)
		log.Printf("[INFO] Inserted location: %d", locationId)
	}
	tx.Commit()
	tx.Close()
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
		tx.Close()
		return 0, err
	}
	var relationId int64
	if insertRelationResult.Next() {
		relationId = insertRelationResult.Record().GetByIndex(0).(int64)
		log.Printf("[INFO] Inserted LOCATED_IN relation: %d", relationId)

	} else {
		log.Println("[WARN] Cannot find match")
		tx.Rollback()
		tx.Close()
		return 0, err
	}
	tx.Commit()
	tx.Close()
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
	tx.Close()
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
	tx.Close()
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
	tx.Close()
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
		tx.Close()
		return 0, err
	}
	var caseUseClassRelationId int64
	if insertCaseUseClassRelationResult.Next() {
		caseUseClassRelationId = insertCaseUseClassRelationResult.Record().GetByIndex(0).(int64)
		log.Printf("[INFO] Inserted case-uc relation: %d", caseUseClassRelationId)
	}
	tx.Commit()
	tx.Close()
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
		tx.Close()
		return 0, err
	}
	var locationPropTypeRelationId int64
	if insertLocationPropTypeRelationResult.Next() {
		locationPropTypeRelationId = insertLocationPropTypeRelationResult.Record().GetByIndex(0).(int64)
		log.Printf("[INFO] Inserted location-pt relation: %d", locationPropTypeRelationId)
	}
	tx.Commit()
	tx.Close()
	return locationPropTypeRelationId, nil
}

func (accessor *DBAccessorImpl) GetSimilarCases(query QueryRequest) ([]QueryResponse, error) {
	tx, err := accessor.DBSession.BeginTransaction()
	if err != nil {
		log.Println("[ERROR]: failed to start transaction")
		log.Fatal(err)
		return []QueryResponse{}, err
	}
	queryResponses := make([]QueryResponse, 0)
	getPropTypeResult, err := tx.Run("MATCH (l:Location)--(p:SpecificPropType)--(gp:GenericPropType) WHERE l.postalCode = $postalCode AND l.floor = $floor AND l.unit = $unit RETURN p.name, gp.name",
		map[string]interface{}{
			"postalCode": query.PostalCode,
			"floor":      query.Floor,
			"unit":       query.Unit,
		})
	if err != nil {
		log.Println("[ERROR]: failed to retrieve propType")
		log.Println(err)
		tx.Close()
		return []QueryResponse{}, err
	}
	var specificPropType SpecificPropType
	var genericPropType GenericPropType
	if getPropTypeResult.Next() {
		specificPropType = SpecificPropType(getPropTypeResult.Record().GetByIndex(0).(string))
		genericPropType = GenericPropType(getPropTypeResult.Record().GetByIndex(1).(string))
	} else {
		log.Printf("[ERROR] Location with postal code: %s, floor: %d, unit: %d does not exist", query.PostalCode, query.Floor, query.Unit)
		err = errors.New("LocationError")
		tx.Close()
		return []QueryResponse{}, err
	}

	getGenericUseClassResult, err := tx.Run("MATCH (su:SpecificUseClass)--(gu:GenericUseClass) WHERE su.name = $specificUseClass RETURN gu.name",
		map[string]interface{}{
			"specificUseClass": query.ProposedUseClass,
		})
	if err != nil {
		log.Println("[ERROR]: failed to retrieve genericUseClass")
		log.Println(err)
		tx.Close()
		return []QueryResponse{}, err
	}
	var genericUseClass GenericUseClass
	if getGenericUseClassResult.Next() {
		genericUseClass = GenericUseClass(getGenericUseClassResult.Record().GetByIndex(0).(string))
	}

	// get similar cases
	statement := `
	MATCH (c:Case)--(l:Location)--(sp:SpecificPropType)--(gp:GenericPropType), 
	(c:Case)--(su:SpecificUseClass)--(gu:GenericUseClass) 
	WHERE (sp.name = $specificPropType and su.name = $specificUseClass) 
	OR (sp.name = $specificPropType and gu.name = $genericUseClass)
	OR (gp.name = $genericPropType and su.name = $specificUseClass)
	OR (gp.name = $genericPropType and gu.name = $genericUseClass)
	RETURN c.case_id, c.proposed_use, c.decision, c.evaluation, c.GFA, 
	l.postalCode, l.floor, l.unit, sp.name, gp.name, su.name, gu.name
	`
	getSimilarCasesResult, err := tx.Run(statement,
		map[string]interface{}{
			"specificPropType": specificPropType,
			"genericPropType":  genericPropType,
			"specificUseClass": query.ProposedUseClass,
			"genericUseClass":  genericUseClass,
		})
	if err != nil {
		log.Println("[ERROR]: failed to retrieve similar cases")
		log.Println(err)
		tx.Close()
		return []QueryResponse{}, err
	}
	for getSimilarCasesResult.Next() {
		caseId := getSimilarCasesResult.Record().GetByIndex(0).(string)
		caseProposedUseDesc := getSimilarCasesResult.Record().GetByIndex(1).(string)
		caseDecision := getSimilarCasesResult.Record().GetByIndex(2).(string)
		caseEvaluation := getSimilarCasesResult.Record().GetByIndex(3).(string)
		caseGFA := getSimilarCasesResult.Record().GetByIndex(4).(float64)
		casePostalCode := getSimilarCasesResult.Record().GetByIndex(5).(string)
		caseFloor := getSimilarCasesResult.Record().GetByIndex(6).(int64)
		caseUnit := getSimilarCasesResult.Record().GetByIndex(7).(int64)
		caseSP := SpecificPropType(getSimilarCasesResult.Record().GetByIndex(8).(string))
		caseGP := GenericPropType(getSimilarCasesResult.Record().GetByIndex(9).(string))
		caseSU := SpecificUseClass(getSimilarCasesResult.Record().GetByIndex(10).(string))
		caseGU := GenericUseClass(getSimilarCasesResult.Record().GetByIndex(11).(string))

		queryResponses = append(queryResponses, QueryResponse{
			CaseSpec: Case{
				Id:              caseId,
				ProposedUseDesc: caseProposedUseDesc,
				Decision:        caseDecision,
				Evaluation:      caseEvaluation,
				GFA:             caseGFA,
			},
			LocationSpec: Location{
				PostalCode: casePostalCode,
				Floor:      caseFloor,
				Unit:       caseUnit,
			},
			SpecificPropertyType: caseSP,
			GenericPropertyType:  caseGP,
			SpecificUseClass:     caseSU,
			GenericUseClass:      caseGU,
		})
	}
	tx.Commit()
	tx.Close()
	return queryResponses, nil

}

// GetSimilarCases(query Query) []Case
// UpsertRelation(relation Relation) int64
