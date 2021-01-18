package blockchain

import (
	"time"
)

type BlockChain struct {
	blocks       []Block
	transactions []Transaction
}

func New() *BlockChain {
	c := &BlockChain{
		blocks:       make([]Block, 0),
		transactions: make([]Transaction, 0),
	}

	preHash := "genesis block"
	hash, nonce := ProofOfWork(preHash, make([]Transaction, 0))
	c.NewBlock(preHash, hash, nonce)

	return c
}

func (c *BlockChain) NewBlock(preHash, hash string, nonce int) Block {
	b := Block{
		Timestamp:   time.Now().UnixNano(),
		Transaction: c.transactions,
		Hash:        hash,
		PreHash:     preHash,
		Nonce:       nonce,
	}

	c.transactions = make([]Transaction, 0)
	c.blocks = append(c.blocks, b)

	return b
}

func (c *BlockChain) NewTransaction(sender, recipient string, amount float64) {
	t := Transaction{Sender: sender, Recipient: recipient, Amount: amount}
	c.transactions = append(c.transactions, t)
}

func (c *BlockChain) Transactions() []Transaction {
	return c.transactions
}

func (c *BlockChain) Blocks() []Block {
	return c.blocks
}

func (c *BlockChain) Last() Block {
	return c.blocks[len(c.blocks)-1]
}

func (c *BlockChain) Resolve(d *BlockChain) bool {
	if !d.Validate() {
		return false
	}

	if len(c.blocks) >= len(d.blocks) {
		return false
	}

	c.blocks = d.blocks
	return true
}

func (c *BlockChain) Validate() bool {
	for i := 1; i < len(c.blocks); i++ {
		pre := c.blocks[i-1]
		current := c.blocks[i]

		if pre.Hash != current.PreHash {
			return false
		}

		if _, ok := Validate(pre.Hash, pre.Transaction, current.Nonce); !ok {
			return false
		}
	}

	return true
}
