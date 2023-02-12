// This package is a wrapper for the go-iden3-crypto package.
// It implements the crypto.Signer interface.
// It will be help to use the core package in the same way as the go standard library.

// I'm using BJJ keys for signing and verifying messages
// because circom circuits uses less constrains for check signatures.

// Also, this wrapper adds an easy way to switch to other signature schemes.

package core

import (
	"crypto"
	"errors"
	"io"
	"math/big"

	bjj "github.com/iden3/go-iden3-crypto/babyjub"
)

// PublicKey is a wrapper for bjj.PublicKey.
// That add Equal method to implement golang interfaces.
type PublicKey struct {
	*bjj.PublicKey
}

// Equal returns true if x is equal to the public key.
func (pub *PublicKey) Equal(x crypto.PublicKey) bool {
	other, ok := x.(*PublicKey)
	if !ok {
		return false
	}
	return pub.X.Cmp(other.X) == 0 && pub.Y.Cmp(other.Y) == 0
}

// PrivateKey is a wrapper for bjj.PrivateKey.
// That add Sign and Public method to implement golang interfaces.
type PrivateKey struct {
	bjj.PrivateKey
}

// Public returns the public key corresponding to priv.
func (pk *PrivateKey) Public() crypto.PublicKey {
	return &PublicKey{pk.PrivateKey.Public()}
}

// Sign returs EdDSA signature.
func (pk *PrivateKey) Sign(
	_ io.Reader,
	digest []byte,
	opts crypto.SignerOpts,
) ([]byte, error) {
	if opts != nil {
		return nil, errors.New("failed overload hash function")
	}
	msg := big.NewInt(0).SetBytes(digest)
	sig := pk.SignPoseidon(msg)
	return sig.Compress().MarshalText()
}
