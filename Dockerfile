FROM golang:1.20.2

WORKDIR /app

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

COPY . .
RUN go mod tidy

EXPOSE 3000

ENTRYPOINT CompileDaemon --build="go build -o main cmd/main.go" --command="./main"
