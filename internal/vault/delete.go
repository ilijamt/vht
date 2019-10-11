package vault

import (
	"errors"
	"fmt"
	"github.com/hashicorp/vault/api"
	"io"
)

// DeletePaths removes all the paths in the list from Vault
func DeletePaths(paths []string, client *api.Client, w io.Writer) (err error) {
	if client == nil {
		return errors.New(ErrMissingVaultClient)
	}
	var path string
	for _, path = range paths {
		_, err = client.Logical().Delete(path)
		_, _ = fmt.Fprintf(w, "%s path deleted: %t\n", path, err == nil)
	}
	return nil
}
