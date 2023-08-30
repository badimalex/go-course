// webapp/webapp_test.go

package webapp

import (
	"encoding/json"
	"go-course/homework-12/pkg/crawler"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestIndexHandler(t *testing.T) {
	idx := map[string][]int{"test": {1, 2}}

	docs := []crawler.Document{}

	app := New(idx, docs)
	router := mux.NewRouter()
	router.HandleFunc("/index", app.indexHandler)

	req, _ := http.NewRequest("GET", "/index", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	var responseIdx map[string][]int
	json.Unmarshal(rr.Body.Bytes(), &responseIdx)

	if len(responseIdx) != 1 {
		t.Errorf("Expected 1 items, got %d", len(responseIdx))
	}
}

func TestDocsHandler(t *testing.T) {
	idx := map[string][]int{"test": {1, 2}}
	docs := []crawler.Document{
		{ID: 1, URL: "http://example.com", Title: "Example", Body: "Body"},
		{ID: 2, URL: "http://test.com", Title: "Test", Body: "Test Body"},
	}

	app := New(idx, docs)
	router := mux.NewRouter()
	router.HandleFunc("/docs", app.docsHandler)

	req, _ := http.NewRequest("GET", "/docs", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	var responseDocs []crawler.Document
	json.Unmarshal(rr.Body.Bytes(), &responseDocs)

	if len(responseDocs) != 2 {
		t.Errorf("Expected 2 items, got %d", len(responseDocs))
	}
}
