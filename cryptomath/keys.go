package cryptomath

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

func GenerateKeys(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

func SignData(privateKey *rsa.PrivateKey, msg []byte) ([]byte, error) {
	hash := sha256.New()
	hash.Write(msg)
	hashed := hash.Sum(nil)

	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, hashed, nil)
	if err != nil {
		return nil, err
	}
	return signature, nil
}

func VerifySignature(publicKey *rsa.PublicKey, msg, signature []byte) error {
	hash := sha256.New()
	hash.Write(msg)
	hashed := hash.Sum(nil)

	err := rsa.VerifyPSS(publicKey, crypto.SHA256, hashed, signature, nil)
	if err != nil {
		return err
	}
	return nil
}
