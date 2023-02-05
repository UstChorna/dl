package main

import (
	cr "crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"
)

/* TODO:
1. Create PubKey struct with {e, n}
2. Create PrivKey struct with {d, n}
3. Encode and Decody raw bytes instead of big.Ints.
Example:
func (private *priKey) decrypt(ciptext []rune) []rune {
	for _, r := range ciptext {
		big.NewInt(0).SetString(r, 2)
		bigB.Exp(bigB, private.d, private.n)
		tmp = append(tmp, int(bigB.Int64()))
		// or use Unicode?
	}
	...
}
*/

const keySize = 2048

func generatePQ() (*big.Int, *big.Int) {
	Q, _ := cr.Prime(cr.Reader, keySize)
	P, _ := cr.Prime(cr.Reader, keySize)
	return Q, P
}

func findE(v *big.Int) (e *big.Int) {
	e = big.NewInt(2)
	for {
		nod := big.NewInt(1).GCD(nil, nil, e, v)
		if nod.Cmp(big.NewInt(1)) == 0 {
			return e
		}
		e = e.Add(e, big.NewInt(1))
	}
}

func calPrivate(e *big.Int, totient *big.Int) *big.Int {
	return big.NewInt(0).ModInverse(e, totient)
}

func getKey() (partKey *big.Int, pubKey *big.Int, privKey *big.Int) {
	P, Q := generatePQ()
	n := big.NewInt(0).Mul(P, Q)
	P = big.NewInt(0).Sub(P, big.NewInt(1))
	Q = big.NewInt(0).Sub(Q, big.NewInt(1))
	totient := big.NewInt(0).Mul(P, Q)
	e := findE(totient)
	if e.Cmp(totient) == 1 {
		return nil, nil, nil
	}

	d := calPrivate(e, totient)

	return n, e, d
}

func main() {
	partyKey, pubKey, privKey := getKey()
	if privKey == nil {
		log.Fatal("math is broken :)")
	}

	msg, _ := big.NewInt(0).SetString(os.Args[1], 10)

	fmt.Println("original message:", msg)
	encodedMsg := big.NewInt(0).Exp(msg, pubKey, partyKey)
	fmt.Println("encoded message:", encodedMsg)
	decodedMsg := big.NewInt(0).Exp(encodedMsg, privKey, partyKey)
	fmt.Println("decoded message:", decodedMsg)
}
