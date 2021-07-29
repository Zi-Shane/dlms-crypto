package main

import (
	"fmt"

	dlms_ecdsa "github.com/Zi-Shane/dlms-crypto/ecdsa"
)

func main() {
	dlms_ecdsa.Setkeypair("bc316842562ab04ca987f39fbc5368899a0f2d6059e1247b68d3dc4f26c75669f808c6bb115b3b43e7f3a23d3e5f4bb3628183615a5604e1e603c9563bb24942", "ce69efe1e68415ad5b9f4c8b2f3025cb1332ddbd881073309a53a526fd3d7dbd")
	text := "DEADC0DE"

	outR, outS := dlms_ecdsa.Suite1Sign(text)
	fmt.Printf("\tSignature: \t%s\n\t\t\t%s\n", outR, outS)
}
