package blockchain

import "encoding/json"

type Transaction struct {
	Sender    string  `json:"sender"`
	Recipient string  `json:"recipient"`
	Amount    float64 `json:"amount"`
}

func (t *Transaction) String() string {
	bytea, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}

	return string(bytea)
}
