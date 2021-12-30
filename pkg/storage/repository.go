package storage

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	bookCollection = "books"
)

// SaveBook saves the book details into DB
func (s *Storage) SaveBook(book *BookDetails) (*BookDetails, error) {

	book.Created = time.Now()
	book.Active = true

	collection := s.Database.Collection(bookCollection)
	r, err := collection.InsertOne(context.TODO(), book)
	if err != nil {
		return nil, err
	}

	book.ID = r.InsertedID.(primitive.ObjectID).Hex()

	return book, nil
}

// GetBookList fetches list of all the book details from DB
func (s *Storage) GetBookList() ([]BookDetails, error) {
	var bookList []BookDetails

	filter := bson.M{"active": true}
	collection := s.Database.Collection(bookCollection)
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	err = cur.All(context.TODO(), &bookList)
	if err != nil {
		return nil, err
	}

	return bookList, nil
}
