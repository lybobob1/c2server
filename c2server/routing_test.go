package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testRouter(t *testing.T) {
	r := newRouter()

	testServer := httptest.NewServer(r)

	resp, err := http.Get(testServer.URL + "/hello")

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status wasn't OK, got %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	respString := string(b)
	expected := "Hello World"

	if respString != expected {
		t.Errorf("Response should be %s, was %s", expected, respString)
	}
}

func testRouterForNonExistentRoute(t *testing.T) {
	r := newRouter()

	testServer := httptest.NewServer(r)

	resp, err := http.Post(testServer.URL+"/hello", "", nil)

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status wasn't 405, got %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	respString := string(b)
	expected := ""

	if respString != expected {
		t.Errorf("Response should be %s, was %s", expected, respString)
	}
}
func TestStaticFileServer(t *testing.T) {
	r := newRouter()
	testServer := httptest.NewServer(r)

	resp, err := http.Get(testServer.URL + "/assets/")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Response was not OK got %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	expectedContentType := "text/html; charset=utf-8"

	if expectedContentType != contentType {
		t.Errorf("Expected %s, got %s", expectedContentType, contentType)
	}
}
