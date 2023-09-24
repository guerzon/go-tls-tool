package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(caCmd)
}

var caCmd = &cobra.Command{
	Use:   "ca",
	Short: "Commands to create a CA.",
	Long:  `Commands to create a CA.`,
}
