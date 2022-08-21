package vault_test

import (
	"github.com/ilijamt/vht/internal/vault"
	"github.com/stretchr/testify/require"
	"io"
	"testing"
)

func TestTree(t *testing.T) {

	t.Run("Invalid path", func(t *testing.T) {
		client, err := vault.Client()
		require.NoError(t, err)
		require.NotNil(t, client)
		paths, err := vault.TreeSerial("invalid/test", client)
		require.Empty(t, paths)
		require.NoError(t, err)
	})

	t.Run("Invalid vault client", func(t *testing.T) {
		paths, err := vault.Tree("", nil, 10)
		require.Error(t, err)
		require.Empty(t, paths)
		paths, err = vault.TreeSerial("", nil)
		require.Error(t, err)
		require.Empty(t, paths)
	})

	t.Run("Listing a path", func(t *testing.T) {
		client, err := vault.Client()
		require.NoError(t, err)
		require.NotNil(t, client)

		var paths []string
		for i := 0; i < 100; i++ {
			path, err := writeRandomData("secret/data", client, 5)
			require.NoError(t, err)
			paths = append(paths, path)
		}

		pathsTree, err := vault.Tree("secret/metadata", client, 10)
		require.NoError(t, err)
		require.NotEmpty(t, pathsTree)

		pathsSerial, err := vault.TreeSerial("secret/metadata", client)
		require.NoError(t, err)
		require.NotEmpty(t, pathsSerial)

		require.Equal(t, len(pathsSerial), len(pathsTree))

		deleted, err := vault.DeletePaths(pathsTree, 10, client, io.Discard)
		require.NoError(t, err)
		require.EqualValues(t, deleted, len(pathsTree))
		require.ElementsMatch(t, pathsSerial, pathsTree)
		require.EqualValues(t, len(paths), len(vault.FilterOnlyDataPaths(pathsTree)))
	})
}
