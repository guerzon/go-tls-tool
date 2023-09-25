package cmd

import (
	"log/slog"
	"os"
	"path/filepath"
	"strconv"

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
				slog.Error("ca: private key " + fStat.Name() + " already exists. If you want to replace this key, pass the [--force | -f] command-line argument.")
				os.Exit(1)
			}
			slog.Warn("ca: private key " + fStat.Name() + " exists and will be replaced.")
			slog.Warn("ca: if a certificate was previously created, you need to recreate that certificate as well.")
		}

		slog.Info("ca: creating a private key ...")
		err := ca.CreatePrivateKey(keyName, keyPath, keySize)
		if err != nil {
			slog.Error("ca: cannot create a private key. Error: ", "msg", err.Error())
			os.Exit(2)
		}
		slog.Info("ca: private key " + keyName + " created with length " + strconv.Itoa(keySize) + ".")
	},
}
