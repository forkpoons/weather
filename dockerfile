FROM golang:1.17-alpine

WORKDIR /app
RUN apk update && apk add make