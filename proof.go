package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"math"
	"strconv"
	"strings"
)

func ProofOfWork(last int) int {
	for nonce := 0; nonce < math.MaxInt64; nonce++ {
		if Validate(last, nonce) {
			return nonce
		}
	}

	panic("hash not found.")
}

func Validate(last, current int) bool {
	str := strconv.Itoa(last) + strconv.Itoa(current)
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

		if !Validate(pre.Nonce, next.Nonce) {
			return false
		}

		pre = next
		index = index + 1
	}

	return true
}
