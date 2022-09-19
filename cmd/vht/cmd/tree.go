package cmd

import (
	"fmt"
	"github.com/ilijamt/vht/internal/vault"
	v "github.com/ilijamt/vht/pkg/vault"
	"github.com/spf13/cobra"
	"regexp"
	"strings"
)

var treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "Print out a list of all the secrets in a path",
	Long:  `Print out a list of all the secrets in a path recursively`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		client, err := v.Client()
		if err != nil {
			return err
		}
		var paths []string
		rootPath, _ := cmd.Flags().GetString("root-path")
		filter, _ := cmd.Flags().GetString("path-filter")
		serial, _ := cmd.Flags().GetBool("serial")
		concurrent, err := cmd.Flags().GetInt8("concurrent")
		if err != nil {
			return err
		}
		rFilter, err := regexp.Compile(filter)
		if err != nil {
			return err
		}
		mountPoint := strings.Split(rootPath, "/")[0] + "/"
		isv2, err := vault.IsV2(mountPoint, client)
		if err != nil {
			return err
		}
		if isv2 {
			rootPath = strings.ReplaceAll(mountPoint+"/metadata/"+strings.Join(strings.Split(rootPath, "/")[1:], "/"), "//", "/")
		}
		if paths, err = getTree(serial, rootPath, client, concurrent); err != nil {
			return err
		}
		if len(paths) == 0 {
			return nil
		}
		for _, path := range vault.FilterDataPaths(paths, rFilter) {
			fmt.Println(strings.ReplaceAll(path, "/metadata/", "/"))
		}
		return err
	},
}

func init() {
	rootCmd.AddCommand(treeCmd)
	treeCmd.Flags().StringP("root-path", "r", "", "The root path to look into")
	treeCmd.Flags().StringP("path-filter", "k", ".*", "Regex to apply to the path")
	treeCmd.Flags().Int8P("concurrent", "n", 10, "How many keys to process concurrently")
	treeCmd.Flags().BoolP("serial", "s", false, "Do not use concurrency to build the path tree")
	_ = cobra.MarkFlagRequired(treeCmd.Flags(), "root-path")
}
