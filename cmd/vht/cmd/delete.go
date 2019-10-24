package cmd

import (
	"fmt"
	"github.com/ilijamt/vht/internal/vault"
	"github.com/spf13/cobra"
	"regexp"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a path recursively",
	Long:  `Deletes a whole path recursively`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		client, err := vault.Client()
		if err != nil {
			return err
		}
		var paths []string
		rootPath, _ := cmd.Flags().GetString("root-path")
		filter, _ := cmd.Flags().GetString("path-filter")
		force, _ := cmd.Flags().GetBool("force")
		serial, _ := cmd.Flags().GetBool("serial")
		concurrent, err := cmd.Flags().GetInt8("concurrent")
		if err != nil {
			return err
		}
		rFilter, err := regexp.Compile(filter)
		if err != nil {
			return err
		}

		if paths, err = getTree(serial, rootPath, client, concurrent); err != nil {
			return err
		}

		if len(paths) == 0 {
			return nil
		}

		filteredPaths := vault.FilterDataPaths(paths, rFilter)
		for _, path := range filteredPaths {
			fmt.Println(path)
		}
		var confirmed bool
		if force {
			fmt.Println("Force flag detected skipping delete confirmation.")
			confirmed = true
		} else {
			fmt.Println("WARNING: This action is irreversible, please confirm?")
			confirmed = askForConfirmation()
		}
		if confirmed {
			_, err = vault.DeletePaths(filteredPaths, concurrent, client, cmd.OutOrStderr())
		}
		return err
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringP("root-path", "r", "", "The root path to look into")
	deleteCmd.Flags().StringP("path-filter", "k", ".*", "Regex to apply to the path")
	deleteCmd.Flags().BoolP("force", "f", false, "Skip confirmation to remove the path")
	deleteCmd.Flags().Int8P("concurrent", "n", 10, "How many keys to process concurrently")
	deleteCmd.Flags().BoolP("serial", "s", false, "Do not use concurrency to build the path tree")
	_ = cobra.MarkFlagRequired(deleteCmd.Flags(), "root-path")
}
