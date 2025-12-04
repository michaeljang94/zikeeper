FROM golang:tip-alpine3.22 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o zikeeper

FROM golang:tip-alpine3.22

COPY --from=build /app/zikeeper /app/zikeeper

WORKDIR /app

# EXPOSE 80
# EXPOSE 443
EXPOSE 8080

CMD ["/app/zikeeper"]