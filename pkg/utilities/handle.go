package utilities

import (
	"bufio"
	"net"
)

var clients = make(map[string]net.Conn)
var leaving = make(chan message)
var messages = make(chan message)

type message struct {
	text    string
	address string
}

func Handle(conn net.Conn) {
	username, _ := GenerateUsername()
	clients[username] = conn
	input := bufio.NewScanner(conn)

	messages <- newMessage(username, " joined.")

	for input.Scan() {
		messages <- newMessage(username, ": "+input.Text())
	}

	delete(clients, conn.RemoteAddr().String())

	leaving <- newMessage(username, " has left.")
	conn.Close()
}

func newMessage(username, msg string) message {
	return message{
		text:    username + msg,
		address: username,
	}
}
