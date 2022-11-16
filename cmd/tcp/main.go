package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BradPreston/go-chat/pkg/server"
	"github.com/BradPreston/go-chat/pkg/utilities"
)

func main() {
	listener, err := server.Listen("0.0.0.0:8080")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println("*** Chat App Is Running! ***")
	go utilities.Brodcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go utilities.Handle(conn)
	}
}
