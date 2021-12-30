package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/vivek080/go-assignment/pkg/book"
	"github.com/vivek080/go-assignment/pkg/http/rest"
	"github.com/vivek080/go-assignment/pkg/storage"
)

func main() {

	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	client, err := storage.GetClient()
	if err != nil {
		fmt.Println(err)
		// panic(err)
	}

	s := book.NewService(client)
	router := rest.NewHandler(s)
	auth := rest.NewMiddleware(router)
	log.Fatal(http.ListenAndServe(":5000", auth))

}
