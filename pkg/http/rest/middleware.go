package rest

import (
	"bytes"
	"encoding/base64"
	"net/http"
	"os"
	"strings"
)

// Middleware implements authentication for certain endpoints
type Middleware struct {
	next http.Handler
}

// NewMiddleware is a func for creating Middleware
func NewMiddleware(next http.Handler) *Middleware {
	m := Middleware{next: next}
	return &m
}

// ServeHTTP implements authentication and if successfully authenticated, will chain the next http.Handler
func (m *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	const basicAuthPrefix = "Basic "
	user := []byte(os.Getenv("BASIC_AUTH_USERNAME"))
	password := []byte(os.Getenv("BASIC_AUTH_PASSWORD"))

	auth := r.Header.Get("Authorization")
	if strings.HasPrefix(auth, basicAuthPrefix) {
		payload, err := base64.StdEncoding.DecodeString(auth[len(basicAuthPrefix):])
		if err != nil {
			w.WriteHeader(401)
			return
		}
		pair := bytes.SplitN(payload, []byte(":"), 2)
		if len(pair) == 2 && bytes.Equal(pair[0], user) && bytes.Equal(pair[1], password) {
			m.next.ServeHTTP(w, r)
			return
		}
	}
	http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
}
