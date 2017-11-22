package main

import (
	"ISpj1-2/myrsa"
	"crypto/cipher"
	"crypto/des"
	"encoding/binary"
	"fmt"
	"math/big"
)

func ReadDesKey() []byte {
	fmt.Println("Please enter the DES key by a 64bits integer: ")
	var desNumber uint64
	fmt.Scanln(&desNumber)
	desKey := make([]byte, 8)
	binary.LittleEndian.PutUint64(desKey, uint64(desNumber))
	return desKey
}

func ReadText() (ret []byte) {
	fmt.Println("Please enter the text to encrypt: ")
	var s string
	fmt.Scanln(&s)
	ret = []byte(s)
	return
}

func Power(a, b, mod *big.Int) (ret *big.Int) {
	ret = big.NewInt(1)
	for b.Cmp(big.NewInt(0)) != 0 {
		tmp := big.NewInt(0)
		if tmp.And(b, big.NewInt(1)); tmp.Cmp(big.NewInt(1)) == 0 {
			ret.Mul(ret, a)
			ret.Mod(ret, mod)
		}
		a.Mul(a, a)
		a.Mod(a, mod)
		b.Div(b, big.NewInt(2))
	}
	return
}

func RsaEncrypt(desKey []byte, rsaPublicKey, n *big.Int) (encryptedDesKey *big.Int) {
	desNumber := binary.LittleEndian.Uint64(desKey)
	desBigNumber := big.NewInt(0)
	desBigNumber.SetUint64(desNumber)
	encryptedDesKey = Power(desBigNumber, rsaPublicKey, n)
	return
}

func RsaDecrypt(desKey, rsaPrivateKey, n *big.Int) (decryptedDesKey *big.Int) {
	decryptedDesKey = Power(desKey, rsaPrivateKey, n)
	return
}

func main() {
	n, publicKey, privateKey := myrsa.GenerateRsaKey()
	fmt.Println("RSA public key: (", n.String(), ", ", publicKey.String(), ").")
	fmt.Println("RSA private key: (", n.String(), ", ", privateKey.String(), ").")
	desKey := ReadDesKey()
	cipherBlock, err := des.NewCipher(desKey)
	encryptedDesKey := RsaEncrypt(desKey, publicKey, n)
	fmt.Println("The RSA encrypted DES key is:", encryptedDesKey.String())
	if err != nil {
		panic(err)
	}
	decryptedDesKey := RsaDecrypt(encryptedDesKey, privateKey, n)
	fmt.Println("The RSA decrypted DES key is:", decryptedDesKey.String())
	cipherEncrypterBlockMode := cipher.NewCBCEncrypter(cipherBlock, make([]byte, 8))
	cipherDecrypterBlockMode := cipher.NewCBCDecrypter(cipherBlock, make([]byte, 8))
	text := ReadText()
	dstText := make([]byte, len(text))
	cipherEncrypterBlockMode.CryptBlocks(dstText, text)
	fmt.Println("The encrypted bytes: ", dstText)
	cipherDecrypterBlockMode.CryptBlocks(text, dstText)
	fmt.Println("The decrypted bytes: ", text)
}
