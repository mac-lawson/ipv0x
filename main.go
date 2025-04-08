package main

import (
	"flag"
	"fmt"
	"os"

	client "ocnet.maclawson.org/source/browser_client"
	"ocnet.maclawson.org/source/chain"
	"ocnet.maclawson.org/source/server"
	"ocnet.maclawson.org/source/typ"
)

func main() {
	// Define the client and server flags
	clientCmd := flag.NewFlagSet("client", flag.ExitOnError)
	serverCmd := flag.NewFlagSet("server", flag.ExitOnError)

	// Check if a subcommand has been provided
	if len(os.Args) < 2 {
		fmt.Println("ERR: expected 'client' or 'server' subcommands. No transmission.")
		os.Exit(1)
	}

	// Switch between subcommands
	switch os.Args[1] {
	case "client":
		// Parse client command
		clientCmd.Parse(os.Args[2:])
		runClient()
	case "server":
		// Parse server command
		serverCmd.Parse(os.Args[2:])
		runServer()
	default:
		fmt.Println("[IPvOX] HELP DEBUG: No transmit mode given: expected 'client' or 'server' subcommands")
		os.Exit(1)
	}
}

// Client mode logic
func runClient() {
	fmt.Println("Transmitter in client mode...")
	arcStart := &typ.Packet{
		Sender:    "me",
		Receiver:  "demo_user",
		Data:      []byte("HEAD: ARC START, ARC PACKET INIT"),
		Signature: "gml",
	}

	client.V0xClient(arcStart)
}

// Server mode logic
func runServer() {
	fmt.Println("Transmitter in server/command/host mode...")
	bc := chain.NewBlockchain()
	server.StartServer(bc)
}
