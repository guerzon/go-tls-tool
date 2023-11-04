package cmd

import "github.com/spf13/cobra"

var CertForce bool

func init() {
	rootCommand.AddCommand(certCommand)
	certCommand.PersistentFlags().BoolVarP(&CertForce, "force", "f", false, "Replace the private key or certificate if it exists.")
}

var certCommand = &cobra.Command{
	Use:   "cert",
	Short: "Commands to sign certificates.",
	Long:  `Commands to sign certificates.`,
}
