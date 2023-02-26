package core

import (
	"fmt"

	bjj "github.com/iden3/go-iden3-crypto/babyjub"
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
func NewDID(pubKey *PublicKey) *DID {
	return &DID{
		ID:      pubKey.String(),
		Method:  DefaultMethod,
		Chain:   DefaultChain,
		Network: DefaultNetwork,
	}
}

// String returns the string representation of the DID.
func (d DID) String() string {
	return fmt.Sprintf("did:%s:%s:%s:%s", d.Method, d.Network, d.Chain, d.ID)
}

func (d DID) ResolvePubKey() (*PublicKey, error) {
	pk := &bjj.PublicKey{}
	if err := pk.UnmarshalText([]byte(d.ID)); err != nil {
		return nil, err
	}
	return &PublicKey{pk}, nil
}
