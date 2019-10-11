package vault

import (
	"errors"
	"fmt"
	"github.com/hashicorp/vault/api"
	"strings"
)

// Tree fetches a list of all the paths in Vault.
// @TODO: Rebuild it to be concurrent instead of serial
func Tree(rootPath string, client *api.Client) (paths []string, err error) {
	if client == nil {
		return []string{}, errors.New(ErrMissingVaultClient)
	}
	var secret *api.Secret
	secret, err = client.Logical().List(rootPath)
	if err != nil {
		return []string{}, err
	}
	if secret == nil || secret.Data == nil || secret.Data["keys"] == nil {
		return []string{}, nil
	}
	keys := secret.Data["keys"].([]interface{})
	for _, path := range keys {
		switch v := path.(type) {
		case string:
			lastChar := v[len(v)-1:]
			if lastChar == "/" {
				newPath := strings.ReplaceAll(fmt.Sprintf("%s/%s", rootPath, v), "//", "/")
				p, e := Tree(newPath, client)
				if e != nil {
					return []string{}, err
				}
				paths = append(paths, newPath)
				paths = append(paths, p...)
			} else {
				paths = append(paths, strings.ReplaceAll(fmt.Sprintf("%s/%s", rootPath, v), "//", "/"))
			}
		default:
			fmt.Printf("Unkown type of key: %v\n", v)
		}
	}
	return paths, err
}
