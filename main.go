package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	port := "8989"
	if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
	if len(os.Args) == 2 {
		port = os.Args[1]
	}
	lisn, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("error for listen: %v", err)
	}
	defer lisn.Close()
	log.Printf("listn in port : %s", port)

	for {
		conn, err := lisn.Accept()
		if err != nil {
			log.Printf("err for accept thisreader: %v", err)
			continue
		}
		go HandleConnection(conn)
	}
}
