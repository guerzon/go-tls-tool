# go-tls-tool

SSL/TLS tool written in Go for demo purposes. This is work in progress.

## Howto

Sub-commands:

- `conn` - for interacting with an SSL/TLS endpoint.
- `ca` - for creating your own CA.
- `cert` - for working with certificate files.

Use `make` to build the binary:

```bash
make build
./bin/gotls --help
```

### Testing a TLS endpoint

1. Get the supported SSL/TLS versions.
2. List certificate tree.
3. Get validity information.
4. List valid domains for the certificate.
5. Check supported ciphers.

### Creating your own CA

1. Create the YAML files containing your configuration. For reference, see:

    - `config-ca.yml` - contains CA info.
    - `config-certs.yaml` - contains server certificate info.

2. Generate a private key called `private.pem` in the current directory:

    `./bin/gotls ca create-key --key-name private.pem`

3. Create the CA certificate `private.crt`:

    `./bin/gotls ca create-cert --cert-name private.crt  --ca-key private.pem`

### Working with certificates

For signing certificates using a CA from a previous section, or for inspecting a PEM-encoded certificate file.

1. Check if a PEM-encoded certificate and a private key match.
2. Print the contents of a certificate.
