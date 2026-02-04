package main

import (
	"fmt"
	"log"
	"net"
	"os"

	NetCat "NetCat/dependencies"
)

func main() {
	port := "8989"
	NetCat.Historyclean()
	if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
	if len(os.Args) == 2 {
		port = os.Args[1]
	}
	tempport, ok := NetCat.PersoAtoi(port)
	if !ok || tempport < 1024 || tempport > 49151 {
		fmt.Println("Invalid port number. Please provide a port between 1024 and 49151.")
		return
	}
	lisn, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Printf("error for listen: %v", err)
		return
	}
	defer lisn.Close()
	log.Printf("listn in port : %s", port)

	for {
		conn, err := lisn.Accept()
		if err != nil {
			log.Printf("err for accept thisreader: %v", err)
			continue
		}
		go NetCat.HandleConnection(conn)
	}
}
