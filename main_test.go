package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatusOK(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	statusHandler()(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
}

func TestStatusBody(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	statusHandler()(rec, req)
	var body statusBody
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatalf("bad json: %v", err)
	}
	if !body.Healthy {
		t.Fatal("expected healthy=true")
	}
	if body.Service != "santosh-pipeline" {
		t.Fatalf("unexpected service name: %s", body.Service)
	}
	if body.Stage != "deploy-demo" {
		t.Fatalf("unexpected stage: %s", body.Stage)
	}
	if body.Revision == "" {
		t.Fatal("expected a revision")
	}
}
