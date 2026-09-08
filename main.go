package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type statusBody struct {
	Service  string `json:"service"`
	Stage    string `json:"stage"`
	Healthy  bool   `json:"healthy"`
	Revision string `json:"revision"`
	Version  string `json:"version"`
}

func statusHandler() http.HandlerFunc {
	revision := os.Getenv("GIT_REVISION")
	if revision == "" {
		revision = "dev"
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(statusBody{
			Service:  "santosh-pipeline",
			Stage:    "deploy-demo",
			Healthy:  true,
			Revision: revision,
			Version:  "1.0.0",
		})
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	srv := &http.Server{Addr: ":" + port, Handler: statusHandler()}
	log.Printf("santosh-pipeline demo listening on :%s", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
