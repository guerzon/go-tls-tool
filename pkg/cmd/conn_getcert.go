package cmd

import "github.com/spf13/cobra"

func init() {
	connCmd.AddCommand(getcert)
	getcert.Flags().StringP("host", "H", "", "FQDN or IP address of the target endpoint")
	getcert.Flags().StringP("port", "P", "", "Port number of the target endpoint")
}

var getcert = &cobra.Command{
	Use:   "getcert",
	Short: "Inspect the SSL certificate of a TLS endpoint.",
	Long:  `Inspect the SSL certificate of a TLS endpoint.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
