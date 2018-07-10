package quasar

import (
	"fmt"
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
		b := c.NewBlock(preHash, proof)

		fmt.Println(b)
	}

	if !ValidateChain(c) {
		t.Error("invalid chain")
	}

}
