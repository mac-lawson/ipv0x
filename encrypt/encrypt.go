package encrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

func EncryptData(publicKey *rsa.PublicKey, msg []byte) ([]byte, error) {
	label := []byte("")
	hash := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, publicKey, msg, label)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}
