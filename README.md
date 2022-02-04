# Library App

A RESTful API Microservice for Managing Book Details


## Local Deployment

before runing this project please create and `.env` file and copy the details from `.sample.env` file

To deploy this project run

```bash
  # to build the binary for the project
  go build
  # then run the binary start the service
  ./library-app 
```
you will see `Book Server started at port 80` in the terminal

## Deployment using K8s
To deploy this project using K8s yaml file 

```bash
  # apply this command to create all the deployment present in the directory
  kubectl apply -f k8s/
```
after successful creation of services you can access the endpoints using
```
localhost:30081
```


## API Documentation
Please visit the API Documentation for the API Endpoints and the body data

[Documentation](https://github.com/vivek080/library-app/blob/main/arch/openapi.json)


## Unit testing

To run the unit testing:

```
go test -coverprofile cover.out 
```

To view coverage report in browser, run:

```
go tool cover -html=cover.out