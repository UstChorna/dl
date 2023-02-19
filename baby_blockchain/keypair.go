// Currently, this is very abstract keypair. I should return implementation not interface.

package core

import (
	"crypto"

	bjj "github.com/iden3/go-iden3-crypto/babyjub"
)

// KeyPair for bjj keys that implements go interfaces.
type KeyPair struct {
	private crypto.Signer
	public  crypto.PublicKey
}

// NewKeyPair returns a new key pair.
func NewKeyPair() *KeyPair {
	pk := bjj.NewRandPrivKey()
	pub := pk.Public()
	return &KeyPair{&PrivateKey{pk}, &PublicKey{pub}}
}

// NewKeyPairFromPk returns a new key pair from private key.
func NewKeyPairFromPk(pk *bjj.PrivateKey) *KeyPair {
	pub := pk.Public()
	return &KeyPair{&PrivateKey{*pk}, &PublicKey{pub}}
}

// Private returns private key from key pair.
func (kp *KeyPair) Private() crypto.Signer {
	return kp.private
}

// Public returns public key from key pair.
func (kp *KeyPair) Public() crypto.PublicKey {
	return kp.public
}

func (kp *KeyPair) PublicToBytes() []byte {
	raw := kp.public.(*PublicKey).Compress()
	return raw[:]
}
