package main

import (
	"log"

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
