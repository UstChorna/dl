package core_test

import (
	"testing"

	core "github.com/ilya-korotya/distributed_lab/baby_blockchain"
	"github.com/stretchr/testify/require"
)

func TestVerifyMessage(t *testing.T) {
	keyPair := core.NewKeyPair()
	msg := []byte("Crypto future is here!")
	sig, err := core.SignMessage(keyPair.Private(), msg)
	require.NoError(t, err)
	require.NotNil(t, sig)

	err = core.VerifyMessage(keyPair.Public(), msg, sig)
	require.NoError(t, err)
}

func TestVerifyMessageInvalid(t *testing.T) {
	AliceKeyPair := core.NewKeyPair()
	BobKeyPair := core.NewKeyPair()

	msg := []byte("Hi, I'm Alice")
	sig, err := core.SignMessage(BobKeyPair.Private(), msg)
	require.NoError(t, err)
	require.NotNil(t, sig)

	err = core.VerifyMessage(AliceKeyPair.Public(), msg, sig)
	require.ErrorIs(t, err, core.ErrInvalidSignature)
}
