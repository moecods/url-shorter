package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUrlHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/test2", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UrlHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusFound)
	}

	expectedLocation := "https://www.bing.com"
	if rr.Header().Get("Location") != expectedLocation {
		t.Errorf("handler returned wrong location header: got %v want %v", rr.Header().Get("Location"), expectedLocation)
	}
}
