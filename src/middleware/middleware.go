package main

import (
    "log"
    "net/http"
    "time"
)

// Define a middleware function to log requests
func logRequestMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        lrw := NewLoggingResponseWriter(w)

        next.ServeHTTP(lrw, r)

        log.Printf(
            "[%s] %s %s %d %d",
            r.Method,
            r.RequestURI,
            r.RemoteAddr,
            lrw.statusCode,
            time.Since(start).Milliseconds(),
        )
    })
}

// Define a custom response writer to capture response status code
type loggingResponseWriter struct {
    http.ResponseWriter
    statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
    return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
    lrw.statusCode = code
    lrw.ResponseWriter.WriteHeader(code)
}