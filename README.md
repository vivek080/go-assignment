# Go-Assignment

This Project is used to create and fetch Book Details in DB


## Local Deployment

before runing this project please create and `.env` file and copy the details from `.sample.env` file

To deploy this project run

```bash
  go run main.go
```
you will see `Book Server started at port 5000` in the terminal

you can use the `/book` endpoint to create a book in the DB using POST method

you can use the `/book` endpoint to get list of all the books from the DB using GET method
## Unit testing

To run the unit testing:

```
go test -coverprofile cover.out 
```

To view coverage report in browser, run:

```
go tool cover -html=cover.out