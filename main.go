package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/vivek080/library-app/pkg/book"
	"github.com/vivek080/library-app/pkg/http/rest"
	"github.com/vivek080/library-app/pkg/storage"
)

func main() {

	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("unable to load env file. Err: %s", err)
	}

	client, err := storage.GetClient()
	if err != nil {
		log.Fatalf("Unable to connect to mongo client. Err: %s", err)
		// panic(err)
	}

	s := book.NewService(client)
	router := rest.NewHandler(s)
	auth := rest.NewMiddleware(router)
	log.Println("Library App Server started at port 5000")
	log.Fatal(http.ListenAndServe(":5000", auth))

}
