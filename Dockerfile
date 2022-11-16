FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o app ./cmd/tcp/main.go

EXPOSE 8080

CMD ["./app"]

ENTRYPOINT ["./app", "--port", "8080"]
