package cmd

import (
	"errors"
	"fmt"
	"github.com/hashicorp/vault/api"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"strings"
	"time"
)

var rootCmd = &cobra.Command{
	Use:   "vht",
	Short: "Vault Helper Tool",
	Long:  `A simple vault helper tool that simplifies the usage of Vault`,
}

func Execute() error {
	return rootCmd.Execute()
}

func getVaultCredentialsFromEnvironment() (addr, token string, err error) {
	addr = os.Getenv("VAULT_ADDR")
	if token = os.Getenv("VAULT_TOKEN"); token == "" {
		var data []byte
		usr, _ := user.Current()
		dir := usr.HomeDir
		data, err = ioutil.ReadFile(fmt.Sprintf("%s/.vault-token", dir))
		if err != nil {
			fmt.Println(err)
			token = ""
		} else {
			token = string(data)
		}
	}
	if addr == "" || token == "" {
		err = errors.New("missing vault address or token")
	}
	return
}

func getVaultClient() (client *api.Client, err error) {
	vaultAddr, vaultToken, err := getVaultCredentialsFromEnvironment()
	if err != nil {
		return nil, err
	}
	var httpClient = &http.Client{
		Timeout: 10 * time.Second,
	}
	if client, err = api.NewClient(&api.Config{Address: vaultAddr, HttpClient: httpClient}); err != nil {
		return nil, err
	}
	client.SetToken(vaultToken)
	return client, err
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
