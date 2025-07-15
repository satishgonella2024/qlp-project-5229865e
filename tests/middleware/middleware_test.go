package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestLogRequestMiddleware(t *testing.T) {
    // Create a test request
    req, err := http.NewRequest("GET", "/health", nil)
    if err != nil {
        t.Fatal(err)
    }

    // Create a test response recorder
    rr := httptest.NewRecorder()

    // Create a new server mux and wrap it with the logging middleware
    mux := http.NewServeMux()
    loggedRouter := logRequestMiddleware(mux)

    // Add routes
    mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
    })

    // Serve the request
    loggedRouter.ServeHTTP(rr, req)

    // Check the response status code
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check if the request was logged
    // (Add assertions to check the log output)
}