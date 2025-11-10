package crypto

import (
	"crypto/ecdh"
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func GenerateDHKeyPair() (string, string, error) {
	privateKey, err := ecdh.X25519().GenerateKey(rand.Reader)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate private key: %w", err)
	}
	publicKey := privateKey.PublicKey()
	return base64.StdEncoding.EncodeToString(publicKey.Bytes()), base64.StdEncoding.EncodeToString(privateKey.Bytes()), nil
}
