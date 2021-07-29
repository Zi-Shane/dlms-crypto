package aesgcm

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

func ExampleNewGCMEncrypter(cipherkey, text, IV, AAD string) {
	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	key, _ := hex.DecodeString(cipherkey)
	plaintext, _ := hex.DecodeString(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	// nonce(a.k.a IV) = SystemTitle || IC
	nonce, _ := hex.DecodeString(IV)
	// aad = SC-AE || IC
	aad, _ := hex.DecodeString(AAD)

	aesgcm, err := cipher.NewGCMWithTagSize(block, 12)
	if err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, aad)
	fmt.Printf("%x\n", ciphertext)
}

func ExampleNewGCMDecrypter(cipherkey, encyptedData, IV, AAD string) {
	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	key, _ := hex.DecodeString(cipherkey)
	ciphertext, _ := hex.DecodeString(encyptedData)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	// nonce(a.k.a IV) = SystemTitle || IC
	nonce, _ := hex.DecodeString(IV)
	// aad = SC-AE || IC
	aad, _ := hex.DecodeString(AAD)

	aesgcm, err := cipher.NewGCMWithTagSize(block, 12)
	if err != nil {
		panic(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, aad)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%x\n", plaintext)

}
