FROM golang:1.20.2

WORKDIR /app

COPY . .
RUN go mod tidy

EXPOSE 3000

ENTRYPOINT go run cmd/server.go
