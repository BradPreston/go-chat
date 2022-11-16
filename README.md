# Golang Chat App
This is a very simple TCP chat app that runs in the terminal.

## How to run
```
go run ./cmd/netcat/netcat.go
```

## How the app works
This is chat app that runs on a TCP server. When you connect via netcat.go, an API call is made to [random-data-api](https://random-data-api.com/api/v2/users) (from inside handle.go) to give you a randomly generated email. The app concurrently handles client connections to the server and notifies the other users of a client connecting. When a client disconnects, the other clients are also notified of their leaving. 