package ecdsa

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
)

// Key instanciation
var pubkey = &ecdsa.PublicKey{
	Curve: elliptic.P256(),
	X:     fromBase16("bc316842562ab04ca987f39fbc5368899a0f2d6059e1247b68d3dc4f26c75669"),
	Y:     fromBase16("f808c6bb115b3b43e7f3a23d3e5f4bb3628183615a5604e1e603c9563bb24942"),
}

var privkey = &ecdsa.PrivateKey{
	PublicKey: *pubkey,
	D:         fromBase16("ce69efe1e68415ad5b9f4c8b2f3025cb1332ddbd881073309a53a526fd3d7dbd"),
}

//sign will sign the provided signhash byte array using privkey
func sign(signhash []byte) (outR, outS string) {
	// If signhash is longer than the bit-length of the private key's curve
	// order, signhash will be truncated to that length. It returns the
	// signature as a pair of big integers.
	r, s, err := ecdsa.Sign(rand.Reader, privkey, signhash)
	if err != nil {
		log.Fatalln(err)
	}

	outR = fmt.Sprintf("%064x", r)
	outS = fmt.Sprintf("%064x", s)
	fmt.Printf("\tSignature: \t%s\n\t\t\t%s\n", outR, outS)
	return
}

// hash produce a signhash byte array compatible with signature and verification processes.
func hash(inMsg string) []byte {
	msg, err := hex.DecodeString(inMsg)
	if err != nil {
		panic(err)
	}
	// the hash used:
	h := sha256.New()
	h.Write(msg)
	var signhash []byte = h.Sum(nil)
	return signhash
}

// verify will verify the provided signhash upon signature r,s using pubkey
func verify(signhash []byte, inR, inS string) {
	// if we are not signing, we are verifying :
	r := fromBase16(inR)
	s := fromBase16(inS)
	verifystatus := ecdsa.Verify(pubkey, signhash, r, s)
	fmt.Println("\tVerification returned: ", verifystatus)
}

// fromBase16 is a helper method to use the prime in hex form, inspired from crypto/rsa/rsa_test.go
func fromBase16(base16 string) *big.Int {
	i, ok := new(big.Int).SetString(base16, 16)
	if !ok {
		log.Fatalln("trying to convert from base16 a bad number: ", base16)
	}
	return i
}

func ExampleECDSA() {
	// Test when everything is fine:
	signhash := hash("DEADC0DE")
	r, s := sign(signhash)
	verify(signhash, r, s)
}

func ExampleECDSASign(text string) {
	signhash := hash(text)
	fmt.Printf("\tSigning hash: %x \n", signhash)
	sign(signhash)
}

func ExampleECDSAVerify(text, signature string) {
	// Test when everything is fine:
	signhash := hash(text)
	fmt.Printf("\tSigning hash: %x \n", signhash)

	r := signature[:len(signature)/2]
	s := signature[len(signature)/2:]

	verify(signhash, r, s)
}
