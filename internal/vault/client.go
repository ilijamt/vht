package vault

import (
	"errors"
	"fmt"
	"github.com/hashicorp/vault/api"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"time"
)

func getVaultCredentialsFromEnvironment() (addr, token string, err error) {
	addr = os.Getenv("VAULT_ADDR")
	if token = os.Getenv("VAULT_TOKEN"); token == "" {
		var data []byte
		usr, _ := user.Current()
		dir := usr.HomeDir
		data, err = ioutil.ReadFile(fmt.Sprintf("%s/.vault-token", dir))
		if err != nil {
			token = ""
		} else {
			token = string(data)
		}
	}
	if addr == "" || token == "" {
		err = errors.New(ErrMissingVaultAddrOrCredentials)
	}
	return
}

func Client() (client *api.Client, err error) {
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
