package certs

import (
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

// Generate a certificate using a configuration file and a private key in PEM format.
func CreateCertificate(path string, keyPEM string, caConfig string, certName string) error {

	keyFile := filepath.Join(path, keyPEM)
	certFile := filepath.Join(path, certName)

	// load the private key
	keyPEMFile, err := os.ReadFile(keyFile)
	if err != nil {
		return fmt.Errorf("cannot open key %s: %s", keyFile, err)
	}
	privateKey, _ := pem.Decode(keyPEMFile)
	if privateKey == nil {
		return fmt.Errorf("no PEM data found in file %s", keyFile)
	}

	// load the configuration
	configFile, err := os.ReadFile(caConfig)
	if err != nil {
		return fmt.Errorf("cannot open config file %s: %s", caConfig, err)
	}
	certConfig := Certificate{}
	if err = yaml.Unmarshal(configFile, &certConfig); err != nil {
		return fmt.Errorf("cannot process config file %s: %s", caConfig, err)
	}

	// TODO: validate required cert fields
	if certConfig.ValidForYears == 0 {
		fmt.Println("[-] CA certificate validity (validForYears) cannot be empty or 0.")
		os.Exit(3)
	}

	// create an x509 certificate object
	// TODO: if the field does not exist, it should not be added to the cert.
	certificate := &x509.Certificate{
		SerialNumber: certConfig.Serial,
		Subject: pkix.Name{
			Country:            []string{certConfig.Subject.Country},
			Organization:       []string{certConfig.Subject.Organization},
			OrganizationalUnit: []string{certConfig.Subject.OrganizationalUnit},
			Locality:           []string{certConfig.Subject.Locality},
			Province:           []string{certConfig.Subject.Province},
			StreetAddress:      []string{certConfig.Subject.StreetAddress},
			PostalCode:         []string{certConfig.Subject.PostalCode},
			CommonName:         certConfig.Subject.CommonName,
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(certConfig.ValidForYears, 0, 0),
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	// convert key to RSA
	privateKeyRSA, err := x509.ParsePKCS1PrivateKey(privateKey.Bytes)
	if err != nil {
		return fmt.Errorf("cannot convert private key to RSA %s: %s", keyPEM, err)
	}

	// generate the certificate
	certBytes, err := x509.CreateCertificate(rand.Reader, certificate, certificate, &privateKeyRSA.PublicKey, privateKeyRSA)
	if err != nil {
		return fmt.Errorf("cannot generate certificate: %s", err)
	}

	// write to file
	err = os.WriteFile(certFile, certBytes, 0644)
	if err != nil {
		return fmt.Errorf("cannot write certificate to file %s: %s", certFile, err)
	}

	return nil
}
