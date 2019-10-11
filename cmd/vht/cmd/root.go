package cmd

import (
	"fmt"
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
