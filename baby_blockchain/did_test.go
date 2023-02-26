package core_test

import (
	"crypto"
	"fmt"
	"testing"

	bjj "github.com/iden3/go-iden3-crypto/babyjub"
	core "github.com/ilya-korotya/distributed_lab/baby_blockchain"
	"github.com/stretchr/testify/require"
)

var (
	// generating by hands for testing
	testingPK  = bjj.PrivateKey([32]byte{168, 140, 75, 205, 254, 48, 108, 255, 155, 247, 142, 64, 171, 137, 29, 147, 142, 40, 79, 195, 190, 95, 140, 232, 98, 131, 248, 41, 89, 74, 70, 151})
	testingPUB = "f98147ef3a5e9e38d8d99b8a6635f020f079fa4cee783ca0ddb9c4c38414011f"
)

func TestNewDID(t *testing.T) {
	testingKeyPair := core.NewKeyPairFromPk(&testingPK)

	tests := []struct {
		name   string
		pubKey crypto.PublicKey
		want   *core.DID
	}{
		{
			name:   "New DID",
			pubKey: testingKeyPair.Public(),
			want: &core.DID{
				ID:      testingPUB,
				Method:  core.DefaultMethod,
				Chain:   core.DefaultChain,
				Network: core.DefaultNetwork,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := core.NewDID(testingKeyPair.Public().(*core.PublicKey))
			require.Equal(t, tt.want.String(), got.String())

			fmt.Printf("did: %s for public key: %s", got.String(), testingPUB)
		})
	}
}

func TestResolvePubKey(t *testing.T) {
	testingKeyPair := core.NewKeyPairFromPk(&testingPK)
	did := core.NewDID(testingKeyPair.Public().(*core.PublicKey))
	pubkey, err := did.ResolvePubKey()
	require.NoError(t, err)
	require.Equal(t, pubkey, testingKeyPair.Public().(*core.PublicKey))
}
