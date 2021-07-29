package tools

import (
	"log"
	"math/big"
)

func HextoBigInt(base16 string) *big.Int {
	i, ok := new(big.Int).SetString(base16, 16)
	if !ok {
		log.Fatalln("trying to convert from base16 a bad number: ", base16)
	}
	return i
}
