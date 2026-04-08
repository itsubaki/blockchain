package blockchain

import (
	"time"
)

type BlockChain struct {
	blocks       []*Block
	transactions []*Transaction
}

func New() *BlockChain {
	chain := &BlockChain{
		blocks:       make([]*Block, 0),
		transactions: make([]*Transaction, 0),
	}

	preHash := "genesis block"
	hash, nonce := ProofOfWork(preHash, make([]*Transaction, 0))
	chain.NewBlock(preHash, hash, nonce)
	return chain
}

func (c *BlockChain) NewBlock(preHash, hash string, nonce int) *Block {
	block := &Block{
		Timestamp:   time.Now().UnixNano(),
		Transaction: c.transactions,
		Hash:        hash,
		PreHash:     preHash,
		Nonce:       nonce,
	}

	c.transactions = make([]*Transaction, 0)
	c.blocks = append(c.blocks, block)
	return block
}

func (c *BlockChain) NewTransaction(sender, recipient string, amount float64) {
	c.transactions = append(c.transactions, &Transaction{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
	})
}

func (c *BlockChain) Transactions() []*Transaction {
	return c.transactions
}

func (c *BlockChain) Blocks() []*Block {
	return c.blocks
}

func (c *BlockChain) Last() *Block {
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
		prev := c.blocks[i-1]
		curr := c.blocks[i]
		if prev.Hash != curr.PreHash {
			return false
		}

		if _, ok := Validate(prev.Hash, prev.Transaction, curr.Nonce); !ok {
			return false
		}
	}

	return true
}
