package main

import (
	"github.com/Zi-Shane/dlms-crypto/ecdsa"
)

func main() {
	text := "DEADC0DE"

	ecdsa.ExampleECDSASign(text)
}
