package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(certCmd)
}

var certCmd = &cobra.Command{
	Use:   "cert",
	Short: "Commands to sign certificates.",
	Long:  `Commands to sign certificates.`,
}
