FROM golang:1.23.4

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY cmd/app ./cmd/app

RUN go build -o main ./cmd/app

EXPOSE 8080
CMD ["./main"]
