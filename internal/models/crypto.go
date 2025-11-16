package models

// X25519KeyPair represents a Curve25519 key pair used for ECDH key exchange.
// Keys are stored as base64-encoded strings for JSON serialization.
type X25519KeyPair struct {
	Private string `json:"private"` // Base64-encoded private key (32 bytes)
	Public  string `json:"public"`  // Base64-encoded public key (32 bytes)
}

// Ed25519KeyPair represents an Ed25519 key pair used for digital signatures.
// Keys are stored as base64-encoded strings for JSON serialization.
type Ed25519KeyPair struct {
	Private string `json:"private"` // Base64-encoded private key (64 bytes)
	Public  string `json:"public"`  // Base64-encoded public key (32 bytes)
}

// OneTimePreKey represents a one-time pre-key in the X3DH protocol.
// These keys are used once and then marked as used to prevent replay attacks.
type OneTimePreKey struct {
	KeyID   uint32        `json:"key_id"`   // Unique identifier for this pre-key
	KeyPair X25519KeyPair `json:"key_pair"` // The actual key pair
	Used    bool          `json:"used"`     // Whether this key has been used
}

// SignedPreKey represents a signed pre-key in the X3DH protocol.
// The signature is created using the identity key to prove ownership.
type SignedPreKey struct {
	KeyID     uint32        `json:"key_id"`    // Unique identifier for this pre-key
	KeyPair   X25519KeyPair `json:"key_pair"`  // The actual key pair
	Signature string        `json:"signature"` // Base64-encoded Ed25519 signature
}

// X3DHKeyBundle contains all the keys needed for the X3DH key agreement protocol.
// This is the complete bundle stored locally by each user.
type X3DHKeyBundle struct {
	UserID         string          `json:"user_id"`           // User identifier
	IdentityKey    Ed25519KeyPair  `json:"identity_key"`      // Long-term identity key
	SignedPreKey   SignedPreKey    `json:"signed_pre_key"`    // Signed pre-key (rotated periodically)
	OneTimePreKeys []OneTimePreKey `json:"one_time_pre_keys"` // One-time pre-keys (consumed on use)
}

// SignedPreKeyPublic is the public representation of a signed pre-key.
// Used when sharing keys with other users (private key is excluded).
type SignedPreKeyPublic struct {
	KeyID     uint32 `json:"key_id"`     // Unique identifier for this pre-key
	PublicKey string `json:"public_key"` // Base64-encoded public key
	Signature string `json:"signature"`  // Base64-encoded Ed25519 signature
}

// OneTimePreKeyPublic is the public representation of a one-time pre-key.
// Used when sharing keys with other users (private key is excluded).
type OneTimePreKeyPublic struct {
	KeyID     uint32 `json:"key_id"`     // Unique identifier for this pre-key
	PublicKey string `json:"public_key"` // Base64-encoded public key
}

// PublicBundle is the public key bundle shared with other users in the X3DH protocol.
// This contains only public information needed to establish a shared secret.
type PublicBundle struct {
	UserID         string               `json:"user_id"`                    // User identifier
	IdentityPublic string               `json:"identity_public"`            // Base64-encoded identity public key
	SignedPreKey   SignedPreKeyPublic   `json:"signed_pre_key"`             // Signed pre-key (public)
	OneTimePreKey  *OneTimePreKeyPublic `json:"one_time_pre_key,omitempty"` // Optional one-time pre-key
}

func (b *PublicBundle) Validate() error {

	return nil
}
