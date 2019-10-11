package vault

import (
	"errors"
	"fmt"
	"github.com/hashicorp/vault/api"
	"io"
)

// Health checks the vault info and prints out the details in the writer
func Health(w io.Writer, client *api.Client) (err error) {
	if client == nil {
		return errors.New(ErrMissingVaultClient)
	}
	response, err := client.Sys().Health()
	if err != nil {
		return err
	}

	_, _ = fmt.Fprintf(w, "Cluster name: %s\n", response.ClusterName)
	_, _ = fmt.Fprintf(w, "Version: %s\n", response.Version)
	_, _ = fmt.Fprintf(w, "Sealed: %t\n", response.Sealed)
	return nil
}
