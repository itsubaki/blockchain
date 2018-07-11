package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"math"
	"strconv"
	"strings"
)

func Hash(preHash string, nonce int) string {
	str := preHash + strconv.Itoa(nonce)
	sha := sha256.Sum256([]byte(str))
	return hex.EncodeToString(sha[:])
}

func ProofOfWork(preHash string) (string, int) {
	for nonce := 0; nonce < math.MaxInt64; nonce++ {
		hash, ok := Validate(preHash, nonce)
		if ok {
			return hash, nonce
		}
	}

	panic("hash not found.")
}

func Validate(preHash string, current int) (string, bool) {
	str := Hash(preHash, current)
	sha := sha256.Sum256([]byte(str))
	hash := hex.EncodeToString(sha[:])

	if strings.HasPrefix(hash, target) {
		return hash, true
	}

	return "", false
}

func ValidateChain(chain *BlockChain) bool {
	pre := chain.blocks[0]
	index := 1

	for index < len(chain.blocks) {
		next := chain.blocks[index]
		if pre.Hash != next.PreHash {
			return false
		}

		_, ok := Validate(pre.Hash, next.Nonce)
		if !ok {
			return false
		}

		pre = next
		index = index + 1
	}

	return true
}
