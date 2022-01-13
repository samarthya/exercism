package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// Diffie-Hellman-Merkle key exchange
// Private keys should be generated randomly.

// generate random privateKey in (1, p)
func PrivateKey(p *big.Int) *big.Int {
	diff := new(big.Int).Sub(p, big.NewInt(2))
	// rand.Int(n) returns a number [0, n).
	randomKey, _ := rand.Int(rand.Reader, diff)
	return randomKey.Add(randomKey, big.NewInt(2))
}

// generate public key: g ** a mod q
func PublicKey(private, p *big.Int, g int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(g), private, p)
}

func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	public = PublicKey(private, p, g)
	return
}

// generate secretKey by public2 ** private1 mod p
func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}
