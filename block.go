package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

type Block struct {
	Index       int            `json:"index"`
	Timestamp   int64          `json:"timestamp"`
	Transaction []*Transaction `json:"transaction"`
	Proof       int            `json:"proof"`
	PreHash     string         `json:"previous_hash"`
}

func (b *Block) Hash() string {
	bytea, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}

	sha := sha256.Sum256(bytea)
	return hex.EncodeToString(sha[:])
}
