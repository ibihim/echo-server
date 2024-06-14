package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEchoHandler(t *testing.T) {
	// Create a new HTTP request
	requestBody := "Hello, Echo Server!"
	req, err := http.NewRequest("POST", "/", strings.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "text/plain")

	// Create a new HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the echo handler
	handler := http.HandlerFunc(echoHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Get the response body
	responseBody, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Convert the response body to a string
	responseString := string(responseBody)

	// Check if the response body contains the request body
	if !strings.Contains(responseString, requestBody) {
		t.Errorf("response body does not contain the request body: got %v, want %v", responseString, requestBody)
	}

	// Check if the response body contains the request method
	if !strings.Contains(responseString, "POST") {
		t.Errorf("response body does not contain the request method: got %v, want %v", responseString, "POST")
	}

	// Check if the response body contains the request URL
	if !strings.Contains(responseString, "/") {
		t.Errorf("response body does not contain the request URL: got %v, want %v", responseString, "/")
	}

	// Check if the response body contains the request header
	if !strings.Contains(responseString, "Content-Type: text/plain") {
		t.Errorf("response body does not contain the request header: got %v, want %v", responseString, "Content-Type: text/plain")
	}
}
