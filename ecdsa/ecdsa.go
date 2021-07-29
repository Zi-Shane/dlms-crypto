package ecdsa

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"

	helper "github.com/Zi-Shane/dlms-crypto/tools"
)

/**
 * PublicKey:  "bc316842562ab04ca987f39fbc5368899a0f2d6059e1247b68d3dc4f26c75669"
 *             "f808c6bb115b3b43e7f3a23d3e5f4bb3628183615a5604e1e603c9563bb24942"
 * PrivateKey: "ce69efe1e68415ad5b9f4c8b2f3025cb1332ddbd881073309a53a526fd3d7dbd"
**/

// Key instanciation
var pubkey *ecdsa.PublicKey
var privkey *ecdsa.PrivateKey

func Setkeypair(pubkeyXY, privkeyD string) {
	pubkeyX := pubkeyXY[:len(pubkeyXY)/2]
	pubkeyY := pubkeyXY[len(pubkeyXY)/2:]

	pubkey = &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     helper.HextoBigInt(pubkeyX),
		Y:     helper.HextoBigInt(pubkeyY),
	}

	privkey = &ecdsa.PrivateKey{
		PublicKey: *pubkey,
		D:         helper.HextoBigInt(privkeyD),
	}
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

//sign will sign the provided signhash byte array using privkey
func Suite1Sign(text string) (outR, outS string) {
	// If signhash is longer than the bit-length of the private key's curve
	// order, signhash will be truncated to that length. It returns the
	// signature as a pair of big integers.
	signhash := hash(text)
	// fmt.Printf("\tSigning hash: %x \n", signhash)

	r, s, err := ecdsa.Sign(rand.Reader, privkey, signhash)
	if err != nil {
		log.Fatalln(err)
	}

	outR = fmt.Sprintf("%064x", r)
	outS = fmt.Sprintf("%064x", s)
	// fmt.Printf("\tSignature: \t%s\n\t\t\t%s\n", outR, outS)
	return
}

// verify will verify the provided signhash upon signature r,s using pubkey
func Suite1Verify(text, signature string) bool {
	signhash := hash(text)
	// fmt.Printf("\tSigning hash: %x \n", signhash)

	r := helper.HextoBigInt(signature[:len(signature)/2])
	s := helper.HextoBigInt(signature[len(signature)/2:])
	verifystatus := ecdsa.Verify(pubkey, signhash, r, s)
	// fmt.Println("\tVerification returned: ", verifystatus)
	return verifystatus
}
