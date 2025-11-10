package crypto

import (
	"crypto/ecdh"
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func GenerateX25519KeyPair() (string, string, error) {
	curve := ecdh.X25519()

	privateKey, err := curve.GenerateKey(rand.Reader)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate private key: %w", err)
	}
	publicKey := privateKey.PublicKey()
	return base64.StdEncoding.EncodeToString(publicKey.Bytes()), base64.StdEncoding.EncodeToString(privateKey.Bytes()), nil
}

func GenerateX25519SharedKey(publicKey string, privateKey string) (string, error) {
	curve := ecdh.X25519()

	publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return "", fmt.Errorf("failed to decode public key: %w", err)
	}

	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to decode private key: %w", err)
	}

	privKey, err := curve.NewPrivateKey(privateKeyBytes)
	if err != nil {
		return "", fmt.Errorf("failed to import private key: %w", err)
	}

	pubKey, err := curve.NewPublicKey(publicKeyBytes)
	if err != nil {
		return "", fmt.Errorf("failed to import private key: %w", err)
	}

	sharedKey, err := privKey.ECDH(pubKey)
	if err != nil {
		return "", fmt.Errorf("failed to generate shared key: %w", err)
	}

	return base64.StdEncoding.EncodeToString(sharedKey), nil
}
