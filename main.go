package main

import (
	"ISpj1-2/myrsa"
	"crypto/cipher"
	"crypto/des"
	"encoding/binary"
	"fmt"
)

func ReadDesKey() []byte {
	fmt.Println("Please enter the DES key by a 64bits integer: ")
	var desNumber uint64
	fmt.Scanln(&desNumber)
	desKey := make([]byte, 8)
	binary.LittleEndian.PutUint64(desKey, uint64(desNumber))
	return desKey
}

func main() {
	n, publicKey, privateKey := myrsa.GenerateRsaKey()
	fmt.Println("RSA public key: (", n, ", ", publicKey, ").")
	fmt.Println("RSA private key: (", n, ", ", privateKey, ").")
	desKey := ReadDesKey()
	cipherBlock, err := des.NewCipher(desKey)
	if err != nil {
		panic(err)
	}
	cipherBlockMode := cipher.NewCBCDecrypter(cipherBlock, make([]byte, 8))

}
