package certs

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"path/filepath"
)

// Generate a private key in PEM-format.
func CreatePrivateKey(path string, keyName string, keySize int) error {

	// generate the private key
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		return fmt.Errorf("cannot generate private key: %s", err)
	}

	// convert to PEM
	privateKeyPem := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	// convert to bytes
	privateKeyBytes := pem.EncodeToMemory(&privateKeyPem)

	// write to file
	if err = os.WriteFile(filepath.Join(path, keyName), privateKeyBytes, 0600); err != nil {
		return fmt.Errorf("cannot write key to file: %s", err)
	}

	return nil
}
