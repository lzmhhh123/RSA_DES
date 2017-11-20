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

func CountSizeUint128(a Uint128) int {
	ret := 1
	for i := 16; i >= 1; i-- {
		if a.n[i-1] > 0 {
			ret = i
			break
		}
	}
	return ret
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
	ret.size = CountSizeUint128(ret)
	return ret
}

func Multiply(a Uint128, b Uint128) Uint128 {
	var tmp [16]int
	for i := 1; i <= a.size; i++ {
		for j := 1; j <= b.size; j++ {
			tmp[i+j-1] += int(a.n[i]) * int(b.n[j])
		}
	}
	var c Uint128
	for i := 1; i <= 16; i++ {
		c.n[i-1] = byte(tmp[i-1] & 0xff)
		tmp[i] += tmp[i-1] / 0x100
	}
	c.size = CountSizeUint128(c)
	return c
}

/*
  Generate 128 bits rsa keys
*/
func GenerateRsaKey() (publicKey [16]byte, privateKey [16]byte) {
	p, q := GeneratePrime(), GeneratePrime()
	fmt.Println(p, q)
	n := Multiply(p, q)

}
