package vault_test

import (
	"github.com/ilijamt/vht/internal/vault"
	"github.com/stretchr/testify/require"
	"io"
	"strings"
	"testing"
)

func TestDeletePaths(t *testing.T) {
	t.Run("Invalid vault client", func(t *testing.T) {
		deleted, err := vault.DeletePaths([]string{}, 10, nil, io.Discard)
		require.EqualError(t, err, vault.ErrMissingVaultClient)
		require.EqualValues(t, deleted, 0)
	})

	t.Run("Empty paths", func(t *testing.T) {
		client, err := vault.Client()
		require.NoError(t, err)
		require.NotNil(t, client)
		deleted, err := vault.DeletePaths([]string{}, 10, client, io.Discard)
		require.NoError(t, err)
		require.EqualValues(t, deleted, 0)
	})

	t.Run("Valid paths", func(t *testing.T) {
		client, err := vault.Client()
		require.NoError(t, err)
		require.NotNil(t, client)

		var paths []string
		for i := 0; i < 100; i++ {
			path, err := writeRandomData("secret/data", client, 3)
			require.NoError(t, err)
			paths = append(paths, path)
		}

		var deletePaths []string
		for _, path := range paths {
			deletePaths = append(deletePaths, strings.ReplaceAll(path, "secret/data", "secret/metadata"))
		}

		deleted, err := vault.DeletePaths(deletePaths, 10, client, io.Discard)
		require.NoError(t, err)
		require.EqualValues(t, len(deletePaths), deleted)
	})
}
