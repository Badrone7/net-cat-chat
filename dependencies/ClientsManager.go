package NetCat

import (
	"log"
	"net"
)

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

func UserDelete(clients map[net.Conn]string, conn net.Conn) {
	mutex.Lock()
	defer mutex.Unlock()
	delete(clients, conn)
}
