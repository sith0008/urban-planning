package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/neo4j/neo4j-go-driver/neo4j"
	. "github.com/sith0008/urban-planning/knowledge-graph/er"
)

// TODO add neo4j configs
const (
	DB_Host     = "localhost"
	DB_Port     = "7687"
	DB_Username = "neo4j"
	DB_Password = "password"
)

type ProcessedCaseLocation struct {
	Case     ProcessedCase     `json: case`
	Location ProcessedLocation `json: location`
}

type ProcessedCase struct {
	Id               string           `json: id`
	ProposedUseClass SpecificUseClass `json: proposedUseClass`
	ProposedUseDesc  string           `json: proposedUseDesc`
	GFA              float64          `json: gfa`
	Decision         string           `json: decision`
	Evaluation       string           `json: evaluation`
}

type ProcessedLocation struct {
	PostalCode   float64          `json: postalCode`
	Lot          string           `json: lot`
	Floor        int64            `json: floor`
	Unit         int64            `json: unit`
	PropertyType SpecificPropType `json: propertyType`
}

func main() {
	_, dbSession, err := InitialiseDBSession(DB_Host, DB_Port, DB_Username, DB_Password)
	if err != nil {
		// TODO: handle error
	}
	err = ClearDatabase(dbSession)
	err = InitialiseUseClass(dbSession, UseClassMap)
	if err != nil {
		// TODO: handle error
	}
	err = InitialisePropType(dbSession, PropTypeMap)
	if err != nil {
		// TODO: handle error
	}
	var cases []ProcessedCaseLocation
	jsonFile, err := os.Open("processed-samples.json")
	defer jsonFile.Close()
	casesJson, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(casesJson), &cases)
	err = InitialiseSampleData(dbSession, cases)
}

func InitialiseSampleData(dbSession neo4j.Session, cases []ProcessedCaseLocation) error {
	tx, err := dbSession.BeginTransaction()
	if err != nil {
		log.Println("[ERROR]: failed to start transaction")
		log.Fatal(err)
		return err
	}
	for _, c := range cases {
		// insert case
		insertCaseResult, err := tx.Run("CREATE (c:Case) SET c.proposed_use = $proposedUse, c.GFA = $GFA, c.decision = $decision, c.evaluation = $evaluation RETURN id(c)",
			map[string]interface{}{
				"proposedUse": c.Case.ProposedUseDesc,
				"GFA":         c.Case.GFA,
				"decision":    c.Case.Decision,
				"evaluation":  c.Case.Evaluation,
			})
		if err != nil {
			log.Println("[ERROR]: failed to insert case")
			log.Println(err)
			return err
		}
		var caseId int64
		if insertCaseResult.Next() {
			caseId = insertCaseResult.Record().GetByIndex(0).(int64)
			log.Printf("[INFO] Inserted case: %d", caseId)
		}
		// insert case-uc relation
		insertCaseUseClassRelationResult, err := tx.Run("MATCH (c:Case), (s:SpecificUseClass) WHERE id(c) = $caseId AND s.name = $proposedUseClass CREATE (c)-[r:HAS_USE_CLASS]->(s) RETURN id(r)",
			map[string]interface{}{
				"caseId":           caseId,
				"proposedUseClass": c.Case.ProposedUseClass,
			})
		if err != nil {
			log.Println("[ERROR]: failed to insert HAS_USE_CLASS relation")
			log.Println(err)
			return err
		}
		var caseUseClassRelationId int64
		if insertCaseResult.Next() {
			caseUseClassRelationId = insertCaseUseClassRelationResult.Record().GetByIndex(0).(int64)
			log.Printf("[INFO] Inserted case-uc relation: %d", caseUseClassRelationId)
		}
		// insert location
		var locationId int64
		getLocationResult, err := tx.Run("MATCH (l:location) WHERE l.postalCode = $postalCode AND l.floor = $floor AND l.unit = $unit RETURN id(l)",
			map[string]interface{}{
				"postalCode": c.Location.PostalCode,
				"floor":      c.Location.Floor,
				"unit":       c.Location.Unit,
			})
		if getLocationResult.Next() {
			locationId = getLocationResult.Record().GetByIndex(0).(int64)
			log.Printf("[INFO] Retrieved location: %d", locationId)
		} else {
			insertLocationResult, err := tx.Run("CREATE (l:Location) SET l.postalCode = $postalCode, l.floor = $floor, l.unit = $unit RETURN id(l)",
				map[string]interface{}{
					"postalCode": c.Location.PostalCode,
					"floor":      c.Location.Floor,
					"unit":       c.Location.Unit,
				})
			if err != nil {
				log.Println("[ERROR]: failed to insert location")
				log.Println(err)
				return err
			}
			if insertLocationResult.Next() {
				locationId = insertLocationResult.Record().GetByIndex(0).(int64)
				log.Printf("[INFO] Inserted location: %d", locationId)
			}
		}
		// insert location-proptype relation
		insertLocationPropTypeRelationResult, err := tx.Run("MATCH (l:Location), (s:SpecificPropType) WHERE id(l) = $locationId AND s.name = $propType CREATE (l)-[r:HAS_PROP_TYPE]->(s) RETURN id(r)",
			map[string]interface{}{
				"locationId": locationId,
				"propType":   c.Location.PropertyType,
			})
		if err != nil {
			log.Println("[ERROR]: failed to insert HAS_PROP_TYPE relation")
			log.Println(err)
			return err
		}
		var locationPropTypeRelationId int64
		if insertCaseResult.Next() {
			locationPropTypeRelationId = insertLocationPropTypeRelationResult.Record().GetByIndex(0).(int64)
			log.Printf("[INFO] Inserted location-pt relation: %d", locationPropTypeRelationId)
		}
		// insert case-location relation
		insertRelationResult, err := tx.Run("MATCH (c:Case),(l:Location) WHERE id(c) = $caseId AND id(l) = $locationId CREATE (c)-[r:LOCATED_IN]->(l) RETURN id(r)",
			map[string]interface{}{
				"caseId":     caseId,
				"locationId": locationId,
			})
		if err != nil {
			log.Println("[ERROR]: failed to insert LOCATED_IN relation")
			log.Println(err)
			return err
		}
		var relationId int64
		if insertRelationResult.Next() {
			relationId = insertRelationResult.Record().GetByIndex(0).(int64)
			log.Printf("[INFO] Inserted LOCATED_IN relation: %d", relationId)

		} else {
			log.Println("[WARN] Cannot find match")
			tx.Rollback()
			return err
		}

	}

	tx.Commit()
	return nil
}

func InitialiseUseClass(dbSession neo4j.Session, useClassMap map[SpecificUseClass]GenericUseClass) error {
	tx, err := dbSession.BeginTransaction()
	if err != nil {
		log.Println("[ERROR]: failed to start transaction")
		log.Fatal(err)
		return err
	}
	for specific, generic := range useClassMap {
		insertSpecificResult, err := tx.Run("CREATE (s:SpecificUseClass) SET s.name = $name RETURN id(s)",
			map[string]interface{}{
				"name": specific,
			})
		if err != nil {
			log.Printf("[ERROR]: failed to insert %s", specific)
			return err
		}
		var specificId int64
		if insertSpecificResult.Next() {
			specificId = insertSpecificResult.Record().GetByIndex(0).(int64)
			log.Printf("[INFO] Inserted specific use class: %d", specificId)
		}
		getGenericResult, err := tx.Run("MATCH (g:GenericUseClass) WHERE g.name = $name RETURN id(g)",
			map[string]interface{}{
				"name": generic,
			})
		if err != nil {
			return err
		}
		var genericId int64
		if getGenericResult.Next() {
			genericId = getGenericResult.Record().GetByIndex(0).(int64)
			log.Printf("[INFO] Retrieved generic use class: %d", genericId)
		} else {
			insertGenericResult, err := tx.Run("CREATE (g:GenericUseClass) SET g.name = $name RETURN id(g)",
				map[string]interface{}{
					"name": generic,
				})
			if err != nil {
				log.Printf("[ERROR]: failed to insert %s", generic)
				return err
			}

			if insertGenericResult.Next() {
				genericId = insertGenericResult.Record().GetByIndex(0).(int64)
				log.Printf("[INFO] Inserted generic use class: %d", genericId)
			}
		}
		insertRelationResult, err := tx.Run("MATCH (s:SpecificUseClass),(g:GenericUseClass) WHERE id(s) = $specificId AND id(g) = $genericId CREATE (s)-[r:IS_A]->(g) RETURN id(r)",
			map[string]interface{}{
				"specificId": specificId,
				"genericId":  genericId,
			})
		if err != nil {
			log.Println("[ERROR]: failed to insert IS_A relation")
			log.Println(err)
			return err
		}
		var relationId int64
		if insertRelationResult.Next() {
			relationId = insertRelationResult.Record().GetByIndex(0).(int64)
			log.Printf("[INFO] Inserted IS_A relation: %d", relationId)

		} else {
			log.Println("[WARN] Cannot find match")
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

func InitialisePropType(dbSession neo4j.Session, propTypeMap map[SpecificPropType]GenericPropType) error {
	tx, err := dbSession.BeginTransaction()
	if err != nil {
		log.Println("[ERROR]: failed to start transaction")
		log.Fatal(err)
		return err
	}
	for specific, generic := range propTypeMap {
		insertSpecificResult, err := tx.Run("CREATE (s:SpecificPropType) SET s.name = $name RETURN id(s)",
			map[string]interface{}{
				"name": specific,
			})
		if err != nil {
			log.Printf("[ERROR]: failed to insert %s", specific)
			return err
		}
		var specificId int64
		if insertSpecificResult.Next() {
			specificId = insertSpecificResult.Record().GetByIndex(0).(int64)
			log.Printf("[INFO] Inserted specific property type: %d", specificId)
		}
		getGenericResult, err := tx.Run("MATCH (g:GenericPropType) WHERE g.name = $name RETURN id(g)",
			map[string]interface{}{
				"name": generic,
			})
		if err != nil {
			return err
		}
		var genericId int64
		if getGenericResult.Next() {
			genericId = getGenericResult.Record().GetByIndex(0).(int64)
			log.Printf("[INFO] Retrieved generic property type: %d", genericId)
		} else {
			insertGenericResult, err := tx.Run("CREATE (g:GenericPropType) SET g.name = $name RETURN id(g)",
				map[string]interface{}{
					"name": generic,
				})
			if err != nil {
				log.Printf("[ERROR]: failed to insert %s", generic)
				return err
			}

			if insertGenericResult.Next() {
				genericId = insertGenericResult.Record().GetByIndex(0).(int64)
				log.Printf("[INFO] Inserted generic property type: %d", genericId)
			}
		}
		insertRelationResult, err := tx.Run("MATCH (s:SpecificPropType),(g:GenericPropType) WHERE id(s) = $specificId AND id(g) = $genericId CREATE (s)-[r:IS_A]->(g) RETURN id(r)",
			map[string]interface{}{
				"specificId": specificId,
				"genericId":  genericId,
			})
		if err != nil {
			log.Println("[ERROR]: failed to insert IS_A relation")
			log.Println(err)
			return err
		}
		var relationId int64
		if insertRelationResult.Next() {
			relationId = insertRelationResult.Record().GetByIndex(0).(int64)
			log.Printf("[INFO] Inserted IS_A relation: %d", relationId)

		} else {
			log.Println("[WARN] Cannot find match")
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

func ClearDatabase(dbSession neo4j.Session) error {
	tx, err := dbSession.BeginTransaction()
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

func InitialiseDBSession(dbHost string, dbPort string, dbUsername string, dbPassword string) (neo4j.Driver, neo4j.Session, error) {
	var (
		driver  neo4j.Driver
		session neo4j.Session
		err     error
	)

	if driver, err = neo4j.NewDriver("bolt://"+dbHost+":"+dbPort, neo4j.BasicAuth(dbUsername, dbPassword, "")); err != nil {
		log.Fatalln("[ERROR]: Connection failure")
		return nil, nil, err
	}
	if session, err = driver.Session(neo4j.AccessModeWrite); err != nil {
		log.Fatalln("[ERROR]: Session error")
		return nil, nil, err
	}
	return driver, session, nil
}
