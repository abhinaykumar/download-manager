package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"./downloader"
)

func TestDownload(t *testing.T) {

	var jsonParams = []byte(`{"url": "http://localhost:8080/data.txt","threads": 4}`)
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(jsonParams))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(downloader.Download)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `File is downloaded from url: http://localhost:8080/data.txt`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
