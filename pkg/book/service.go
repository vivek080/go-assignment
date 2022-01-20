package book

import (
	"github.com/vivek080/library-app/pkg/storage"
)

// Service implements all functions of Services interface
type Service struct {
	Repo Repository
}

// NewService is a function for creating a new service.
func NewService(repo Repository) *Service {
	return &Service{repo}
}

// SaveBook saves the book details in DB
func (s *Service) SaveBook(book *storage.BookDetails) (*storage.BookDetails, error) {
	bData, err := s.Repo.SaveBook(book)
	return bData, err
}

// GetBookList gets the list of book details from DB
func (s *Service) GetBookList() ([]storage.BookDetails, error) {
	bData, err := s.Repo.GetBookList()
	return bData, err
}
