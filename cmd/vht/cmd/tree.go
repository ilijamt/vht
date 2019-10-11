package cmd

import (
	"fmt"
	"github.com/ilijamt/vht/internal/vault"
	"github.com/spf13/cobra"
	"regexp"
)

var treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "Print out a list of all the secrets in a path",
	Long:  `Print out a list of all the secrets in a path recursively`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		client, err := vault.Client()
		if err != nil {
			return err
		}
		var paths []string
		rootPath, _ := cmd.Flags().GetString("root-path")
		filter, _ := cmd.Flags().GetString("path-filter")
		rFilter, err := regexp.Compile(filter)
		if err != nil {
			return err
		}
		paths, err = vault.Tree(rootPath, client)
		if err != nil {
			return err
		}
		if len(paths) == 0 {
			return nil
		}
		for _, path := range vault.FilterDataPaths(paths, rFilter) {
			fmt.Println(path)
		}
		return err
	},
}

func init() {
	rootCmd.AddCommand(treeCmd)
	treeCmd.Flags().StringP("root-path", "r", "", "The root path to look into")
	treeCmd.Flags().StringP("path-filter", "k", ".*", "Regex to apply to the path")
	_ = cobra.MarkFlagRequired(treeCmd.Flags(), "root-path")
}
