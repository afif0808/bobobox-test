FROM golang:alpine


WORKDIR /app

COPY go.mod .
COPY go.sum .


COPY . .



