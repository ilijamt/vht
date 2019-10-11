package cmd

import (
	"github.com/ilijamt/vht/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows the version of the application",
	Long:  `Shows the version of the application`,
	Run: func(cmd *cobra.Command, args []string) {
		version.PrintVersion(cmd.OutOrStdout())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
