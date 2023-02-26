package core

import (
	"fmt"

	"github.com/iden3/go-iden3-crypto/poseidon"
)

const (
	DefaultMethod  = "dl"
	DefaultChain   = "eth"
	DefaultNetwork = "testnet"
)

// DID is a decentralized identifier.
type DID struct {
	ID      string
	Method  string
	Chain   string
	Network string
}

// NewDID creates a new DID.
func NewDID(pubKey []byte) (*DID, error) {
	id, err := poseidon.HashBytes(pubKey)
	if err != nil {
		return nil, err
	}
	return &DID{
		ID:      id.Text(16),
		Method:  DefaultMethod,
		Chain:   DefaultChain,
		Network: DefaultNetwork,
	}, nil
}

// String returns the string representation of the DID.
func (d DID) String() string {
	return fmt.Sprintf("did:%s:%s:%s:%s", d.Method, d.Network, d.Chain, d.ID)
}
