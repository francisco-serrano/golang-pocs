FROM golang:latest

RUN mkdir /files
COPY . /files
WORKDIR /files

RUN go build -o /files/serialization serialization.go

ENTRYPOINT ["/files/serialization"]