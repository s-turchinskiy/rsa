package decrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"rsa/internal/common"
)

type RSADecrypt struct {
	message    string
	privateKey *rsa.PrivateKey
}

func NewOnce(message string, privateKey *rsa.PrivateKey) *RSADecrypt {
	return &RSADecrypt{
		message:    message,
		privateKey: privateKey,
	}
}

func (d *RSADecrypt) Decrypt() (string, error) {

	ct, _ := base64.StdEncoding.DecodeString(d.message)
	label := []byte("OAEP Encrypted")
	rng := rand.Reader
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rng, d.privateKey, ct, label)
	if err != nil {
		return "", common.WrapError(fmt.Errorf("%w, message: \"%s\"", err, d.message))
	}
	return string(plaintext), nil

}
