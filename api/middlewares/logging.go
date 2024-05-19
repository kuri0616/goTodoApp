package middlewares

import (
	"log"
	"net/http"
)

type resLoggingWriter struct {
	http.ResponseWriter
	code int
}

func newResLoggingWriter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{w, http.StatusOK}
}

func (w *resLoggingWriter) WriteHeader(code int) {
	w.code = code
	w.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceID := newTraceID()
		log.Printf("traceID: %d, method: %s, path: %s", traceID, r.Method, r.URL.Path)
		ctx := SetTraceID(r.Context(), traceID)
		r = r.WithContext(ctx)
		resWriter := newResLoggingWriter(w)
		next.ServeHTTP(resWriter, r)
		log.Printf("traceID: %d, status: %d", traceID, resWriter.code)
	})
}
