package blockchain

import "encoding/json"

type Block struct {
	Timestamp   int64          `json:"timestamp"`
	Transaction []*Transaction `json:"transaction"`
	Hash        string         `json:"hash"`
	PreHash     string         `json:"previous_hash"`
	Nonce       int            `json:"nonce"`
}

func (b *Block) String() string {
	bytea, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}

	return string(bytea)
}
