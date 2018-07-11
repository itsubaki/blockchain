package blockchain

import (
	"testing"
)

func TestMine(t *testing.T) {
	c := NewBlockChain()

	for i := 0; i < 10; i++ {
		nonce := ProofOfWork(c.Last().Nonce)

		c.NewTransaction("sender", "recipient", 1.2)
		c.NewTransaction("sender", "recipient", 1.4)
		c.NewTransaction("sender", "recipient", 1.6)

		preHash := c.Last().Hash()
		c.NewBlock(preHash, nonce)
	}

	if !ValidateChain(c) {
		t.Error("invalid chain")
	}
}

func TestInvalidBlock(t *testing.T) {
	c := NewBlockChain()

	for i := 0; i < 10; i++ {
		nonce := ProofOfWork(c.Last().Nonce)

		c.NewTransaction("sender", "recipient", 1.2)
		c.NewTransaction("sender", "recipient", 1.4)
		c.NewTransaction("sender", "recipient", 1.6)

		preHash := c.Last().Hash()
		c.NewBlock(preHash, nonce)
	}

	preHash := c.Last().Hash()
	nonce := 1000 // invalid nonce
	c.NewBlock(preHash, nonce)

	if ValidateChain(c) {
		t.Error("invalid chain")
	}
}
