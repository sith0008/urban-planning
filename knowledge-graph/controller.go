package main

import (
	"encoding/json"
	. "github.com/urban-planning/knowledge-graph/accessor/api"
	"log"
	"net/http"
)

// router.HandleFunc("/getCases", c.GetCases)
// router.HandleFunc("/upsertCase", c.InsertNewCase)
// router.HandleFunc("/removeCase", c.RemoveCase)
// router.HandleFunc("/getSimilarCases", c.GetSimilarCases)

func (c *KnowledgeGraphComponent) GetCases(w http.ResponseWriter, r *http.Request) {

}

func (c *KnowledgeGraphComponent) UpsertCase(w http.ResponseWriter, r *http.Request) {
	var req map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		// TODO: handle error
	}
	pastCase := Case{
		Use: req["use"].(string),
	}
	result, err := c.Accessor.UpsertCase(pastCase)
	if err != nil {
		// TODO: handle error
		log.Println(err)
	}
	json.NewEncoder(w).Encode(&result)
}

func (c *KnowledgeGraphComponent) RemoveCase(w http.ResponseWriter, r *http.Request) {

}

func (c *KnowledgeGraphComponent) GetSimilarCases(w http.ResponseWriter, r *http.Request) {

}
