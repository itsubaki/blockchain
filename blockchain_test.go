package blockchain_test

import (
	"fmt"

	"github.com/itsubaki/blockchain"
)

func Example() {
	chain := blockchain.New()

	for range 3 {
		chain.NewTransaction("alice", "bob", 1.2)
		chain.NewTransaction("alice", "bob", 1.4)
		chain.NewTransaction("alice", "bob", 1.6)

		preHash := chain.Last().Hash
		hash, nonce := blockchain.ProofOfWork(preHash, chain.Last().Transaction)
		chain.NewBlock(preHash, hash, nonce)
	}

	for _, b := range chain.Blocks() {
		fmt.Println(b.Nonce, b.PreHash, b.Hash)
	}

	if !chain.Validate() {
		panic("invalid")
	}

	// Output:
	// 25017 genesis block 0000ff2103db684dd59a277d347e8d62c1d9f310795541dd6f4be5aefdccd7b6
	// 40877 0000ff2103db684dd59a277d347e8d62c1d9f310795541dd6f4be5aefdccd7b6 0000f32d5f8c20d461406d02cc6d2e83d17ad597b0fe5d02e268640657d6aa27
	// 50522 0000f32d5f8c20d461406d02cc6d2e83d17ad597b0fe5d02e268640657d6aa27 0000095b683f6745214508874bc8cb5690f8b064a7c0607cf1bfedc19ddbf0e5
	// 12247 0000095b683f6745214508874bc8cb5690f8b064a7c0607cf1bfedc19ddbf0e5 0000ebffb63c10d65b085e71d1912286c068d0c289f924a43b21c8e7bcb9e520
}

func Example_invalid() {
	chain := blockchain.New()

	for range 3 {
		chain.NewTransaction("alice", "bob", 1.2)
		chain.NewTransaction("alice", "bob", 1.4)
		chain.NewTransaction("alice", "bob", 1.6)

		preHash := chain.Last().Hash
		hash, nonce := blockchain.ProofOfWork(preHash, chain.Last().Transaction)
		chain.NewBlock(preHash, hash, nonce)
	}

	chain.NewBlock(chain.Last().Hash, "foo", 1000)
	if !chain.Validate() {
		fmt.Println("invalid")
	}

	// Output:
	// invalid
}
