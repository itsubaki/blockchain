package blockchain_test

import (
	"fmt"
	"testing"

	"github.com/itsubaki/blockchain"
)

func TestMine(t *testing.T) {
	bc := blockchain.New()

	for i := 0; i < 10; i++ {
		bc.NewTransaction("alice", "bob", 1.2)
		bc.NewTransaction("alice", "bob", 1.4)
		bc.NewTransaction("alice", "bob", 1.6)

		preHash := bc.Last().Hash
		hash, nonce := blockchain.ProofOfWork(preHash, bc.Last().Transaction)
		bc.NewBlock(preHash, hash, nonce)
	}

	for _, b := range bc.Blocks() {
		fmt.Println(b)
	}

	if !bc.Validate() {
		t.Error("invalid chain")
	}
}

func TestInvalidBlock(t *testing.T) {
	bc := blockchain.New()

	for i := 0; i < 10; i++ {
		bc.NewTransaction("alice", "bob", 1.2)
		bc.NewTransaction("alice", "bob", 1.4)
		bc.NewTransaction("alice", "bob", 1.6)

		preHash := bc.Last().Hash
		hash, nonce := blockchain.ProofOfWork(preHash, bc.Last().Transaction)
		bc.NewBlock(preHash, hash, nonce)
	}

	bc.NewBlock(bc.Last().Hash, "foo", 1000)

	if bc.Validate() {
		t.Error("invalid chain")
	}
}
