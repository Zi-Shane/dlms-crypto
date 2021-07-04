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

var pemPath = flag.String("path", "", "path of cert.pem file")

func GetPublickey(pemPath string) {
	// 	pemString := `-----BEGIN CERTIFICATE-----
	// MIIBfTCCARqgAwIBAgIDJkYPMAoGCCqGSM49BAMCME4xFDASBgNVBAMMC1Jvb3QgQ0EgMjU2MRcwFQYDVQQKDA5HdXJ1eCBTZWN1cml0eTEQMA4GA1UEBwwHVGFtcGVyZTELMAkGA1UEBhMCRkkwHhcNMjEwNDEyMDU0MzI2WhcNMjIwNDEyMDU0MzI1WjAbMRkwFwYDVQQDDBA0MTQyNDM0NDQ1NDY0NzQ4MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEvDFoQlYqsEyph/OfvFNoiZoPLWBZ4SR7aNPcTybHVmn4CMa7EVs7Q+fzoj0+X0uzYoGDYVpWBOHmA8lWO7JJQqMaMBgwCQYDVR0TBAIwADALBgNVHQ8EBAMCB4AwCgYIKoZIzj0EAwIDUQAwTgYIKoZIzj0EAwICIAmYxkxBHy6l6QvMotr0G4oYa0HzSVbC+vgWrPBv71+fAiALSTGGL8CYZ6OwdoTSxTPJYnvwzo1gMV4wajT2bwxc0g==
	// -----END CERTIFICATE-----`

	pemString, err := ioutil.ReadFile(pemPath)
	if err != nil {
		panic(err.Error())
	}
	block, _ := pem.Decode([]byte(pemString))
	parseResult, _ := x509.ParseCertificate(block.Bytes)
	key := parseResult.PublicKey.(*ecdsa.PublicKey)
	fmt.Println(key.X.Text(16))
	fmt.Println(key.Y.Text(16))
}

func main() {
	flag.Parse()

	if len(*pemPath) == 0 {
		log.Fatalf("Missing required --path parameter")
	}

	GetPublickey(*pemPath)
}
