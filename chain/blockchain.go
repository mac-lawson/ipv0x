package chain

import (
	"crypto/sha256"
	"time"

	"ocnet.maclawson.org/source/typ"
)

/*
1. Initialize a new blockchain:

	package main

	import (
		"fmt"
		"ocnet.maclawson.org/source/typ"
		"chain"
	)

	func main() {
		bc := chain.NewBlockchain()
		fmt.Println(bc)
	}

2. Add a new block to the blockchain:

	package main

	import (
		"fmt"
		"ocnet.maclawson.org/source/typ"
		"chain"
	)

	func main() {
		bc := chain.NewBlockchain()
		packet := typ.Packet{Data: "Sample Data"}
		bc.AddBlock(packet)
		fmt.Println(bc)
	}
*/
type Block struct {
	Index     int
	Timestamp string
	Packet    typ.Packet
	PrevHash  string
	Hash      string
}

type Blockchain struct {
	blocks []Block
}

func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.Packet.Data) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	return string(h.Sum(nil))
}

func CreateBlock(prevBlock Block, packet typ.Packet) Block {
	block := Block{
		Index:     prevBlock.Index + 1,
		Timestamp: time.Now().String(),
		Packet:    packet,
		PrevHash:  prevBlock.Hash,
	}
	block.Hash = calculateHash(block)
	return block
}

func (bc *Blockchain) AddBlock(packet typ.Packet) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := CreateBlock(prevBlock, packet)
	bc.blocks = append(bc.blocks, newBlock)
}

func createGenesisBlock() Block {
	return Block{
		Index:     0,
		Timestamp: time.Now().String(),
		Packet:    typ.Packet{},
		PrevHash:  "",
		Hash:      calculateHash(Block{}),
	}
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]Block{createGenesisBlock()}}
}
