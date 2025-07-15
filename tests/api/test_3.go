{
    "code": "package main

import (
    \"encoding/json\"
    \"log\"
    \"net/http\"
    \"os\"
)

type healthResponse struct {
    Status string `json:\"status\"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set(\"Content-Type\", \"application/json\")
    resp := healthResponse{Status: \"ok\"}
    json.NewEncoder(w).Encode(resp)
}

func logMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf(\"%s %s\", r.Method, r.RequestURI)
        next.ServeHTTP(w, r)
    })
}

func main() {
    http.Handle(\"/health\", logMiddleware(http.HandlerFunc(healthHandler)))
    log.Println(\"Server started on :8080\")
    log.Fatal(http.ListenAndServe(\":8080\", nil))
}",

    "tests": "package main

import (
    \"encoding/json\"
    \"net/http\"
    \"net/http/httptest\"
    \"testing\"
)

func TestHealthHandler(t *testing.T) {
    req, err := http.NewRequest(\"GET\", \"/health\", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(healthHandler)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf(\"handler returned wrong status code: got %v want %v\", status, http.StatusOK)
    }

    expected := `{\"status\":\"ok\"}`
    if rr.Body.String() != expected {
        t.Errorf(\"handler returned unexpected body: got %v want %v\", rr.Body.String(), expected)
    }
}

func TestLogMiddleware(t *testing.T) {
    req, err := http.NewRequest(\"GET\", \"/health\", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := logMiddleware(http.HandlerFunc(healthHandler))
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf(\"handler returned wrong status code: got %v want %v\", status, http.StatusOK)
    }

    expected := `{\"status\":\"ok\"}`
    var resp healthResponse
    err = json.Unmarshal(rr.Body.Bytes(), &resp)
    if err != nil {
        t.Errorf(\"failed to unmarshal response: %v\", err)
    }

    if resp.Status != \"ok\" {
        t.Errorf(\"unexpected response status: got %v want %v\", resp.Status, \"ok\")
    }
}",

    "documentation": "This Go code sets up a simple web server with a health check endpoint (/health) and a middleware for logging incoming requests. The healthHandler function returns a JSON response with a status of 'ok'. The logMiddleware function logs the HTTP method and request URI for each incoming request. The main function sets up the server to handle the /health endpoint with the logMiddleware and healthHandler.",

    "dependencies": ["encoding/json", "log", "net/http", "os"]
}