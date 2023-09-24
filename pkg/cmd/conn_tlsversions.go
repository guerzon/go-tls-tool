package cmd

import "github.com/spf13/cobra"

func init() {
	connCmd.AddCommand(tlsversions)
}

var tlsversions = &cobra.Command{
	Use:   "tlsversions",
	Short: "Get the supported TLS versions of a TLS endpoint.",
	Long:  `Get the supported TLS versions of a TLS endpoint.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
