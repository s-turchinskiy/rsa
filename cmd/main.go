package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"rsa/cmd/config"
	"rsa/internal/decrypt"
	"rsa/internal/encrypt"
)

func main() {

	_ = godotenv.Load("./cmd/.env")

	var err error
	config, err := configer.GetConfig()
	checkError(err)

	switch config.Operation {
	case configer.Encrypt:

		rsaEncrypt := encrypt.New(
			config.Encrypt.Message,
			config.Encrypt.PublicKeyFile)

		encryptedMessage, err := rsaEncrypt.Encrypt()
		checkError(err)
		fmt.Print(encryptedMessage)

	case configer.Decrypt:

		privateKey, err := decrypt.ReadPrivateKey(config.Decrypt.PrivateKeyFile)
		checkError(err)

		decryptedMessage, err := decrypt.NewOnce(
			config.Decrypt.Message,
			privateKey).Decrypt()

		checkError(err)
		fmt.Print(decryptedMessage)

	case configer.DecryptBatch:

		decryptedMessages, err := decrypt.NewBatch(
			config.DecryptBatch.DataFile,
			config.DecryptBatch.PrivateKeyFile,
			config.DecryptBatch.NumThreads).Decrypt()

		checkError(err)
		fmt.Print(decryptedMessages)

	default:
		err = fmt.Errorf("operation not specified")
		checkError(err)

	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
