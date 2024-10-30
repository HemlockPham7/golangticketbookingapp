FROM golang:latest

WORKDIR /src/app

RUN go install github.com/air-verse/air@latest

COPY . .

RUN go mod tidy