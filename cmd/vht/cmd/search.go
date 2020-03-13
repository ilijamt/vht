package cmd

import (
	"fmt"
	"github.com/gammazero/workerpool"
	"github.com/hashicorp/vault/sdk/helper/jsonutil"
	"github.com/ilijamt/vht/internal/vault"
	"github.com/spf13/cobra"
	"regexp"
	"sort"
	"strings"
	"sync"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search in the secrets data",
	Long:  `Search through all the secrets to find out where the code is used`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		client, err := vault.Client()
		if err != nil {
			return err
		}
		var paths []string
		rootPath, _ := cmd.Flags().GetString("root-path")
		serial, _ := cmd.Flags().GetBool("serial")
		dnd, _ := cmd.Flags().GetBool("dont-display-data")
		concurrent, err := cmd.Flags().GetInt8("concurrent")
		if err != nil {
			return err
		}

		mountPoint := strings.Split(rootPath, "/")[0] + "/"
		isv2, err := vault.IsV2(mountPoint, client)
		if err != nil {
			return err
		}

		dataFilter, _ := cmd.Flags().GetString("data-filter")
		rDataFilter, err := regexp.Compile(dataFilter)
		if err != nil {
			return err
		}

		pathFilter, _ := cmd.Flags().GetString("path-filter")
		rPathFilter, err := regexp.Compile(pathFilter)
		if err != nil {
			return err
		}

		// we need to assemble the path if we are searching in a v2 kv store
		if isv2 {
			rootPath = strings.ReplaceAll(mountPoint+"/metadata/"+strings.Join(strings.Split(rootPath, "/")[1:], "/"), "//", "/")
		}
		if paths, err = getTree(serial, rootPath, client, concurrent); err != nil {
			return err
		}

		if len(paths) == 0 {
			return nil
		}

		type FilteredData struct {
			Path     string
			Data     map[string]interface{}
			Metadata map[string]interface{}
		}

		var data []FilteredData
		var muLock = &sync.Mutex{}
		var task = func(path string) func() {
			return func() {
				var pathv = path
				if isv2 {
					path = strings.ReplaceAll(path, "/metadata/", "/")
					parts := strings.Split(path, "/")
					pathv = parts[0] + "/data/" + strings.Join(parts[1:], "/")
				}
				secret, err := client.Logical().Read(pathv)
				if err != nil {
					fmt.Println(err)
					return
				}

				j, err := jsonutil.EncodeJSON(secret.Data)
				if err != nil {
					fmt.Println(err)
					return
				}

				if rDataFilter.MatchString(string(j)) {
					muLock.Lock()
					var payload = secret.Data
					var metadata map[string]interface{}
					if isv2 {
						payload = secret.Data["data"].(map[string]interface{})
						metadata = secret.Data["metadata"].(map[string]interface{})
					}
					data = append(data, FilteredData{Path: path, Data: payload, Metadata: metadata})
					muLock.Unlock()
				}

			}
		}

		wp := workerpool.New(int(concurrent))
		filteredPaths := vault.FilterDataPaths(paths, rPathFilter)
		for _, key := range filteredPaths {
			wp.Submit(task(key))
		}

		wp.StopWait()

		sort.Slice(data, func(i, j int) bool {
			return data[i].Path < data[j].Path
		})

		for _, secret := range data {
			fmt.Printf("%s\n", secret.Path)
			if !dnd {
				fmt.Println(strings.Repeat("-", len(secret.Path)))
				for key, val := range secret.Data {
					fmt.Printf("%s = %v\n", key, val)
				}
				fmt.Println()
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringP("root-path", "r", "", "The root path to look into")
	searchCmd.Flags().StringP("data-filter", "f", ".*", "Regex to apply to the data")
	searchCmd.Flags().BoolP("dont-display-data", "d", false, "Do not display the data dump")
	searchCmd.Flags().StringP("path-filter", "k", ".*", "Regex to apply to the path")
	searchCmd.Flags().Int8P("concurrent", "n", 10, "How many keys to process concurrently")
	searchCmd.Flags().BoolP("serial", "s", false, "Do not use concurrency to build the path tree")
	_ = cobra.MarkFlagRequired(searchCmd.Flags(), "root-path")
}
