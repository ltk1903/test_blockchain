package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

type Block struct {
	PrevHash string
	Transaction []*Transaction
	Hash string
}

func NewBlock(prevHash string, txs []*Transaction) *Block {
	block := &Block{
		PrevHash: prevHash,
		Transaction: txs,
	}
	block.Hash = block.calculateHash()
	return block
}

func (b *Block) calculateHash() string {
	copyBlock := *b
	copyBlock.Hash = ""

	blockBytes, _:= json.Marshal(copyBlock)
	hash := sha256.Sum256(blockBytes)
	return hex.EncodeToString(hash[:])
}

