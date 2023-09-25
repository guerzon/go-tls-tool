package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gotls",
	Short: "gotls is a command line tool for TLS",
	Long: `gotls is a command line tool for TLS
		Can be used to generating CAs and signing x509 certificates,
		and for interacting with TLS endpoints.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {

	// to implement
	rootCmd.PersistentFlags().StringP("version", "V", "", "Prints the gotls tool version")

}
