package cmd

import (
	"github.com/ilijamt/vht/internal/vault"
	"github.com/spf13/cobra"
)

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify connection to Vault",
	Long:  `Tries to verify that we can actually open a connection to vault`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		client, err := vault.Client()
		if err != nil {
			return err
		}
		return vault.Health(cmd.OutOrStdout(), client)
	},
}

func init() {
	rootCmd.AddCommand(verifyCmd)
}
