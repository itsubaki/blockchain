package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"math"
	"strconv"
	"strings"
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
	t := &Transaction{Sender: sender, Recipient: recipient, Amount: amount}
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

func ProofOfWork(b *Block) int {
	hash := b.Hash()
	for nonce := 0; nonce < math.MaxInt64; nonce++ {
		if Validate(hash, nonce) {
			return nonce
		}
	}

	panic("hash not found.")
}

func Validate(preHash string, current int) bool {
	str := preHash + strconv.Itoa(current)
	sha := sha256.Sum256([]byte(str))
	hash := hex.EncodeToString(sha[:])

	if strings.HasPrefix(hash, target) {
		return true
	}

	return false
}

func ValidateChain(chain *BlockChain) bool {
	pre := chain.blocks[0]
	index := 1

	for index < len(chain.blocks) {
		next := chain.blocks[index]
		if pre.Hash() != next.PreHash {
			return false
		}

		if !Validate(pre.Hash(), next.Nonce) {
			return false
		}

		pre = next
		index = index + 1
	}

	return true
}
