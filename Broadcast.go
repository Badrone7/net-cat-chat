package main

import (
	"log"
	"net"
)

func Broadcast(sender net.Conn, message string) {
	mutex.Lock()
	defer mutex.Unlock()
	for conn := range clients {
		if conn != sender {
			_, err := conn.Write([]byte(message))
			if err != nil {
				log.Printf("Error broadcasting to %s: %v\n", clients[conn], err)
			}
		}
	}
}
