package cmd

import (
	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generates bash completion scripts",
	Long: `To load completion run

. <(vht completion)

To configure your bash shell to load completions for each session add to your bashrc

# ~/.bashrc or ~/.profile
. <(vht completion)
`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = rootCmd.GenBashCompletion(cmd.OutOrStdout());
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
