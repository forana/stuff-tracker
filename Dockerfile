FROM golang:1.12-alpine AS builder

RUN apk update && apk upgrade && apk add --no-cache bash git openssh build-base

RUN mkdir -p /app
WORKDIR /app

COPY . .
RUN go test ./...
RUN go build .

FROM alpine:latest

RUN mkdir -p /app
WORKDIR /app

COPY --from=0 /app/stuff-tracker /app/stuff-tracker
COPY --from=0 /app/static /app/static
COPY --from=0 /app/migrations /app/migrations

CMD /app/stuff-tracker
