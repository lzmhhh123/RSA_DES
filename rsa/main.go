package myrsa

import (
	"math"
	"math/big"
	"math/rand"
)

/*
  Judge the number whether prime
*/
func IsPrime(number int64) bool {
	lim := math.Sqrt(float64(number))
	for i := int64(2); float64(i) <= lim; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

/*
  Generate prime randly
*/
func GeneratePrime() *big.Int {
	prime := rand.Int63()
	for !IsPrime(prime) {
		prime = rand.Int63()
	}
	ret := big.NewInt(prime)
	return ret
}

/*
  Generate 128 bits rsa keys
*/
func GenerateRsaKey() (n *big.Int, publicKey *big.Int, privateKey *big.Int) {
	p, q := GeneratePrime(), GeneratePrime()
	var tmp *big.Int
	n = tmp.Mul(p, q)
	fiN := tmp.Mul(tmp.Sub(p, big.NewInt(1)), tmp.Sub(q, big.NewInt(1)))
	return n, p, q
}
