package core

import (
	"crypto/sha256"
	"fmt"
)

const (
	DefaultMethod  = "dl"
	DefaultChain   = "eth"
	DefaultNetwork = "testnet"
)

// DID is a decentralized identifier.
type DID struct {
	ID      [32]byte
	Method  string
	Chain   string
	Network string
}

// NewDID creates a new DID.
func NewDID(pubKey []byte) *DID {
	// TODO: change to poseidon hash later.
	id := sha256.Sum256(pubKey)
	return &DID{
		ID:      id,
		Method:  DefaultMethod,
		Chain:   DefaultChain,
		Network: DefaultNetwork,
	}
}

// String returns the string representation of the DID.
func (d DID) String() string {
	return fmt.Sprintf("did:%s:%s:%s", d.Method, d.Network, d.Chain)
}
