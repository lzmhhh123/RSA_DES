package main

import (
	"ISpj1-2/rsa"
	"fmt"
)

func main() {
	n, publicKey, privateKey := myrsa.GenerateRsaKey()
	fmt.Println(n, publicKey, privateKey)
}
