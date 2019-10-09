package cmd

import (
	"fmt"
	"github.com/ilijamt/vht"
	"github.com/spf13/cobra"
	"runtime"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows the version of the application",
	Long:  `Shows the version of the application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s\n", vht.BuildVersion);
		fmt.Printf("Git Commit Hash: %s\n", vht.BuildHash);
		fmt.Printf("Build Date: %s\n", vht.BuildDate);
		fmt.Printf("OS: %s\n", runtime.GOOS)
		fmt.Printf("Architecture: %s\n", runtime.GOARCH)

	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
