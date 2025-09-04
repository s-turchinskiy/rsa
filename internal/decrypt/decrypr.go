package decrypt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"rsa/internal/common"
)

func ReadPrivateKey(privateKeyFile string) (*rsa.PrivateKey, error) {

	priv, err := os.ReadFile(privateKeyFile)
	if err != nil {
		err = fmt.Errorf("path: %s, error: %w", privateKeyFile, err)
		return nil, common.WrapError(err)
	}

	privPem, _ := pem.Decode(priv)
	var privPemBytes []byte
	if privPem.Type != "RSA PRIVATE KEY" {
		err = fmt.Errorf("RSA private key is of the wrong type :%s", privPem.Type)
		return nil, common.WrapError(err)
	}
	privPemBytes = privPem.Bytes

	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS1PrivateKey(privPemBytes); err != nil {
		if parsedKey, err = x509.ParsePKCS8PrivateKey(privPemBytes); err != nil { // note this returns type `interface{}`
			err = fmt.Errorf("unable to parse RSA private key %w", err)
			return nil, common.WrapError(err)
		}
	}

	var privateKey *rsa.PrivateKey
	var ok bool
	privateKey, ok = parsedKey.(*rsa.PrivateKey)
	if !ok {
		err = fmt.Errorf("unable to parse RSA private key %w", err)
		return nil, common.WrapError(err)
	}

	return privateKey, nil

}
