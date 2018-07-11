package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strconv"
)

type Block struct {
	Index       int            `json:"index"`
	Timestamp   int64          `json:"timestamp"`
	Transaction []*Transaction `json:"transaction"`
	PreHash     string         `json:"previous_hash"`
	Nonce       int            `json:"nonce"`
}

func (b *Block) Hash() string {
	str := b.PreHash + strconv.Itoa(b.Nonce)
	sha := sha256.Sum256([]byte(str))
	return hex.EncodeToString(sha[:])
}

func (b *Block) String() string {
	bytea, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}

	return string(bytea)
}
