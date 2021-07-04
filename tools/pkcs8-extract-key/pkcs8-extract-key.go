package main

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

var pemPath = flag.String("path", "", "path of PrivateKey.pem file")

func GetPrivatekey(pemPath string) {
	// pemString := `-----BEGIN PRIVATE KEY-----
	// MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgzmnv4eaEFa1bn0yLLzAlyxMy3b2IEHMwmlOlJv09fb2hRANCAAS8MWhCViqwTKmH85+8U2iJmg8tYFnhJHto09xPJsdWafgIxrsRWztD5/OiPT5fS7NigYNhWlYE4eYDyVY7sklC
	// 	-----END PRIVATE KEY-----`

	pemString, err := ioutil.ReadFile(pemPath)
	if err != nil {
		panic(err.Error())
	}
	block, _ := pem.Decode([]byte(pemString))
	parseResult, _ := x509.ParsePKCS8PrivateKey(block.Bytes)
	key := parseResult.(*ecdsa.PrivateKey)
	fmt.Println(key.D.Text(16))
}

func main() {
	flag.Parse()

	if len(*pemPath) == 0 {
		log.Fatalf("Missing required --path parameter")
	}

	GetPrivatekey(*pemPath)
}
