package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify connection to Vault",
	Long:  `Tries to verify that we can actually open a connection to vault`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		client, err := getVaultClient()
		if err != nil {
			return err
		}

		response, err := client.Sys().Health()
		if err != nil {
			return err
		}
		fmt.Printf("Cluster name: %s\n", response.ClusterName)
		fmt.Printf("Version: %s\n", response.Version)
		fmt.Printf("Sealed: %t\n", response.Sealed)
		return err
	},
}

func init() {
	rootCmd.AddCommand(verifyCmd)
}
