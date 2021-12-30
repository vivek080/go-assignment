package storage

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func GetClient() (*Storage, error) {
	URI := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_Name")

	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
	if err != nil {
		return nil, err
	}
	err = client.Connect(context.TODO())
	if err != nil {
		return nil, err
	}
	db := client.Database(dbName)
	s := Storage{Client: client, Database: db}
	return &s, nil
}
