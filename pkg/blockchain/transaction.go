package blockchain

import (
	"time"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"

	"test_blockchain/pkg/wallet"
)

type Transaction struct {
	From string
	To string
	Amount float64
	Timestamp int64
	Signature string
}

// type Block struct {
// 	Index int
// 	Timestamp time.Time
// 	Transaction []Transaction
// 	PrevHash string
// 	Hash string
// 	MerkleRoot string
// 	Validator string
// 	Signature string
// 

func NewTransaction(from, to string, amount float) *Transaction {
	return &Transaction{
		From: from,
		To: to,
		Amount: amount,
		Timestamp: time.Now().Unix(),
	}
}

func(tx *Transaction) Hash() []byte {
	copyTx := *tx
	copyTx.Signature = nil
	bytes, _:= json.Marshal(copyTx)
	hash := sha256.Sum256(bytes)
	return hash[:]
}

func (tx *Transaction) Sign(priv *ecdsa.PrivateKey) error {
	hash := tx.Hash()
	sig, err := wallet.Sign(hash, priv)
	if err != nil {
		return err
	}
	tx.Signature = sig
	return nil
}

func (tx *Transaction) Verify(pub *ecdsa.PublicKey) bool {
	return wallet.Verify(tx.Hash(), tx.Signature, pub)
}