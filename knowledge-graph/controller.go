package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	. "github.com/sith0008/urban-planning/knowledge-graph/er"
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
		Id:              req["caseId"].(string),
		ProposedUseDesc: req["proposedUseDesc"].(string),
		GFA:             req["GFA"].(float64),
		Decision:        req["decision"].(string),
		Evaluation:      req["evaluation"].(string),
	}
	location := Location{
		PostalCode: req["postalCode"].(string),
		LotNumber:  req["lotNumber"].(string),
		Floor:      int64(req["floor"].(float64)),
		Unit:       int64(req["unit"].(float64)),
	}
	specificUseClass := SpecificUseClass(req["proposedUseClass"].(string))
	specificPropType := SpecificPropType(req["propertyType"].(string))

	caseId, err := c.Accessor.UpsertCase(pastCase)
	if err != nil {
		// TODO: handle error
		log.Println(err)
	}
	locationId, err := c.Accessor.UpsertLocation(location)
	if err != nil {
		// TODO: handle error
		log.Println(err)
	}
	_, err = c.Accessor.UpsertCaseLocRelation(caseId, locationId)
	if err != nil {
		// TODO: handle error
		log.Println(err)
	}
	_, err = c.Accessor.UpsertCaseUseClassRelation(caseId, specificUseClass)
	if err != nil {
		// TODO: handle error
		log.Println(err)
	}
	_, err = c.Accessor.UpsertLocationPropTypeRelation(locationId, specificPropType)
	if err != nil {
		// TODO: handle error
		log.Println(err)
	}
	json.NewEncoder(w).Encode(&caseId)
}

func (c *KnowledgeGraphComponent) RemoveCase(w http.ResponseWriter, r *http.Request) {
	var req map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		// TODO: handle error
		log.Fatal(err)
	}
	caseId := int64(req["id"].(float64))
	err = c.Accessor.RemoveCase(caseId)
	if err != nil {
		// TODO: handle error
		log.Fatal(err)
	}
	var retString string = fmt.Sprintf("successfully deleted case id %d", caseId)
	json.NewEncoder(w).Encode(&retString)
}

func (c *KnowledgeGraphComponent) RemoveLocation(w http.ResponseWriter, r *http.Request) {
	var req map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		// TODO: handle error
		log.Fatal(err)
	}
	locationId := int64(req["id"].(float64))
	err = c.Accessor.RemoveLocation(locationId)
	if err != nil {
		// TODO: handle error
		log.Fatal(err)
	}
	var retString string = fmt.Sprintf("successfully deleted location id %d", locationId)
	json.NewEncoder(w).Encode(&retString)
}

func (c *KnowledgeGraphComponent) GetSimilarCases(w http.ResponseWriter, r *http.Request) {
	var req map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		// TODO: handle error
		log.Printf("[ERROR]: %s", err)
	}

	query := QueryRequest{
		ProposedUseClass: SpecificUseClass(req["proposedUseClass"].(string)),
		ProposedUseDesc:  req["proposedUseDesc"].(string),
		GFA:              req["GFA"].(float64),
		PostalCode:       req["postalCode"].(string),
		LotNumber:        req["lotNumber"].(string),
		Floor:            int64(req["floor"].(float64)),
		Unit:             int64(req["unit"].(float64)),
	}
	queryResponses, err := c.Accessor.GetSimilarCases(query)
	fmt.Println(queryResponses)
	// json.NewEncoder(w).Encode(&queryResponses)
	response, err := json.Marshal(queryResponses)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}

func (c *KnowledgeGraphComponent) ClearDatabase(w http.ResponseWriter, r *http.Request) {
	err := c.Accessor.ClearDatabase()
	if err != nil {
		log.Fatal(err)
	}
	var retString string = "successfully cleared database"
	json.NewEncoder(w).Encode(&retString)

}
