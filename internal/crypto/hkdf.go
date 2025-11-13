package crypto

import (
	"crypto/sha256"

	"golang.org/x/crypto/hkdf"
	_ "golang.org/x/crypto/hkdf"
)

func DeriveRootKey(sharedSecret [32]byte) ([32]byte, error) {
	var rootKey [32]byte

	hkdfReader := hkdf.New(sha256.New, sharedSecret[:],
		[]byte("x3dh_root_salt_v1"),
		[]byte("X3DH-Root-Key"),
	)

	_, err := hkdfReader.Read(rootKey[:])
	return rootKey, err
}
