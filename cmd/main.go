package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "gochk",
	Short: "gochk is a command line tool for TLS",
	Long: `gochk is a command line tool for TLS
		Can be used to generating CAs and signing x509 certificates,
		and for interacting with TLS endpoints.`,
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		slog.Error(err.Error())
		os.Exit(1)
	}
}
