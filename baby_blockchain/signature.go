package core

import (
	"crypto"
	"fmt"
	"math/big"

	bjj "github.com/iden3/go-iden3-crypto/babyjub"
)

var (
	// ErrInvalidSignature is returned when message was signed by another private key.
	ErrInvalidSignature = fmt.Errorf("invalid signature")
)

// SignMessage returns signature.
func SignMessage(pk crypto.Signer, msg []byte) ([]byte, error) {
	return pk.Sign(nil, msg, nil)
}

// VerifyMessage verifies the signature.
func VerifyMessage(pub crypto.PublicKey, msg []byte, sig []byte) error {
	switch p := pub.(type) {
	case *PublicKey:
		sigComp := &bjj.SignatureComp{}
		if err := sigComp.UnmarshalText(sig); err != nil {
			return fmt.Errorf("invalid compressed signature format: %w", err)
		}
		bjjSig, err := sigComp.Decompress()
		if err != nil {
			return fmt.Errorf("invalid signature format: %w", err)
		}
		if p.VerifyPoseidon(big.NewInt(0).SetBytes(msg), bjjSig) {
			return nil
		}
		return ErrInvalidSignature
	}
	return fmt.Errorf("invalid public key type")
}
