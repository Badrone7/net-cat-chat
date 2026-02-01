package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	clients = make(map[net.Conn]string)
	mutex   = &sync.Mutex{}
)

var historyMutex = &sync.Mutex{}

func AppendToHistory(message string) {
	historyMutex.Lock()
	defer historyMutex.Unlock()

	file, err := os.OpenFile(
		"history.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0o644,
	)
	if err != nil {
		log.Println("error opening history file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(message)
	if err != nil {
		log.Println("error writing to history file:", err)
	}
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

func HandleConnection(conn net.Conn) {
	// Sdd l'connection dyal had l'client mli nsaliw mno
	defer conn.Close()
	history, err := os.ReadFile("history.txt")
	welcomeLogo, err := os.ReadFile("logo.txt")
	welcomeLogo = append(welcomeLogo, []byte("\n----- Chat History -----\n")...)
	welcomeLogo = append(welcomeLogo, history...)
	if err == nil && len(history) > 0 {
		conn.Write(welcomeLogo)
	} else {
		conn.Write(welcomeLogo)
	}
	if err != nil {
		log.Println("error reading welcome file:", err)
		return
	}
	// wahd reader smaile kaireadi smiya dyal l'client
	reader := bufio.NewReader(conn)
	var name string
	conn.Write([]byte("\n[ENTER YOUR NAME]:"))
	for {
		// read until find \n
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Println("error connection Interruption", err)
			return
		}
		// Nqy smiya mn spaces w l'retour à la ligne
		name = strings.TrimSpace(input)
		if name != "" && NameChecker(name) {
			// name != "" left bocle
			break
		}

		conn.Write([]byte("[ENTER YOUR NAME]:"))
	}
	mutex.Lock()
	clients[conn] = name
	mutex.Unlock()
	log.Printf("%s has joined our chat... \n", name)
	joinMsg := fmt.Sprintf("\n%s has joined our chat...\n", name)

	Broadcast(conn, joinMsg)
	// ----------------------------------
	for {
		tm := time.Now().Format("2006-01-02 15:04:05")
		msg := fmt.Sprintf("\n[%s] [%s] :", tm, name)
		_, err := conn.Write([]byte(msg))
		if err != nil {
			log.Println("error connection Interruption with client", err)
			break
		}
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println(name, "left the channel. type of error : ", err)
			Broadcast(conn, fmt.Sprintf("\n%s has left the chat.\n", name))
			UserDelete(clients, conn)
			break
		}
		if message == "\n" || message == "\r\n" {
			continue
		}
		fnlmsg := fmt.Sprintf("%s %s", msg, message)
		AppendToHistory(fnlmsg)
		Broadcast(conn, fnlmsg)
	}
}

func UserDelete(clients map[net.Conn]string, conn net.Conn) {
	mutex.Lock()
	defer mutex.Unlock()
	delete(clients, conn)
}
