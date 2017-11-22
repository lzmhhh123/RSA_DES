package myrsa

import (
	"math"
	"math/big"
	"math/rand"
)

/*
  Judge the number whether prime
*/
func IsPrime(number uint64) bool {
	lim := math.Sqrt(float64(number))
	for i := uint64(2); float64(i) <= lim; i++ {
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
	prime := rand.Uint64()
	for !IsPrime(prime) {
		prime = rand.Uint64()
	}
	ret := big.NewInt(0)
	ret.SetUint64(prime)
	return ret
}

/*
	exgcd algorithm
*/
func exgcd(a, b *big.Int) (x, y *big.Int) {
	tmp := big.NewInt(0)
	if b.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(1), big.NewInt(0)
	}
	tx, ty := exgcd(b, tmp.Mod(a, b))
	x = ty
	y = tmp.Sub(tx, tmp.Mul(tmp.Div(a, b), ty))
	return
}

/*
  Generate 128 bits rsa keys
*/
func GenerateRsaKey() (*big.Int, *big.Int, *big.Int) {
	p, q := GeneratePrime(), GeneratePrime()
	n := big.NewInt(0)
	n.Mul(p, q)
	tmp, fiN := big.NewInt(0), big.NewInt(0)
	fiN.Mul(tmp.Sub(p, big.NewInt(1)), fiN.Sub(q, big.NewInt(1)))
	e := big.NewInt(65537)
	d, _ := exgcd(e, fiN)
	if d.Cmp(big.NewInt(0)) == -1 {
		d.Add(d, fiN)
	}
	return n, e, d
}
