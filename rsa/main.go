package myrsa

import (
	"fmt"
	"math"
	"math/rand"
)

type Uint128 struct {
	size int
	n    [16]byte
}

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
func GeneratePrime() Uint128 {
	prime := rand.Uint64()
	for !IsPrime(prime) {
		prime = rand.Uint64()
	}
	var ret Uint128
	for i := 1; i <= 16; i++ {
		ret.n[i-1] = byte(uint64(prime) & 0xff)
		prime >>= 8
	}
	for i := 16; i >= 1; i-- {
		if ret.n[i-1] > 0 {
			ret.size = i
			break
		}
	}
	if ret.size == 0 {
		ret.size = 1
	}
	return ret
}

func Multiply(a Uint128, b Uint128) Uint128 {

}

/*
  Generate 128 bits rsa keys
*/
func GenerateRsaKey() (publicKey [16]byte, privateKey [16]byte) {
	p, q := GeneratePrime(), GeneratePrime()
	fmt.Println(p, q)
	n := Multiply(p, q)

}
