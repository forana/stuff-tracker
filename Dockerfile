FROM golang:1.12-alpine AS builder

RUN apk update && apk upgrade && apk add --no-cache bash git openssh

RUN mkdir -p /app
WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 go test ./...
RUN go build .

FROM alpine:latest

RUN mkdir -p /app
WORKDIR /app

COPY --from=0 /app/stuff-tracker /app/stuff-tracker
CMD /app/stuff-tracker
