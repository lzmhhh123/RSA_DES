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

func exgcd(a *big.Int, b *big.Int) (x *big.Int, y *big.Int) {
	if b == big.NewInt(0) {
		return big.NewInt(1), big.NewInt(0)
	}
	var tmp *big.Int
	tx, ty := exgcd(b, tmp.Mod(a, b))
	x = ty
	y = tmp.Sub(tx, tmp.Mul(tmp.Div(a, b), ty))
	return
}

/*
  Generate 128 bits rsa keys
*/
func GenerateRsaKey() (n *big.Int, publicKey *big.Int, privateKey *big.Int) {
	p, q := GeneratePrime(), GeneratePrime()
	var tmp *big.Int
	n = tmp.Mul(p, q)
	fiN := tmp.Mul(tmp.Sub(p, big.NewInt(1)), tmp.Sub(q, big.NewInt(1)))
	e := big.NewInt(65537)
	d, _ := exgcd(e, fiN)
	return n, e, d
}
