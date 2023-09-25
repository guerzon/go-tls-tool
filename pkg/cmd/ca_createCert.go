package cmd

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/guerzon/go-tls-tool/pkg/ca"
	"github.com/spf13/cobra"
)

var certPath string
var certName string
var caConfig string

func init() {
	caCmd.AddCommand(createCert)
	createCert.Flags().StringVar(&certPath, "cert-path", "", "Directory where to save the CA certificate. Defaults to the current directory.")
	createCert.Flags().StringVar(&certName, "cert-name", "ca_cert.pem", "Filename of the CA certificate.")
	createCert.Flags().StringVar(&keyName, "ca-key", "ca_key.pem", "Filename of the private key.")
	createCert.Flags().StringVar(&caConfig, "ca-config", "ca_config.yml", "Filename of the CA's YAML configuration file.")
}

var createCert = &cobra.Command{
	Use:   "create-cert",
	Short: "Create a CA certificate.",
	Long:  `Create a CA certificate.`,
	Run: func(cmd *cobra.Command, args []string) {

		fStat, statErr := os.Stat(filepath.Join(certPath, certName))
		if statErr == nil {
			if !force {
				slog.Error("ca: certificate " + fStat.Name() + " already exists. If you want to replace this file, pass the [--force | -f] command-line argument.")
				os.Exit(1)
			}
			slog.Warn("ca: certificate " + certName + " exists and will be replaced.")
		}

		slog.Info("ca: creating a CA certificate ...")
		err := ca.CreateCACert(certPath, certName, keyName, caConfig)
		if err != nil {
			slog.Error("ca: cannot create CA certificate.", "msg", err.Error())
			os.Exit(2)
		}
		slog.Info("ca: certificate " + certName + " created.")
	},
}
