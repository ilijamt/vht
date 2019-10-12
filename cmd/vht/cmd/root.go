package cmd

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"github.com/ilijamt/vht/internal/vault"
	"github.com/spf13/cobra"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "vht",
	Short: "Vault Helper Tool",
	Long:  `A simple vault helper tool that simplifies the usage of Vault`,
}

func Execute() error {
	return rootCmd.Execute()
}

func getTree(serial bool, rootPath string, client *api.Client, concurrent int8) (paths []string, err error) {
	if serial {
		paths, err = vault.TreeSerial(rootPath, client)
	} else {
		paths, err = vault.Tree(rootPath, client, concurrent)
	}
	return paths, err
}

func askForConfirmation() bool {
	var response string
	fmt.Printf("Please type y/yes/YES or n/no/NO and then press enter: ")
	_, err := fmt.Scanln(&response)
	if err != nil {
		return false
	}
	okayResponses := []string{"y", "yes"}
	nokayResponses := []string{"n", "no"}
	if containsString(okayResponses, strings.ToLower(response)) {
		return true
	} else if containsString(nokayResponses, strings.ToLower(response)) {
		return false
	} else {
		return askForConfirmation()
	}
}

func containsString(slice []string, element string) bool {
	for _, elem := range slice {
		if elem == element {
			return true
		}
	}
	return false
}
