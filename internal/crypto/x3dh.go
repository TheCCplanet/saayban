package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"sayban/internal/models"
)

// new key bundle
func GenerateX3DHKeyBundle() (*models.X3DHKeyBundle, error) {
	// Identity key (Ed25519)
	ikPub, iKPriv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to generate identity key: %w", err)
	}
	// Signed Pre-Key (X25519)
	spk, err := GenerateX25519KeyPair()
	if err != nil {
		return nil, err
	}
	// One-Time Pre-Key (X25519)
	opk, err := GenerateX25519KeyPair()
	if err != nil {
		return nil, err
	}

	// Decode the base64 public key before signing
	spkPublicBytes, err := base64.StdEncoding.DecodeString(spk.Public)
	if err != nil {
		return nil, fmt.Errorf("failed to decode signed pre-key public key: %w", err)
	}

	spkSignature := ed25519.Sign(iKPriv, spkPublicBytes)

	return &models.X3DHKeyBundle{
		UserID: "", // TODO: Set UserID when generating bundle for a specific user
		IdentityKey: models.Ed25519KeyPair{
			Private: base64.StdEncoding.EncodeToString(iKPriv),
			Public:  base64.StdEncoding.EncodeToString(ikPub),
		},
		SignedPreKey: models.SignedPreKey{
			KeyID:     0, // TODO: Implement proper key ID generation/rotation
			KeyPair:   *spk,
			Signature: base64.StdEncoding.EncodeToString(spkSignature), // Convert to base64 string for consistency
		},
		OneTimePreKeys: []models.OneTimePreKey{
			{
				KeyID:   0, // TODO: Implement proper key ID generation
				KeyPair: *opk,
				Used:    false,
			},
		},
	}, nil
}
