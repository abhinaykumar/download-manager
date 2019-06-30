package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"./downloader"
)

func TestDownloadSuccess(t *testing.T) {
	fmt.Printf("====== Testing Download with valid parameters should [PASS] =====")
	var jsonParams = []byte(`{"url": "https://file-examples.com/wp-content/uploads/2017/11/file_example_MP3_700KB.mp3","threads": 4}`)
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
	expected := `File is downloaded from url: https://file-examples.com/wp-content/uploads/2017/11/file_example_MP3_700KB.mp3`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestDownloadFail(t *testing.T) {
	fmt.Printf("==== Testing Download with invalid parameters should [FAIL] =====")
	var jsonParams = []byte(`{"url": "https://github.com/abhinaykumar/fake.html","threads": null}`)
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(jsonParams))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(downloader.Download)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}
