package cmd

import "github.com/spf13/cobra"

func init() {
	certCmd.AddCommand(csr)
}

var csr = &cobra.Command{
	Use:   "csr",
	Short: "Create a certificate signing request.",
	Long:  `Create a certificate signing request.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
