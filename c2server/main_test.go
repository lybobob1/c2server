package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(handler)
	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: expected %v got %v", status, http.StatusOK)
	}

	expected := "Hello World"
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("hanlder returned unexpected body: %v", actual)
	}
}
