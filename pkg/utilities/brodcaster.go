package utilities

import "fmt"

func Brodcaster() {
	for {
		select {
		case msg := <-messages:
			for _, conn := range clients {
				if msg.address == conn.RemoteAddr().String() {
					continue
				}
				fmt.Fprintln(conn, msg.text) // NOTE ignoring network errors
			}

		case msg := <-leaving:
			for _, conn := range clients {
				fmt.Fprintln(conn, msg.text) // NOTE ignoring network errors
			}
		}
	}
}
