package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/guerzon/go-tls-tool/pkg/ca"
	"github.com/spf13/cobra"
)

var keyPath string
var keyName string
var keySize int

func init() {
	caCmd.AddCommand(createKey)
	createKey.Flags().StringVar(&keyPath, "key-path", "", "Directory where to save the private key. Defaults to the current directory.")
	createKey.Flags().StringVar(&keyName, "key-name", "ca_key.pem", "Filename of the private key.")
	createKey.Flags().IntVar(&keySize, "key-size", 4096, "Length of the private key.")
}

var createKey = &cobra.Command{
	Use:   "create-key",
	Short: "Create a private key for the CA.",
	Long:  `Create a private key for the CA.`,
	Run: func(cmd *cobra.Command, args []string) {

		fStat, statErr := os.Stat(filepath.Join(keyPath, keyName))
		if statErr == nil {
			if !force {
				fmt.Printf("[-] Private key \"%s\" already exists. If you want to replace this key, pass the [--force | -f] command-line argument.\n", fStat.Name())
				os.Exit(1)
			}
			fmt.Printf("[+] Private key \"%s\" exists and will be replaced.\n", fStat.Name())
			fmt.Printf("If a certificate was previously created, you need to recreate that certificate as well.\n\n")
		}

		fmt.Println("[+] Creating a private key ...")
		err := ca.CreatePrivateKey(keyName, keyPath, keySize)
		if err != nil {
			fmt.Printf("[-] Cannot create a private key. Error:\n%s\n", err)
			os.Exit(2)
		}
		fmt.Printf("[+] Private key %s created with length %d.\n", keyName, keySize)
	},
}
