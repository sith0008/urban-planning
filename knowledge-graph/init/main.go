package main

import (
	"log"

	"github.com/neo4j/neo4j-go-driver/neo4j"
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
	err = InitialiseUseClass(dbSession)
	if err != nil {
		// TODO: handle error
	}
	err = InitialisePropType(dbSession)
	if err != nil {
		// TODO: handle error
	}
}

func InitialiseUseClass(dbSession neo4j.Session) error {
	tx, err := dbSession.BeginTransaction()
	result, err := tx.Run("CREATE (c:Case) SET c.use = $use RETURN c.use + ', from node ' + id(c)",
		map[string]interface{}{})
	if err != nil {
		// TOOD: handle error
		log.Println(err)
	}
	tx.Commit()
	if result.Next() {
		log.Println(result.Record().GetByIndex(0).(string))
	}
	return nil
}

func InitialisePropType(dbSession neo4j.Session) error {
	tx, err := dbSession.BeginTransaction()
	result, err := tx.Run("CREATE (c:Case) SET c.use = $use RETURN c.use + ', from node ' + id(c)",
		map[string]interface{}{})
	if err != nil {
		// TOOD: handle error
		log.Println(err)
	}
	tx.Commit()
	if result.Next() {
		log.Println(result.Record().GetByIndex(0).(string))
	}
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
