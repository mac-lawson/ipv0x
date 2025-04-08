package client

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"ocnet.maclawson.org/source/typ"
)

func V0xClient(transport_packet *typ.Packet) {
	// Connect to server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()
	encodedData, err := json.Marshal(transport_packet)
	if err != nil {
		fmt.Println("Error encoding:", err)
	}
	// Send a message
	message := encodedData
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}
	fmt.Println("Message sent:", string(message))

	// Receive a response
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	fmt.Println("Received response:", string(buffer[:n]))
}
