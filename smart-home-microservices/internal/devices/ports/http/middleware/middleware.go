package middleware

import (
	"log/slog"
	"net/http"
)

type responseWriter struct {
	w          http.ResponseWriter
	statusCode int
}

func (r *responseWriter) Header() http.Header {
	return r.w.Header()
}

func (r *responseWriter) Write(bytes []byte) (int, error) {
	return r.w.Write(bytes)
}

func (r *responseWriter) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.w.WriteHeader(statusCode)
}

func WithLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("request", "method", r.Method, "path", r.URL.Path)
		writer := &responseWriter{w: w}
		handler.ServeHTTP(writer, r)
		slog.Info("response", "status", writer.statusCode)
	})
}
