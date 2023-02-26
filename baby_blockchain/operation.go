package core

import (
	"encoding/json"
)

type OperationType int

const (
	TransactionTypeIssuedClaim OperationType = iota
)

type Operation struct {
	Issuer *DID          `json:"issuer"` // sender
	User   *DID          `json:"user"`   // receiver
	Proof  *Proof        `json:"proof"`
	Type   OperationType `json:"type"`

	Signature []byte `json:"-"`
}

func (o *Operation) Sign(keyPair *KeyPair) error {
	operationBytes, err := json.Marshal(o)
	if err != nil {
		return err
	}
	signature, err := SignMessage(keyPair.Private(), operationBytes)
	if err != nil {
		return err
	}
	o.Signature = signature
	return nil
}

func VerifyOperation(o *Operation) error {
	senderPubKey, err := o.Issuer.ResolvePubKey()
	if err != nil {
		return err
	}
	operationBytes, err := json.Marshal(o)
	if err != nil {
		return err
	}
	if err := VerifyMessage(senderPubKey, operationBytes, o.Signature); err != nil {
		return err
	}
	return nil
}
