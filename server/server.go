package server

import (
	"log"
	"net"

	"ocnet.maclawson.org/source/chain"
)

// Start the server
func StartServer(bc *chain.Blockchain) {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn, bc)
	}
}
