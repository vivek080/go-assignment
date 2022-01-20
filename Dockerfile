# Stage 1
FROM golang:1.17 AS buildenv
ENV GOOS=linux GOARCH=amd64 CGO_ENABLED=0
ADD . /go/src/app/
WORKDIR /go/src/app/
RUN go build -o library-app

# Stage 2
FROM alpine:latest
RUN apk update
COPY --from=buildenv /go/src/app/library-app /usr/local/bin/library-app
CMD ["library-app"]