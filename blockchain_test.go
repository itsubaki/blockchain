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

		pre := c.Last()
		_, nonce := ProofOfWork(pre)
		new := c.NewBlock(pre.Hash(), nonce)

		fmt.Println(new)
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

		pre := c.Last()
		_, nonce := ProofOfWork(pre)
		c.NewBlock(pre.Hash(), nonce)
	}

	c.NewBlock(c.Last().Hash(), 1000)

	if ValidateChain(c) {
		t.Error("invalid chain")
	}
}
