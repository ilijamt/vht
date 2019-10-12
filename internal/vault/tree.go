package vault

import (
	"errors"
	"fmt"
	"github.com/hashicorp/vault/api"
	"strings"
)

func getPathDetails(path string, client *api.Client) (nodes, leafs []string, err error) {
	var secret *api.Secret
	secret, err = client.Logical().List(path)
	if err != nil {
		return
	}
	if secret == nil || secret.Data == nil || secret.Data["keys"] == nil {
		return []string{}, []string{}, nil
	}

	keys := secret.Data["keys"].([]interface{})
	var newPath string
	for _, p := range keys {
		switch v := p.(type) {
		case string:
			newPath = strings.ReplaceAll(fmt.Sprintf("%s/%s", path, v), "//", "/")
			if v[len(v)-1:] == "/" {
				nodes = append(nodes, newPath)
			} else {
				leafs = append(leafs, newPath)
			}
		default:
		}
	}

	return
}

// Tree fetches a list of all the paths in Vault.
// @TODO: Rebuild it to be concurrent instead of serial
func Tree(rootPath string, client *api.Client, concurrency int8) (paths []string, err error) {
	if client == nil {
		return []string{}, errors.New(ErrMissingVaultClient)
	}

	var nodes, leafs []string
	nodes, leafs, err = getPathDetails(rootPath, client)
	if err != nil || (len(nodes) == 0 && len(leafs) == 0) {
		return []string{}, err
	}

	paths = append(paths, nodes...)
	paths = append(paths, leafs...)

	// Iterate over the nodes so we can get the new data
	for _, path := range nodes {
		newPath := strings.ReplaceAll(fmt.Sprintf("%s", path), "//", "/")
		pths, err := Tree(newPath, client, concurrency)
		if err != nil {
			return []string{}, err
		}
		paths = append(paths, pths...)
	}

	return paths, err
}
