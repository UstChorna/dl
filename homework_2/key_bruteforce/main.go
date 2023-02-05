package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)


// KeyLens defines the key length options
var	KeyLens = []int64{8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096}


// GetKeySpace returns the key space for a given key length
func GetKeySpace(key int64) *big.Int {
	return big.NewInt(0).Exp(
		big.NewInt(2),
		big.NewInt(key),
		nil,
	)
}

// PrintKeySpace prints the key space for each key length in KeyLens
func PrintKeySpace() {
	for _, key := range KeyLens {
		fmt.Printf("Key space for the key length '%d' is: %d\n",
			key, GetKeySpace(key),
		)
	}
}

// BruteForceKey performs a brute force search for a key of the given length
func BruteForceKey(key int64) error {
	keySpace := GetKeySpace(key)
	randKey, err := rand.Int(rand.Reader, keySpace)
	if err != nil {
		return err
	}
	i := big.NewInt(0)
	for i.Cmp(randKey) != 0 {
		i.Add(i, big.NewInt(1))
	}
	return nil
}

func main() {
	PrintKeySpace()

	for _, key := range KeyLens {
		t := time.Now()
		err := BruteForceKey(key)
		if err != nil {
			fmt.Printf("Failed to brute force key length '%d': %v", key, err)
			continue
		}
		fmt.Printf("For brute force key length '%d', we spent '%.8fs'\n",
			key, time.Since(t).Seconds(),
		)
	}
}
