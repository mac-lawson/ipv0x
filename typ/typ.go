package typ

import "crypto/ecdsa"

type Packet struct {
	Sender    string // Public key of the sender
	Receiver  string // Public key of the receiver
	Data      []byte // Encrypted bit array of data
	Signature string // Digital signature of the sender
}

type KnownKey struct {
	Owner       string
	PublicKey   *ecdsa.PublicKey
	Description string
}
