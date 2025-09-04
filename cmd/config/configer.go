package configer

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"runtime"
)

const (
	Unknown = iota
	Encrypt
	Decrypt
	DecryptBatch
)

type Config struct {
	Operation    int
	Encrypt      *RSAEncrypt
	Decrypt      *RSADecrypt
	DecryptBatch *RSADecryptBatch
}

type RSADecrypt struct {
	Message        string
	PrivateKeyFile string
}

type RSADecryptBatch struct {
	DataFile       string
	PrivateKeyFile string
	NumThreads     int
}

type RSAEncrypt struct {
	Message       string
	PublicKeyFile string
}

func GetConfig() (*Config, error) {

	config := &Config{
		Encrypt:      &RSAEncrypt{},
		Decrypt:      &RSADecrypt{},
		DecryptBatch: &RSADecryptBatch{},
	}

	err := parseFlags(config)
	if err != nil {
		return nil, err
	}

	err = parseEnv(config)
	if err != nil {
		return nil, err
	}

	switch config.Operation {
	case Encrypt:

		err = config.Encrypt.validation()
		if err != nil {
			return nil, err
		}

	case Decrypt:

		err = config.Decrypt.validation()
		if err != nil {
			return nil, err
		}

	case DecryptBatch:

		config.DecryptBatch.NumThreads = runtime.NumCPU()

		err = config.DecryptBatch.validation()
		if err != nil {
			return nil, err
		}
	}

	return config, nil

}

func parseFlags(config *Config) error {

	encryptFags := flag.NewFlagSet("encrypt", flag.ContinueOnError)
	decryptFlags := flag.NewFlagSet("decrypt", flag.ContinueOnError)
	decryptBatchFlags := flag.NewFlagSet("decryptbatch", flag.ContinueOnError)

	encryptFags.StringVar(&config.Encrypt.Message, "msg", "", "message to encode")
	encryptFags.StringVar(&config.Encrypt.PublicKeyFile, "pub", "", "path to file with public key")

	decryptFlags.StringVar(&config.Decrypt.Message, "msg", "", "message to decode")
	decryptFlags.StringVar(&config.Decrypt.PrivateKeyFile, "pr", "", "path to file with private key")

	decryptBatchFlags.StringVar(&config.DecryptBatch.DataFile, "data", "", "path to file with data")
	decryptBatchFlags.StringVar(&config.DecryptBatch.PrivateKeyFile, "pr", "", "path to file with private key")

	if len(os.Args) < 2 {
		return nil
	}

	switch os.Args[1] {
	case "encrypt":

		err := encryptFags.Parse(os.Args[2:])
		if err != nil {
			return err
		}

	case "decrypt":

		err := decryptFlags.Parse(os.Args[2:])
		if err != nil {
			return err
		}

	case "decryptbatch":

		err := decryptBatchFlags.Parse(os.Args[2:])
		if err != nil {
			return err
		}

	default:
		return fmt.Errorf("unknown command %s, only encrypt, decrypt, decryptbatch allowed", os.Args[1])
	}

	if encryptFags.Parsed() {
		config.Operation = Encrypt
	}

	if decryptFlags.Parsed() {
		config.Operation = Decrypt
	}

	if decryptBatchFlags.Parsed() {
		config.Operation = DecryptBatch
	}

	return nil
}

func parseEnv(config *Config) error {

	operation := os.Getenv("OPERATION")

	switch operation {
	case "":
		return nil
	case "encrypt":

		config.Operation = Encrypt

		if value := os.Getenv("PUBLIC_KEY_FILE"); value != "" {
			config.Encrypt.PublicKeyFile = value
		}

		if value := os.Getenv("ENCRYPT_MESSAGE"); value != "" {
			config.Encrypt.Message = value
		}

	case "decrypt":

		config.Operation = Decrypt

		if value := os.Getenv("PRIVATE_KEY_FILE"); value != "" {
			config.Decrypt.PrivateKeyFile = value
		}

		if value := os.Getenv("DECRYPT_MESSAGE"); value != "" {
			config.Decrypt.Message = value
		}

	case "decryptbatch":

		config.Operation = DecryptBatch

		if value := os.Getenv("PRIVATE_KEY_FILE"); value != "" {
			config.DecryptBatch.PrivateKeyFile = value
		}

		if value := os.Getenv("DECRYPT_DATA_FILE"); value != "" {
			config.DecryptBatch.DataFile = value
		}

	default:
		return fmt.Errorf("unclown operatin %s, only encrypt and decrypt allowed", operation)
	}

	return nil

}

func (e RSAEncrypt) validation() error {

	var errs []error

	if e.Message == "" {
		err := fmt.Errorf("message is empty")
		errs = append(errs, err)
	}

	if e.PublicKeyFile == "" {
		err := fmt.Errorf("publicKeyFile is empty")
		errs = append(errs, err)
	}

	return errors.Join(errs...)

}

func (d RSADecrypt) validation() error {

	var errs []error

	if d.Message == "" {
		err := fmt.Errorf("message is empty")
		errs = append(errs, err)
	}

	if d.PrivateKeyFile == "" {
		err := fmt.Errorf("privateKeyFile is empty")
		errs = append(errs, err)
	}

	return errors.Join(errs...)

}

func (d RSADecryptBatch) validation() error {

	var errs []error

	if d.DataFile == "" {
		err := fmt.Errorf("dataFile is empty")
		errs = append(errs, err)
	}

	if d.PrivateKeyFile == "" {
		err := fmt.Errorf("privateKeyFile is empty")
		errs = append(errs, err)
	}

	return errors.Join(errs...)

}
