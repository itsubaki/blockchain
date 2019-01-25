package blockchain

import (
	"time"
)

type BlockChain struct {
	blocks       []*Block
	transactions []*Transaction
}

func NewBlockChain() *BlockChain {
	c := &BlockChain{
		blocks:       []*Block{},
		transactions: []*Transaction{},
	}

	preHash := "genesis block"
	hash, nonce := ProofOfWork(preHash, []*Transaction{})
	c.NewBlock(preHash, hash, nonce)

	return c
}

func (c *BlockChain) NewBlock(preHash, hash string, nonce int) *Block {
	b := &Block{
		Timestamp:   time.Now().UnixNano(),
		Transaction: c.transactions,
		Hash:        hash,
		PreHash:     preHash,
		Nonce:       nonce,
	}

	c.transactions = []*Transaction{}
	c.blocks = append(c.blocks, b)

	return b
}

func (c *BlockChain) NewTransaction(sender, recipient string, amount float64) {
	t := &Transaction{Sender: sender, Recipient: recipient, Amount: amount}
	c.transactions = append(c.transactions, t)
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
