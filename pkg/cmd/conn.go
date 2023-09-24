package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(connCmd)
}

var connCmd = &cobra.Command{
	Use:   "conn",
	Short: "Commands used to interact with TLS endpoints.",
	Long:  `Commands used to interact with TLS endpoints.`,
}
