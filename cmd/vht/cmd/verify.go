package cmd

import (
	"github.com/ilijamt/vht/internal/vault"
	v "github.com/ilijamt/vht/pkg/vault"
	"github.com/spf13/cobra"
)

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify connection to Vault",
	Long:  `Tries to verify that we can actually open a connection to vault`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		client, err := v.Client()
		if err != nil {
			return err
		}
		return vault.Health(cmd.OutOrStdout(), client)
	},
}

func init() {
	rootCmd.AddCommand(verifyCmd)
}
