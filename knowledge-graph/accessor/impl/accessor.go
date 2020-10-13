package impl

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type DBAccessorImpl struct {
	DBSession neo4j.Session
}

func NewDBAccessorImpl(dbSession neo4j.Session) *DBAccessorImpl {
	return &DBAccessorImpl{DBSession: dbSession}

}
