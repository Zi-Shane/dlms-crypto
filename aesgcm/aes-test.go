package aesgcm

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

/**
 * Example Data, belows are Hex String
 * plaintext:  01011000112233445566778899AABBCCDDEEFF0000065F1F0400007E1F04B0
 * key:        000102030405060708090A0B0C0D0E0F
 * nonce(IV):  4D4D4D0000BC614E01234567
 * aad:        30D0D1D2D3D4D5D6D7D8D9DADBDCDDDEDF
 * ciphertext: 801302ff8a7874133d414ced25b42534d28db0047720606b175bd52211be68
 * tag(12):    41db204d39ee6fdb8e356855
 * tag(16):    41db204d39ee6fdb8e356855f6558503
 **/

func ExampleNewGCMEncrypter() {
	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	key, _ := hex.DecodeString("000102030405060708090A0B0C0D0E0F")
	plaintext, _ := hex.DecodeString("01011000112233445566778899AABBCCDDEEFF0000065F1F0400007E1F04B0")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	// nonce(a.k.a IV) = SystemTitle || IC
	nonce, _ := hex.DecodeString("4D4D4D0000BC614E01234567")
	// aad = SC-AE || IC
	aad, _ := hex.DecodeString("30D0D1D2D3D4D5D6D7D8D9DADBDCDDDEDF")

	aesgcm, err := cipher.NewGCMWithTagSize(block, 12)
	if err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, aad)
	fmt.Printf("%x\n", ciphertext)
}

func ExampleNewGCMDecrypter() {
	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	key, _ := hex.DecodeString("000102030405060708090A0B0C0D0E0F")
	ciphertext, _ := hex.DecodeString("801302ff8a7874133d414ced25b42534d28db0047720606b175bd52211be6841db204d39ee6fdb8e356855")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	// nonce(a.k.a IV) = SystemTitle || IC
	nonce, _ := hex.DecodeString("4D4D4D0000BC614E01234567")
	// aad = SC-AE || IC
	aad, _ := hex.DecodeString("30D0D1D2D3D4D5D6D7D8D9DADBDCDDDEDF")

	aesgcm, err := cipher.NewGCMWithTagSize(block, 12)
	if err != nil {
		panic(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, aad)
	if err != nil {
		panic(err.Error())
	}

	// fmt.Printf("%s\n", string(plaintext))
	fmt.Printf("%x\n", plaintext)

}
