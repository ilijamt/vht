package vault_test

import (
	"github.com/ilijamt/vht/internal/vault"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTree(t *testing.T) {
	t.Run("Invalid vault client", func(t *testing.T) {
		paths, err := vault.Tree("", nil)
		require.Error(t, err)
		require.Empty(t, paths)
	})

	t.Run("Listing a path", func(t *testing.T) {

	})
}
