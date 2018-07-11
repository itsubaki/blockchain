package blockchain

import (
	"fmt"
	"testing"
)

func TestMine(t *testing.T) {
	c := NewBlockChain()

	for i := 0; i < 10; i++ {
		c.NewTransaction("alice", "bob", 1.2)
		c.NewTransaction("alice", "bob", 1.4)
		c.NewTransaction("alice", "bob", 1.6)

		preHash := c.Last().Hash
		hash, nonce := ProofOfWork(preHash)
		c.NewBlock(preHash, hash, nonce)
	}

	for _, b := range c.blocks {
		fmt.Println(b)
	}

	if !ValidateChain(c) {
		t.Error("invalid chain")
	}
}

func TestInvalidBlock(t *testing.T) {
	c := NewBlockChain()

	for i := 0; i < 10; i++ {
		c.NewTransaction("alice", "bob", 1.2)
		c.NewTransaction("alice", "bob", 1.4)
		c.NewTransaction("alice", "bob", 1.6)

		preHash := c.Last().Hash
		hash, nonce := ProofOfWork(preHash)
		c.NewBlock(preHash, hash, nonce)
	}

	c.NewBlock(c.Last().Hash, "foo", 1000)

	if ValidateChain(c) {
		t.Error("invalid chain")
	}
}
