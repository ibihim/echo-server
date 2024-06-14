package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	http.HandleFunc("/", echoHandler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	// Log the incoming request
	log.Printf("Received request: %s %s", r.Method, r.URL)

	// Dump the request
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Printf("Error dumping request: %v\n", err)
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	// Log the dumped request
	log.Printf("Request dump:\n%s", string(requestDump))

	// Set the response content type to plain text
	w.Header().Set("Content-Type", "text/plain")

	// Write the dumped request as the response body
	w.Write(requestDump)
}
