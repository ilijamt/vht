package vault

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/vault/api"
	"golang.org/x/sync/semaphore"
	"io"
	"sync/atomic"
)

// DeletePaths removes all the paths in the list from Vault
func DeletePaths(paths []string, concurrency int8, client *api.Client, w io.Writer) (deleted uint64, err error) {
	if client == nil {
		return deleted, errors.New(ErrMissingVaultClient)
	}
	var path string
	var sem = semaphore.NewWeighted(int64(concurrency))

	for _, path = range paths {
		if err = sem.Acquire(context.Background(), 1); err != nil {
			return deleted, err
		}

		go func(p string) {
			defer sem.Release(1)
			_, err := client.Logical().Delete(p)
			atomic.AddUint64(&deleted, 1)
			_, _ = fmt.Fprintf(w, "%s path deleted: %t\n", p, err == nil)
		}(path)

	}
	return deleted, sem.Acquire(context.Background(), int64(concurrency))
}
