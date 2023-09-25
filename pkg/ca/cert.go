package ca

import (
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/guerzon/go-tls-tool/pkg/cert"
	"gopkg.in/yaml.v3"
)

// Generate a CA certificate certName using the CA configuration caConfig
// and the private key keyPEM.
func CreateCACert(certName string, certPath string, keyPEM string, caConfig string) error {

	// load the private key
	keyPEMFile, err := os.ReadFile(keyPEM)
	if err != nil {
		return fmt.Errorf("cannot open private key %s: %s", keyPEM, err)
	}
	privateKey, _ := pem.Decode(keyPEMFile)
	if privateKey == nil {
		return fmt.Errorf("no PEM data found in file %s", keyPEM)
	}

	// load the configuration
	cfgFile, err := os.ReadFile(caConfig)
	if err != nil {
		return fmt.Errorf("cannot open config file %s: %s", caConfig, err)
	}
	caCertCfg := cert.Cert{}
	if err = yaml.Unmarshal(cfgFile, &caCertCfg); err != nil {
		return fmt.Errorf("cannot process config file %s: %s", caConfig, err)
	}

	// TODO: validate required cert fields
	if caCertCfg.ValidForYears == 0 {
		fmt.Println("[-] CA certificate validity (validForYears) cannot be empty or 0.")
		os.Exit(3)
	}

	// create an x509 certificate object
	caCert := &x509.Certificate{
		SerialNumber: caCertCfg.Serial,
		Subject: pkix.Name{
			Country:            []string{caCertCfg.Subject.Country},
			Organization:       []string{caCertCfg.Subject.Organization},
			OrganizationalUnit: []string{caCertCfg.Subject.OrganizationalUnit},
			Locality:           []string{caCertCfg.Subject.Locality},
			Province:           []string{caCertCfg.Subject.Province},
			StreetAddress:      []string{caCertCfg.Subject.StreetAddress},
			PostalCode:         []string{caCertCfg.Subject.PostalCode},
			CommonName:         caCertCfg.Subject.CommonName,
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(caCertCfg.ValidForYears, 0, 0),
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

	// generate the CA certificate
	caCertificate, err := x509.CreateCertificate(rand.Reader, caCert, caCert, &privateKeyRSA.PublicKey, privateKeyRSA)
	if err != nil {
		return fmt.Errorf("cannot generate a certificate: %s", err)
	}

	// finally, write to file
	err = os.WriteFile(filepath.Join(certPath, certName), caCertificate, 0644)
	if err != nil {
		return fmt.Errorf("cannot write certificate to file %s: %s", filepath.Join(certPath, certName), err)
	}

	return nil
}
