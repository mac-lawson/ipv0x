package main

import (
	"encoding/json"
	"fmt"

	"ocnet.maclawson.org/source/cryptomath"
	"ocnet.maclawson.org/source/encrypt"
	"ocnet.maclawson.org/source/typ"
)

func _main() {
	// Initialize blockchain and add genesis block.
	// bc := chain.NewBlockchain()
	privkey1, pubkey1, err1 := cryptomath.GenerateKeys(2048)
	privkey2, pubkey2, err2 := cryptomath.GenerateKeys(2048)

	if err1 != nil || err2 != nil {
		fmt.Println("error with generating keychain", pubkey1)
	}

	arcStart := &typ.Packet{
		Sender:    "me",
		Receiver:  "demo_user",
		Data:      []byte("HEAD: ARC START, ARC PACKET INIT"),
		Signature: "gml",
	}

	encryptedMsg, errMsg := encrypt.EncryptData(pubkey2, arcStart.Data)
	if errMsg != nil {
		fmt.Println("Error encrypting message for public key", pubkey2.N)
	}

	fmt.Println("Encrypted data", string(encryptedMsg))

	signature, errSign := cryptomath.SignData(privkey1, arcStart.Data)

	if errSign != nil {
		fmt.Println(errSign.Error())
	}
	fmt.Println("Data signature", signature)

	// decrypt the data

	decryptedData, errDecrypt := encrypt.DecryptData(privkey2, encryptedMsg)
	if errDecrypt != nil {
		fmt.Println(errDecrypt.Error())
	}
	fmt.Println("Decrypted data", string(decryptedData))
	j, e := json.Marshal(arcStart)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(string(j))
	// demo of JSON marshaling
	var g typ.Packet
	er1 := json.Unmarshal(j, &g)
	fmt.Println(er1)
	fmt.Println(string(g.Data))
}
