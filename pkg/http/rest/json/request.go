package json

import "github.com/vivek080/go-assignment/pkg/storage"

// BookRequestData defines a request book struct data
type BookRequestData struct {
	Data BookRequestDetails `json:"data"`
}

// BookRequestDetails defines a request Book struct
type BookRequestDetails struct {
	Name          string `json:"name,omitempty"`
	Title         string `json:"title,omitempty"`
	Price         string `json:"price,omitempty"`
	PublishedDate string `json:"publishedDate,omitempty"`
}

// MapfromRequest maps the request object to the storage object
func (bookRequest *BookRequestData) MapfromRequest() *storage.BookDetails {
	var book storage.BookDetails
	book.Name = bookRequest.Data.Name
	book.Title = bookRequest.Data.Title
	book.Price = bookRequest.Data.Price
	book.PublishedDate = bookRequest.Data.PublishedDate
	return &book
}
