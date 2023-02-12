package core_test

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"testing"

	core "github.com/ilya-korotya/distributed_lab/baby_blockchain"
	"github.com/stretchr/testify/require"
	_ "github.com/stretchr/testify/require"
)

var basicKeyPair = core.NewKeyPair()

func TestPublicKeyEqual(t *testing.T) {
	tests := []struct {
		name string
		x    *core.PublicKey
		y    crypto.PublicKey
		eq   bool
	}{
		{
			name: "x and y are equal",
			x:    basicKeyPair.Public().(*core.PublicKey),
			y:    basicKeyPair.Public(),
			eq:   true,
		},
		{
			name: "x and y are not equal",
			x:    basicKeyPair.Public().(*core.PublicKey),
			y:    core.NewKeyPair().Public(),
			eq:   false,
		},
		{
			name: "x and y have different types",
			x:    basicKeyPair.Public().(*core.PublicKey),
			y: func() crypto.PublicKey {
				pr, _ := ecdsa.GenerateKey(elliptic.P224(), nil)
				return pr.Public()
			},
			eq: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.x.Equal(tt.y)
			require.Equal(t, tt.eq, res)
		})
	}
}
