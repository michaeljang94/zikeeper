FROM golang:tip-alpine3.22 AS base

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o zikeeper

EXPOSE 8080

CMD ["/app/zikeeper"]