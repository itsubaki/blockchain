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

	preHash := "genesis block"
	hash, nonce := ProofOfWork(preHash)
	c.NewBlock(preHash, hash, nonce)
	return c
}

func (c *BlockChain) NewBlock(preHash, hash string, nonce int) *Block {
	b := &Block{
		Index:       len(c.blocks) + 1,
		Timestamp:   time.Now().UnixNano(),
		Transaction: c.transcations,
		Hash:        hash,
		PreHash:     preHash,
		Nonce:       nonce,
	}

	c.transcations = []*Transaction{}
	c.blocks = append(c.blocks, b)

	return b
}

func (c *BlockChain) NewTransaction(sender, recipient string, amount float64) {
	t := &Transaction{Sender: sender, Recipient: recipient, Amount: amount}
	c.transcations = append(c.transcations, t)
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
