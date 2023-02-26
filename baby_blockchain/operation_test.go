package core_test

import (
	"testing"

	core "github.com/ilya-korotya/distributed_lab/baby_blockchain"
	"github.com/stretchr/testify/require"
)

func TestVerifyOperation(t *testing.T) {
	issuerKeyPair := core.NewKeyPair()
	userKeyPair := core.NewKeyPair()
	o := &core.Operation{
		Issuer: core.NewDID(issuerKeyPair.Public().(*core.PublicKey)),
		User:   core.NewDID(userKeyPair.Public().(*core.PublicKey)),
		Proof:  &core.Proof{},
		Type:   core.TransactionTypeIssuedClaim,
	}
	err := o.Sign(issuerKeyPair)
	require.NoError(t, err)
	err = core.VerifyOperation(o)
	require.NoError(t, err)
}
