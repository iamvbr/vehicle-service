package middlewares

import (
	"log/slog"
	"net/http"
	"os"
	"time"
)

type StatusResponseWriter struct {
	http.ResponseWriter
	status int
}

func (w *StatusResponseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}
func Logger(inner http.Handler) http.Handler {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		srw := &StatusResponseWriter{ResponseWriter: w}

		defer func(res *StatusResponseWriter, req *http.Request) {
			logger.InfoContext(req.Context(), "REQ",
				"ResponseTime", time.Since(start).Nanoseconds()/1000,
				"Method", req.Method,
				"UserAgent", req.UserAgent(),
				"URI", req.RequestURI,
				"Code", srw.status,
			)
		}(srw, r)

		inner.ServeHTTP(srw, r)

	})
}
