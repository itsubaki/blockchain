package blockchain

import (
	"testing"
)

func TestMine(t *testing.T) {
	c := NewBlockChain()

	for i := 0; i < 10; i++ {
		proof := ProofOfWork(c.Last().Proof)

		c.NewTransaction("sender", "recipient", 1.2)
		c.NewTransaction("sender", "recipient", 1.4)
		c.NewTransaction("sender", "recipient", 1.6)

		preHash := c.Last().Hash()
		c.NewBlock(preHash, proof)
	}

	if !ValidateChain(c) {
		t.Error("invalid chain")
	}
}

func TestInvalidBlock(t *testing.T) {
	c := NewBlockChain()

	for i := 0; i < 10; i++ {
		proof := ProofOfWork(c.Last().Proof)

		c.NewTransaction("sender", "recipient", 1.2)
		c.NewTransaction("sender", "recipient", 1.4)
		c.NewTransaction("sender", "recipient", 1.6)

		preHash := c.Last().Hash()
		c.NewBlock(preHash, proof)
	}

	preHash := c.Last().Hash()
	proof := 1000 // invalid proof
	c.NewBlock(preHash, proof)

	if ValidateChain(c) {
		t.Error("invalid chain")
	}
}
