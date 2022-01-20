package book

import (
	"github.com/vivek080/library-app/pkg/storage"
)

// Repository defines funcs that a storage implementation should implement
type Repository interface {
	SaveBook(book *storage.BookDetails) (*storage.BookDetails, error)
	GetBookList() ([]storage.BookDetails, error)
}

// Servicer defines funcs that a service implementation should implement
type Servicer interface {
	SaveBook(book *storage.BookDetails) (*storage.BookDetails, error)
	GetBookList() ([]storage.BookDetails, error)
}
