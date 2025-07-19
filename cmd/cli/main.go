package main

import (
	"fmt"
	"test_blockchain/pkg/blockchain"
	"test_blockchain/pkg/wallet"
)

func main() {
	// Tạo ví cho Alice và Bob
	alicePriv, _ := wallet.GenerateKeyPair()
	bobPriv, _ := wallet.GenerateKeyPair()

	aliceAddr := wallet.PublicKeyToAddress(&alicePriv.PublicKey)
	bobAddr := wallet.PublicKeyToAddress(&bobPriv.PublicKey)

	fmt.Println("Alice Address:", aliceAddr)
	fmt.Println("Bob Address:  ", bobAddr)

	// Alice chuyển 10 coin cho Bob
	tx := blockchain.NewTransaction(aliceAddr, bobAddr, 10)
	tx.Sign(alicePriv)

	// Xác minh giao dịch
	isValid := tx.Verify(&alicePriv.PublicKey)
	fmt.Println("Transaction valid:", isValid)

	// Tạo block với 1 transaction
	block := blockchain.NewBlock("", []*blockchain.Transaction{tx})
	fmt.Println("Block Hash:", block.Hash)
}
