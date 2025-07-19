package main

import (
	"fmt"
	"simple-blockchain/pkg/blockchain"
	"simple-blockchain/pkg/wallet"
)

func main() {
	// Tạo ví cho Alice và Bob
	alicePriv, _:= wallet.GenerateKeyPair()
	bobPriv, _:= wallet.GenerateKeyPair()

	aliceAddr := wallet.PublicKeyToAddress(&alicePriv.PublicKey)
	bobAddr := wallet.PuclicKeyToAddress(&bobPriv.PublicKey)

	fmt.Println("Alice address:", aliceAddr)
	fmt.Println("Bob address:", bobAddr)

	//Alice chuyển 10 coin cho Bob
	tx := blockchain.NewTransaction(aliceAddr, bobAddr, 10)
	tx := Sign(alicePriv)

	//Xác minh giao dịch
	isValid := tx.Verify(&alicePriv.PublicKey)
	fmt.Println("Transaction valid:", isValid)

	//Tạo block với 1 transaction
	block := blockchain.NewBlock("", []*blockchain.Transaction{tx})
	fmt.Println("Block hash:", block.Hash)
}