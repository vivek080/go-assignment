package rest

import (
	"errors"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vivek080/library-app/pkg/internal/mocks"
	"github.com/vivek080/library-app/pkg/storage"
)

// TestSaveBook_1 test tries to save the book details in DB & expects status code be 201
func TestSaveBook_1(t *testing.T) {
	w := httptest.NewRecorder()
	jsonBody := `{
		"data":{
			"name":"test",
			"title":"test",
			"price":"123",
			"publishedDate":"12/11/2021"
		}
	}`
	r := httptest.NewRequest("POST", "/book", strings.NewReader(jsonBody))
	m := &mocks.MockBookServices{}
	b := storage.BookDetails{ID: "1223", Name: "test", Title: "test", Price: "123", PublishedDate: "12/11/2021", Created: time.Now()}
	m.On("SaveBook", mock.AnythingOfType("*storage.BookDetails")).Return(&b, nil)

	app := BookHandler{m}
	app.createBookDetails(w, r)
	assert.Equal(t, 201, w.Result().StatusCode)
}

// TestSaveBookError_1 test fails to save the book details in DB due to read body error & expects status code be 422
func TestSaveBookError_1(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/book", mocks.ErrReader(0))
	m := &mocks.MockBookServices{}
	app := BookHandler{m}
	app.createBookDetails(w, r)
	assert.Equal(t, 422, w.Result().StatusCode)
	m.AssertNotCalled(t, "SaveBook", mock.AnythingOfType("*storage.BookDetails"))
}

// TestSaveBookError_2 test fails to save the book details in DB due to json body error & expects status code be 400
func TestSaveBookError_2(t *testing.T) {
	w := httptest.NewRecorder()
	jsonBody := `{
		"data":{
			"name":"test",
			"title":"test",
			"price":"123",
			"publishedDate":"12/11/2021"
		
	}`
	r := httptest.NewRequest("POST", "/book", strings.NewReader(jsonBody))
	m := &mocks.MockBookServices{}

	app := BookHandler{m}
	app.createBookDetails(w, r)
	assert.Equal(t, 400, w.Result().StatusCode)
	m.AssertNotCalled(t, "SaveBook", mock.AnythingOfType("*storage.BookDetails"))
}

// TestSaveBookError_3 test fails to save the book details in DB due to SaveBook failure & expects status code be 500
func TestSaveBookError_3(t *testing.T) {
	w := httptest.NewRecorder()
	jsonBody := `{
		"data":{
			"name":"test",
			"title":"test",
			"price":"123",
			"publishedDate":"12/11/2021"
		}
	}`
	r := httptest.NewRequest("POST", "/book", strings.NewReader(jsonBody))
	m := &mocks.MockBookServices{}
	m.On("SaveBook", mock.AnythingOfType("*storage.BookDetails")).Return(nil, errors.New("mock error"))

	app := BookHandler{m}
	app.createBookDetails(w, r)
	assert.Equal(t, 500, w.Result().StatusCode)
}

// TestGetBook_1 test tries to get all the book list from DB & expects status code be 200
func TestGetBook_1(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/book", nil)
	m := &mocks.MockBookServices{}
	b := []storage.BookDetails{{ID: "1223", Name: "test", Title: "test", Price: "123", PublishedDate: "12/11/2021", Created: time.Now()}}
	m.On("GetBookList").Return(b, nil)

	app := BookHandler{m}
	app.getBookDetails(w, r)
	assert.Equal(t, 200, w.Result().StatusCode)
}

// TestGetBookError_1 test fails to get the book list from DB due to GetBookList failure & expects status code be 500
func TestGetBookError_1(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/book", nil)
	m := &mocks.MockBookServices{}
	m.On("GetBookList").Return(nil, errors.New("mock error"))

	app := BookHandler{m}
	app.getBookDetails(w, r)
	assert.Equal(t, 500, w.Result().StatusCode)
}
