package cmd

import "github.com/spf13/cobra"

var force bool

func init() {
	rootCmd.AddCommand(caCmd)
	caCmd.PersistentFlags().BoolVarP(&force, "force", "f", false, "Replace the private key or certificate if it exists.")
}

var caCmd = &cobra.Command{
	Use:   "ca",
	Short: "Commands to create a CA.",
	Long:  `Commands to create a CA.`,
}
