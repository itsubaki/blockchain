package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"math"
	"strconv"
	"strings"
	"time"
)

type BlockChain struct {
	blocks       []*Block
	transcations []*Transaction
}

func NewBlockChain() *BlockChain {
	c := &BlockChain{}
	c.NewBlock("genesis block", 100)
	return c
}

func (c *BlockChain) NewBlock(preHash string, proof int) *Block {
	b := &Block{
		Index:       len(c.blocks) + 1,
		Timestamp:   time.Now().UnixNano(),
		Transaction: c.transcations,
		Proof:       proof,
		PreHash:     preHash,
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

func ProofOfWork(last int) int {
	for i := 0; i < math.MaxInt64; i++ {
		if Validate(last, i) {
			return i
		}
	}

	panic("hash not found.")
}

func Validate(last, current int) bool {
	nonce := strconv.Itoa(last) + strconv.Itoa(current)
	sha := sha256.Sum256([]byte(nonce))
	hash := hex.EncodeToString(sha[:])

	if strings.HasPrefix(hash, "0000") {
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

		if !Validate(pre.Proof, next.Proof) {
			return false
		}

		pre = next
		index = index + 1
	}

	return true
}
