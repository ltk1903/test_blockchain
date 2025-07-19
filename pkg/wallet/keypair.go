package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"crypto/sha256"
	"math/big"
)

func GenerateKeyPair() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
}

func PublicKeyToString(pub *ecdsa.PublicKey) string {
	bytes, _:= x509.MarshalPKIXPublicKey(pub)
	return hex.EncodeToString(bytes)
}

func PublicKeyToAddress(pub *ecdsa.PublicKey) string {
	pubBytes, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return ""
	}
	hash := sha256.Sum256(pubBytes)
	return hex.EncodeToString(hash[:])
}

func Sign(hash []byte, priv *ecdsa.PrivateKey) ([]byte, error) {
	r, s, err := ecdsa.Sign(rand.Reader, priv, hash)
	if err != nil {
		return nil, err
	}
	//merge r & s
	signature := append(r.Bytes(), s.Bytes()...)
	return signature, nil
}

func Verify(hash, signature []byte, pub *ecdsa.PublicKey) bool {
	keyLen := len(signature) / 2
	r := new(big.Int).SetBytes(signature[:keyLen])
	s := new(big.Int).SetBytes(signature[keyLen:])
	return ecdsa.Verify(pub, hash, r, s)
}