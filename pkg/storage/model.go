package storage

import "time"

// BookDetails defines a Book object
type BookDetails struct {
	ID            string    `bson:"_id,omitempty" json:"_id,omitempty"`
	Name          string    `bson:"bookName,omitempty" json:"bookName,omitempty"`
	Title         string    `bson:"title,omitempty" json:"title,omitempty"`
	Price         string    `bson:"price,omitempty" json:"price,omitempty"`
	PublishedDate string    `bson:"publishedDate,omitempty" json:"publishedDate,omitempty"`
	Created       time.Time `bson:"created" json:"created"`
	Modified      time.Time `bson:"modified" json:"modified"`
	Active        bool      `bson:"active" json:"active"`
}
