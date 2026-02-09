package NetCat

import (
	"log"
	"net"
)

/* LimitsChecker checks if the number of connected clients has reached the limit (10).
   If the limit is reached, it sends a message to the client and closes the connection.
   It returns true if the connection is accepted, and false if it is rejected.
*/

func LimitsChecker(conn net.Conn) bool {
	mutex.Lock()
	if len(clients) >= 10 {
		log.Println("connection rejected: chat room is full")
		conn.Write([]byte("chat is full. please try again later.\n"))
		conn.Close()
		mutex.Unlock()
		return false
	}
	mutex.Unlock()
	return true
}

// NameChecker checks if the provided name is unique among the connected clients.
func NameChecker(name string) bool {
	mutex.Lock()
	defer mutex.Unlock()
	for _, existingName := range clients {
		if existingName == name {
			return false
		}
	}
	return true
}

// UserDelete removes a client from the clients map when they disconnect.
func UserDelete(clients map[net.Conn]string, conn net.Conn) {
	mutex.Lock()
	defer mutex.Unlock()
	delete(clients, conn)
}
