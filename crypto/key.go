package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
)

func NormalizePublicKey(address string) (string, error) {
	address = strings.TrimSpace(address)
	addr, err := hex.DecodeString(address)
	if err != nil {
		return "", fmt.Errorf("Near invalid address %s", address)
	}
	if len(addr) != ed25519.PublicKeySize {
		return "", fmt.Errorf("Near invalid address %s", address)
	}
	a := hex.EncodeToString(addr)
	if a != address {
		return "", fmt.Errorf("Near invalid address %s", address)
	}
	return a, nil
}

func GenerateKey() (string, string, error) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return "", "", err
	}
	private := hex.EncodeToString(priv.Seed())
	address := hex.EncodeToString(pub)
	return address, private, nil
}
