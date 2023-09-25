package ca

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"path/filepath"
)

// Generate a private key keyName located in keyPath with size keySize.
// The key is stored as PEM-format.
func CreatePrivateKey(keyName string, keyPath string, keySize int) error {

	// generate the private key
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		return fmt.Errorf("cannot generate a private key: %s", err)
	}

	// convert to PEM
	privateKeyPem := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	// convert to bytes
	privateKeyBytes := pem.EncodeToMemory(&privateKeyPem)
	// write to file
	if err = os.WriteFile(filepath.Join(keyPath, keyName), privateKeyBytes, 0600); err != nil {
		return fmt.Errorf("cannot write to file: %s", err)
	}

	return nil
}
