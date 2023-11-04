package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"

	certs "github.com/guerzon/gochk/pkg/certs"
	"github.com/spf13/cobra"
)

var force bool

var savePath string

var keyName string
var keySize int
var certName string
var caConfig string

func init() {
	rootCommand.AddCommand(caCommand)
	caCommand.PersistentFlags().BoolVarP(&force, "force", "f", false, "Replace the private key or certificate if any of them exists.")
	caCommand.Flags().StringVarP(&caConfig, "config", "c", "ca.yml", "Filename of the CA's configuration file (YAML).")
	caCommand.Flags().StringVarP(&savePath, "output-dir", "d", "", "Directory where to save the private key. Defaults to the current directory.")
	caCommand.Flags().StringVar(&keyName, "key", "ca_key.pem", "Filename of the private key.")
	caCommand.Flags().IntVar(&keySize, "key-size", 4096, "Length of the private key.")
	caCommand.Flags().StringVar(&certName, "cert", "ca_cert.pem", "Filename of the CA certificate.")
}

var caCommand = &cobra.Command{
	Use:   "ca",
	Short: "Commands to create a CA.",
	Long:  `Commands to create a CA.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Check if the path exists
		if savePath != "" {
			_, statErr := os.Stat(savePath)
			if statErr != nil {
				slog.Info(fmt.Sprintf("[ca]: save path %s does not exist, creating ...", savePath))
				err := os.Mkdir(savePath, 0755)
				if err != nil {
					slog.Error(fmt.Sprintf("[ca]: cannot create path %s", savePath), "errmsg", err.Error())
					os.Exit(1)
				}
				slog.Info(fmt.Sprintf("[ca]: created directory: %s.", savePath))
			}
		}

		// Check if the private key exists
		fStat, statErr := os.Stat(filepath.Join(savePath, keyName))
		if statErr == nil {
			if !force {
				slog.Error("[ca]: private key " + fStat.Name() + " already exists. If you want to replace this key, pass the [--force | -f] command-line argument.")
				os.Exit(1)
			}
			slog.Warn("[ca]: private key " + fStat.Name() + " exists and will be replaced.")
			slog.Warn("[ca]: if a certificate was previously created with the same name, it will be replaced.")
		}

		// Check if the certificate exists
		fStat, statErr = os.Stat(filepath.Join(savePath, certName))
		if statErr == nil {
			if !force {
				slog.Error("[ca]: certificate " + fStat.Name() + " already exists. If you want to replace this file, pass the [--force | -f] command-line argument.")
				os.Exit(1)
			}
			slog.Warn("[ca]: certificate " + certName + " exists and will be replaced.")
		}

		slog.Info(fmt.Sprintf("[ca]: creating ca key and certificate in %s", savePath))

		// Generate the private key
		slog.Info("[ca]: creating a private key ...")
		err := certs.CreatePrivateKey(savePath, keyName, keySize)
		if err != nil {
			slog.Error("[ca]: cannot create a private key. Error: ", "errmsg", err.Error())
			os.Exit(2)
		}
		slog.Info("[ca]: private key " + keyName + " created with length " + strconv.Itoa(keySize) + ".")

		// Generate the CA certificate
		slog.Info("[ca]: creating a CA certificate ...")
		err = certs.CreateCertificate(savePath, keyName, caConfig, certName)
		if err != nil {
			slog.Error("[ca]: cannot create CA certificate.", "errmsg", err.Error())
			os.Exit(2)
		}
		slog.Info("[ca]: certificate " + certName + " created.")

	},
}
