package rest

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vivek080/library-app/pkg/book"
)

// NewHandler returns a mux router that staisfies http.Handler interface
func NewHandler(servicer book.Servicer) http.Handler {
	router := mux.NewRouter()

	router.Handle("/book", NewBookHandler(servicer))

	return router
}
