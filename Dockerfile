FROM golang:1.23-alpine AS builder

WORKDIR /url-shortener

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main ./cmd/link/main.go

EXPOSE 8080

CMD ["./main"]