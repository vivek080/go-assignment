package main

import (
	"log"
	"net/http"

	"github.com/vivek080/library-app/pkg/book"
	"github.com/vivek080/library-app/pkg/http/rest"
	"github.com/vivek080/library-app/pkg/storage"
)

func main() {

	client, err := storage.GetClient()
	if err != nil {
		log.Fatalf("Unable to connect to mongo client. Err: %s", err)
		// panic(err)
	}

	s := book.NewService(client)
	router := rest.NewHandler(s)
	auth := rest.NewMiddleware(router)
	log.Println("Library App Server started at port 80")
	log.Fatal(http.ListenAndServe(":80", auth))

}
