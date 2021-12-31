package json

import (
	"time"

	"github.com/vivek080/go-assignment/pkg/storage"
)

type ErrorData struct {
	Data Error `json:"data"`
}

type Error struct {
	Status      int    `json:"status"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

// BookResponseData defines a book reponse data
type BookResponseData struct {
	Data BookResponseDetails `json:"data"`
}

// BookResponseDetails defines a book struct
type BookResponseDetails struct {
	Name          string    `json:"name,omitempty"`
	Title         string    `json:"title,omitempty"`
	Price         string    `json:"price,omitempty"`
	PublishedDate string    `json:"publishedDate,omitempty"`
	Created       time.Time `json:"created"`
	Modified      time.Time `json:"modified"`
}

// BookListResponseData defines a book list struct
type BookListResponseData struct {
	Data []BookResponseDetails `json:"data"`
}

// MapToResponse maps the storage object to the reponse object
func (response *BookResponseData) MapToResponse(book *storage.BookDetails) {
	response.Data.Name = book.Name
	response.Data.Title = book.Title
	response.Data.Price = book.Price
	response.Data.PublishedDate = book.PublishedDate
	response.Data.Created = book.Created
	response.Data.Modified = book.Modified
}

// MapToResponse maps the storage list object to the reponse list object
func (responseList *BookListResponseData) MapToResponse(bookList []storage.BookDetails) {
	if len(bookList) > 0 {
		for _, book := range bookList {
			var response BookResponseDetails
			response.Name = book.Name
			response.Title = book.Title
			response.Price = book.Price
			response.PublishedDate = book.PublishedDate
			response.Created = book.Created
			response.Modified = book.Modified
			responseList.Data = append(responseList.Data, response)
		}
	} else {
		responseList.Data = make([]BookResponseDetails, 0)
	}
}
