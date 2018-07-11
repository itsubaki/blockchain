package blockchain

import (
	"time"
)

const target = "0000"

type BlockChain struct {
	blocks       []*Block
	transcations []*Transaction
}

func NewBlockChain() *BlockChain {
	c := &BlockChain{}
	c.NewBlock("genesis block", 100)
	return c
}

func (c *BlockChain) NewBlock(preHash string, nonce int) *Block {
	b := &Block{
		Index:       len(c.blocks) + 1,
		Timestamp:   time.Now().UnixNano(),
		Transaction: c.transcations,
		PreHash:     preHash,
		Nonce:       nonce,
	}

	c.transcations = []*Transaction{}
	c.blocks = append(c.blocks, b)

	return b
}

func (c *BlockChain) NewTransaction(sender, recipient string, amount float64) int {
	t := &Transaction{sender: sender, recipient: recipient, amount: amount}
	c.transcations = append(c.transcations, t)
	return c.Last().Index + 1
}

func (c *BlockChain) Last() *Block {
	return c.blocks[len(c.blocks)-1]
}

func (c *BlockChain) Resolve(d *BlockChain) bool {
	if !ValidateChain(d) {
		return false
	}

	if len(c.blocks) >= len(d.blocks) {
		return false
	}

	c.blocks = d.blocks
	return true
}
