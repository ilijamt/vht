package cmd

import (
	"fmt"
	"github.com/gammazero/workerpool"
	"github.com/hashicorp/vault/sdk/helper/jsonutil"
	"github.com/ilijamt/vht/internal/vault"
	log "github.com/sirupsen/logrus"
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
		client, err := getVaultClient()
		if err != nil {
			return err
		}
		var paths []string
		rootPath, _ := cmd.Flags().GetString("root-path")
		concurrent, err := cmd.Flags().GetInt8("concurrent")
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

		log.WithFields(log.Fields{
			"root-path": rootPath,
		}).Info("Processing paths")
		paths, err = vault.Tree(rootPath, client)
		if err != nil {
			return err
		}
		if len(paths) == 0 {
			return nil
		}
		log.WithFields(log.Fields{
			"root-path": rootPath,
			"total":     len(paths),
		}).Info("Paths processing finished")

		type FilteredData struct {
			Path string
			Data map[string]interface{}
		}

		var data []FilteredData
		var muLock = &sync.Mutex{}
		var task = func(path string) func() {
			return func() {
				log.WithField("path", path).Debugln("Processing")
				defer log.WithField("path", path).Debugln("Finished")
				secret, err := client.Logical().Read(path)
				if err != nil {
					log.WithField("path", path).Error(err)
					return
				}
				j, err := jsonutil.EncodeJSON(secret.Data)
				if err != nil {
					log.WithField("path", path).Error(err)
					return
				}
				if rDataFilter.MatchString(string(j)) {
					muLock.Lock()
					data = append(data, FilteredData{Path: path, Data: secret.Data})
					muLock.Unlock()
				}
			}
		}

		wp := workerpool.New(int(concurrent))
		filteredPaths := vault.FilterDataPaths(paths, rPathFilter)
		log.WithFields(log.Fields{
			"root-path": rootPath,
			"total":     len(filteredPaths),
		}).Info("Filtered paths processed")
		for _, key := range filteredPaths {
			wp.Submit(task(key))
		}

		wp.StopWait()

		sort.Slice(data, func(i, j int) bool {
			return data[i].Path < data[j].Path
		})

		for _, secret := range data {
			fmt.Printf("%s\n", secret.Path)
			fmt.Println(strings.Repeat("-", len(secret.Path)))
			for key, val := range secret.Data {
				fmt.Printf("%s = %v\n", key, val)
			}
			fmt.Println()
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringP("root-path", "r", "", "The root path to look into")
	searchCmd.Flags().StringP("data-filter", "f", ".*", "Regex to apply to the data")
	searchCmd.Flags().StringP("path-filter", "k", ".*", "Regex to apply to the path")
	searchCmd.Flags().Int8P("concurrent", "n", 10, "How many keys to process concurrently")
	_ = cobra.MarkFlagRequired(searchCmd.Flags(), "root-path")
}
