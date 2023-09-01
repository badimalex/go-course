package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"go-course/homework-13/pkg/crawler"
	"go-course/homework-13/pkg/index"
	"go-course/homework-13/pkg/retriever"
)

var api *API

func TestMain(m *testing.M) {
	data := []crawler.Document{
		{
			ID:    1,
			Title: "Go",
			URL:   "https://golang.org/",
		},
		{
			ID:    2,
			Title: "Python",
			URL:   "https://www.python.org/",
		},
	}

	idx := index.New(data)
	r := retriever.New(idx, data)
	api = New(r, data)

	os.Exit(m.Run())
}

func TestSearch(t *testing.T) {
	req, _ := http.NewRequest("GET", "/search/Go", nil)
	rr := httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Search handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var resultDocs []crawler.Document
	_ = json.NewDecoder(rr.Body).Decode(&resultDocs)
	if len(resultDocs) != 1 || resultDocs[0].ID != 1 {
		t.Errorf("Search did not return expected results")
	}
}

func TestCreateDocument(t *testing.T) {
	newDoc := crawler.Document{
		ID:    3,
		Title: "Java",
		URL:   "https://www.java.com/",
	}

	payload, _ := json.Marshal(newDoc)
	req, _ := http.NewRequest("POST", "/documents", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("CreateDocument handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}

func TestGetDocument(t *testing.T) {
	req, _ := http.NewRequest("GET", "/documents/1", nil)
	rr := httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("GetDocument handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var doc crawler.Document
	_ = json.NewDecoder(rr.Body).Decode(&doc)
	if doc.ID != 1 || doc.Title != "Go" {
		t.Errorf("GetDocument did not return expected result")
	}
}

func TestUpdateDocument(t *testing.T) {
	updatedDoc := crawler.Document{
		ID:    1,
		Title: "Golang",
		URL:   "https://www.golang.org/",
	}

	payload, _ := json.Marshal(updatedDoc)
	req, _ := http.NewRequest("PUT", "/documents/1", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("UpdateDocument handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	req, _ = http.NewRequest("GET", "/documents/1", nil)
	rr = httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)

	var doc crawler.Document
	_ = json.NewDecoder(rr.Body).Decode(&doc)
	if doc.Title != "Golang" {
		t.Errorf("UpdateDocument did not correctly update the document")
	}
}

func TestDeleteDocument(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/documents/1", nil)
	rr := httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("DeleteDocument handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	req, _ = http.NewRequest("GET", "/documents/1", nil)
	rr = httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("DeleteDocument did not correctly delete the document")
	}
}
