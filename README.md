DLMS Crypto Tools
==========

AES-128-GCM
----------

- AES-128-GCM Encryption

```
go run examples/aes-encrypt/main.go
```

- AES-128-GCM Encryption

```
go run examples/aes-decrypt/main.go
```

----------

ECDSA
----------

#### change key pair in `ecdsa/ecdsa.go`

- ECDSA Sign

```
go run examples/ecdsa-sign/main.go
```

- ECDSA Verify

```
go run examples/ecdsa-verify/main.go
```

ECDH
----------

----------

tools for parse .pem
----------
- get privat key from PKCS#8 format  
```
go run tools/pkcs8-extract-key/pkcs8-extract-key.go --path ECDSA_Keypairs/server/Copy_to_Keys/D4142434445464748.pem
```
- get public key from x509 certification  
```
go run tools/x509-extract-key/x509-extract-key.go --path ECDSA_Keypairs/server/Copy_to_Certificates/D4142434445464748.pem
```

