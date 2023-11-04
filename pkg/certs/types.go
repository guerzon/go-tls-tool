package certs

import "math/big"

type Certificate struct {
	Serial        *big.Int    `yaml:"serial"`
	ValidForYears int         `yaml:"validForYears"`
	Subject       certSubject `yaml:"subject"`
	DNSNames      []string    `yaml:"dnsNames"`
}

type certSubject struct {
	Country            string `yaml:"country"`
	Organization       string `yaml:"organization"`
	OrganizationalUnit string `yaml:"organizationalUnit"`
	Locality           string `yaml:"locality"`
	Province           string `yaml:"province"`
	StreetAddress      string `yaml:"streetAddress"`
	PostalCode         string `yaml:"postalCode"`
	SerialNumber       string `yaml:"serialNumber"`
	CommonName         string `yaml:"commonName"`
}
