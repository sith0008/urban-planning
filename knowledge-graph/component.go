package main

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	. "github.com/urban-planning/knowledge-graph/accessor/api"
	. "github.com/urban-planning/knowledge-graph/accessor/impl"
)

// TODO add neo4j configs
const (
	DB_Host     = "127.0.0.1"
	DB_Port     = "3306"
	DB_Username = "user"
	DB_Password = "password"
)

type KnowledgeGraphComponent struct {
	Accessor DBAccessor
}

func NewKnowledgeGraphComponent() (*KnowledgeGraphComponent, error) {
	c := KnowledgeGraphComponent{}
	// TODO: add open connection to DB
	_, dbSession, err := InitialiseDBSession(DB_Host, DB_Port, DB_Username, DB_Password)
	if err != nil {
		return nil, err
	}

	dbAccessor := NewDBAccessorImpl(dbSession)
	c.Accessor = dbAccessor
	return &c, nil
}

func InitialiseDBSession(dbHost string, dbPort string, dbUsername string, dbPassword string) (neo4j.Driver, neo4j.Session, error) {
	var (
		driver  neo4j.Driver
		session neo4j.Session
		err     error
	)

	if driver, err = neo4j.NewDriver(dbHost+"+"+dbPort, neo4j.BasicAuth(dbUsername, dbPassword, "")); err != nil {
		return nil, nil, err
	}
	if session, err = driver.Session(neo4j.AccessModeWrite); err != nil {
		return nil, nil, err
	}
	return driver, session, nil
}
