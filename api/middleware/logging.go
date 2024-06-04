package middleware

import (
	"net/http"
	"time"

	"github.com/charmbracelet/log"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Info(r.Method, r.RequestURI, time.Since(start))
		next.ServeHTTP(w, r)
	})
}