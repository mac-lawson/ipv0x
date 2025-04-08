package server

import (
	"encoding/json"
	"fmt"
	"net"

	"ocnet.maclawson.org/source/chain"
	"ocnet.maclawson.org/source/typ"
)

// Handle incoming connections
func handleConnection(conn net.Conn, bc *chain.Blockchain) {
	defer conn.Close()

	// Read the message
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	var packet typ.Packet
	err1 := json.Unmarshal(buffer[:n], &packet)
	if err1 != nil {
		fmt.Println("Error reading:", err)
		return
	}
	fmt.Println("Received transmission:", string(buffer[:n]))
	fmt.Println("Decoded packet values:", packet)
}
