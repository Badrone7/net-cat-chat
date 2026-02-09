package NetCat

import (
	"log"
	"net"
)

// Broadcast sends a message to all connected clients except the sender.
func Broadcast(sender net.Conn, message string) {
	mutex.Lock()
	defer mutex.Unlock()
	for conn, name := range clients {
		if conn != sender {
			_, err := conn.Write([]byte(message))
			if err != nil {
				log.Printf("\nError broadcasting to %s: %v\n", clients[conn], err)
			}
			sendPrompt(conn, name)
		}
	}
}
