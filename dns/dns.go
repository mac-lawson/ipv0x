package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"
)

// DNSRecord structure
type DNSRecord struct {
	Domain string // Friendly domain name
	PubKey string // Actual public key
}

// Blockchain structure for DNS
type DNSBlock struct {
	Index     int
	Timestamp string
	Record    DNSRecord
	PrevHash  string
	Hash      string
}

type DNSBlockchain struct {
	blocks []DNSBlock
}

// Calculate the hash for a DNS block
func calculateDNSHash(block DNSBlock) string {
	record := fmt.Sprintf("%d%s%s%s", block.Index, block.Timestamp, block.Record, block.PrevHash)
	h := sha256.New()
	h.Write([]byte(record))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Create a new DNS block
func createDNSBlock(prevBlock DNSBlock, record DNSRecord) DNSBlock {
	block := DNSBlock{
		Index:     prevBlock.Index + 1,
		Timestamp: time.Now().String(),
		Record:    record,
		PrevHash:  prevBlock.Hash,
	}
	block.Hash = calculateDNSHash(block)
	return block
}

// Add a block to the DNS blockchain
func (bc *DNSBlockchain) addBlock(record DNSRecord) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := createDNSBlock(prevBlock, record)
	bc.blocks = append(bc.blocks, newBlock)
}

// Create the genesis block for DNS
func createDNSGenesisBlock() DNSBlock {
	return DNSBlock{
		Index:     0,
		Timestamp: time.Now().String(),
		Record:    DNSRecord{},
		PrevHash:  "",
		Hash:      calculateDNSHash(DNSBlock{}),
	}
}

// Initialize a new DNS blockchain
func NewDNSBlockchain() *DNSBlockchain {
	return &DNSBlockchain{[]DNSBlock{createDNSGenesisBlock()}}
}

// Handle incoming DNS connections
func handleDNSConnection(conn net.Conn, bc *DNSBlockchain) {
	defer conn.Close()
	var record DNSRecord
	decoder := json.NewDecoder(conn)
	if err := decoder.Decode(&record); err != nil {
		log.Println("Failed to decode DNS record:", err)
		return
	}
	// Add the DNS record to the blockchain
	bc.addBlock(record)
	fmt.Println("DNS record added to blockchain:", record)
}

// Start the DNS server
func startDNSServer(bc *DNSBlockchain) {
	ln, err := net.Listen("tcp", ":9099")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleDNSConnection(conn, bc)
	}
}

func Start() {
	// Initialize DNS blockchain
	bc := NewDNSBlockchain()
	// Start the DNS server
	startDNSServer(bc)
}
