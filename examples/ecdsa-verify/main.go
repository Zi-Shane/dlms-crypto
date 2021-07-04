package main

import (
	"github.com/Zi-Shane/dlms-crypto/ecdsa"
)

func main() {
	text := "DEADC0DE"
	sig := "d30ee63e712cedb96291d93362549aa704f560967e69b12e041e9f229abb103f6d32525fc8354feba3dfe746cebcc2460261072ba7cde498532cc4afc5d6317a"
	ecdsa.ExampleECDSAVerify(text, sig)
}
