package core

import (
	"encoding/json"
	"math/big"

	"github.com/iden3/go-iden3-crypto/poseidon"
)

type Transaction struct {
	ID         string       `json:"id"`
	Operations []*Operation `json:"operations"`
	Nonce      *big.Int     `json:"nonce"`
}

func NewTransaction(opts []*Operation, nonce *big.Int) (*Transaction, error) {
	rawID := struct {
		Operations []*Operation `json:"operations"`
		Nonce      *big.Int     `json:"nonce"`
	}{
		Operations: opts,
		Nonce:      nonce,
	}
	idBytes, err := json.Marshal(rawID)
	if err != nil {
		return nil, err
	}
	bi, err := poseidon.HashBytes(idBytes)
	if err != nil {
		return nil, err
	}
	return &Transaction{
		ID:         bi.String(),
		Operations: opts,
		Nonce:      nonce,
	}, nil
}
