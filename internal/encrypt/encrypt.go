package encrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
	"rsa/internal/common"
)

type RSAEncrypt struct {
	message           string
	publicKeyLocation string
	publicKey         *rsa.PublicKey
}

func New(message, publicKeyLocation string) *RSAEncrypt {
	return &RSAEncrypt{
		message:           message,
		publicKeyLocation: publicKeyLocation,
	}
}

func (e RSAEncrypt) Encrypt() (string, error) {

	var err error

	if e.publicKey, err = e.readPublicKey(); err != nil {
		return "", err
	}

	return e.rsaOaepEncrypt()
}

func (e RSAEncrypt) rsaOaepEncrypt() (string, error) {
	label := []byte("OAEP Encrypted")
	rng := rand.Reader
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rng, e.publicKey, []byte(e.message), label)
	if err != nil {
		return "", common.WrapError(err)
	}
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func (e RSAEncrypt) readPublicKey() (*rsa.PublicKey, error) {

	pub, err := os.ReadFile(e.publicKeyLocation)
	if err != nil {
		err = fmt.Errorf("path: %s, error: %w", e.publicKeyLocation, err)
		return nil, common.WrapError(err)
	}

	pubPem, _ := pem.Decode(pub)
	if pubPem == nil {
		return nil, common.WrapError(err)
	}

	if pubPem.Type != "RSA PUBLIC KEY" && pubPem.Type != "PUBLIC KEY" {
		err := fmt.Errorf("RSA public key is of the wrong type, Pem Type :%s", pubPem.Type)
		return nil, common.WrapError(err)
	}

	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKIXPublicKey(pubPem.Bytes); err != nil {
		err = fmt.Errorf("unable to parse RSA public key: %w", err)
		return nil, common.WrapError(err)
	}

	var pubKey *rsa.PublicKey
	var ok bool
	if pubKey, ok = parsedKey.(*rsa.PublicKey); !ok {
		err = fmt.Errorf("unable to parse RSA public key: %w", err)
		return nil, common.WrapError(err)
	}

	return pubKey, nil

}
