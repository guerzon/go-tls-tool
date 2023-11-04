# gochk

`gochk` is a command-line tool for various SSL/TLS tasks.

## Usage

Download and run the pre-build binaries:

```bash
# todo
```

If you have `go` installed, you can build your own binary:

```bash
go build -o gochk main.go
./gochk --help

# using make
make build
./bin/gochk --help
```

Sub-commands (see [features](#features) for available commands):

- `ca` - for creating your own CA.
- `conn` - for interacting with an SSL/TLS endpoint.
- `cert` - for working with certificate files.

## Features

### Creating your own CA

1. Create the YAML files containing your configuration. For reference, see the example [ca configuration](./ca.yml).

2. Create the private key and the CA certificate:

    ```bash
    # use all defaults:
    ./gochk ca

    # specify filenames:
    ./gochk ca --key private.pem --cert cert.pem
    ```

## TODO

The following features are not yet implemented.

### Testing a TLS endpoint

1. Get the supported SSL/TLS versions.
2. List certificate tree.
3. Get validity information.
4. List valid domains for the certificate.
5. Check supported ciphers.

### Working with certificates

For signing certificates using a CA from a previous section, or for inspecting a PEM-encoded certificate file.

1. Check if a PEM-encoded certificate and a private key match.
2. Print the contents of a certificate.
