package rest

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/vivek080/go-assignment/pkg/book"
	httpJson "github.com/vivek080/go-assignment/pkg/http/rest/json"
)

// BookHandler handles the book endpoints
type BookHandler struct {
	Servicer book.Servicer
}

// NewBookHandler returns a *BookHandler that satisfies http.Handler interface
func NewBookHandler(servicer book.Servicer) http.Handler {
	bh := BookHandler{Servicer: servicer}
	return &bh
}

// ServeHTTP handles the adding/updating the information
func (bh *BookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		bh.createBookDetails(w, r)
	case "GET":
		bh.getBookDetails(w, r)
	}
}

// createBookDetails creates a Book detail in DB.
func (bh *BookHandler) createBookDetails(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var request httpJson.BookRequestData
	var response httpJson.BookResponseData

	w.Header().Set("Content-Type", "application/json")

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err.Error())
		httpJson.SendHTTPErrorResponse(w, http.StatusUnprocessableEntity, "Unable to read the request", err.Error())
		return
	}

	err = json.Unmarshal(reqBody, &request)
	if err != nil {
		log.Print(err.Error())
		httpJson.SendHTTPErrorResponse(w, http.StatusBadRequest, "Unable to unmarshal the request", err.Error())
		return
	}

	sBook := request.MapfromRequest()

	bookData, err := bh.Servicer.SaveBook(sBook)
	if err != nil {
		log.Print("Unable to save Book details", "err", err)
		httpJson.SendHTTPErrorResponse(w, http.StatusInternalServerError, "Unable to save Book details", err.Error())
		return
	}

	response.MapToResponse(bookData)

	httpJson.SendHTTPSuccessResponse(w, http.StatusCreated, response)

}

// getBookDetails gets book list
func (bh *BookHandler) getBookDetails(w http.ResponseWriter, r *http.Request) {
	var response httpJson.BookListResponseData

	w.Header().Set("Content-Type", "application/json")

	bookList, err := bh.Servicer.GetBookList()
	if err != nil {
		log.Print("Unable to get Book list", "err", err)
		httpJson.SendHTTPErrorResponse(w, http.StatusInternalServerError, "Unable to get Book list", err.Error())
		return
	}

	response.MapToResponse(bookList)

	httpJson.SendHTTPSuccessResponse(w, http.StatusOK, response)
}
