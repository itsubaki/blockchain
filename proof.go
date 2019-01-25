package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"math"
	"strconv"
	"strings"
)

const target = "0000"

func Hash(preHash string, preT []*Transaction, nonce int) string {
	src := preHash + strconv.Itoa(nonce)
	for _, t := range preT {
		src = src + t.String()
	}

	sha := sha256.Sum256([]byte(src))
	hash := hex.EncodeToString(sha[:])
	return hash
}

func ProofOfWork(preHash string, preT []*Transaction) (string, int) {
	for nonce := 0; nonce < math.MaxInt64; nonce++ {
		hash, ok := Validate(preHash, preT, nonce)
		if ok {
			return hash, nonce
		}
	}

	panic("hash not found.")
}

func Validate(preHash string, preT []*Transaction, current int) (string, bool) {
	hash := Hash(preHash, preT, current)
	if strings.HasPrefix(hash, target) {
		return hash, true
	}

	return "", false
}

func ValidateChain(chain *BlockChain) bool {
	for i := 1; i < len(chain.blocks); i++ {
		pre := chain.blocks[i-1]
		current := chain.blocks[i]

		if pre.Hash != current.PreHash {
			return false
		}

		if _, ok := Validate(pre.Hash, pre.Transaction, current.Nonce); !ok {
			return false
		}

	}

	return true
}
