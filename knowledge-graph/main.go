package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	Serving_URL = "127.0.0.1:8080"
)

func CreateHandlers(router *mux.Router) error {
	c, err := NewKnowledgeGraphComponent()
	if err != nil {
		return err
	}
	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		response := "status ok"
		log.Printf("[INFO]: %s", response)
		w.Write([]byte(response))
	}).Methods("GET")
	router.HandleFunc("/getCases", c.GetCases)

	return nil
}

func main() {
	router := mux.NewRouter()
	err := CreateHandlers(router)
	if err != nil {
		log.Println("[ERROR]")
	}
	http.ListenAndServe(Serving_URL, router)
}
