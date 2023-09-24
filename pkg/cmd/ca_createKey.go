package cmd

import (
	"fmt"
	"os"

	"github.com/guerzon/go-tls-tool/pkg/ca"
	"github.com/spf13/cobra"
)

var keyName string
var keyPath string
var keySize int

func init() {
	caCmd.AddCommand(createKey)
	createKey.Flags().StringVar(&keyPath, "key-path", "", "Directory where to save the private key. Defaults to the current directory.")
	createKey.Flags().StringVar(&keyName, "key-name", "ca_private.pem", "Filename of the private key.")
	createKey.Flags().IntVar(&keySize, "key-size", 4096, "Length of the private key.")
}

var createKey = &cobra.Command{
	Use:   "create-key",
	Short: "Create a private key for the CA.",
	Long:  `Create a private key for the CA.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("[+] Creating a private key ...")
		err := ca.CreatePrivateKey(keyName, keyPath, keySize)
		if err != nil {
			fmt.Printf("[-] Cannot create a private key. Error:\n%s\n", err)
			os.Exit(1)
		}
		fmt.Printf("[+] Private key created %s with length %d.\n", keyName, keySize)
	},
}
