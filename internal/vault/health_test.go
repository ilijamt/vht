package vault_test

import (
	"github.com/ilijamt/vht/internal/vault"
	v "github.com/ilijamt/vht/pkg/vault"
	"github.com/stretchr/testify/require"
	"io"
	"testing"
)

func TestHealth(t *testing.T) {
	t.Run("Invalid vault client", func(t *testing.T) {
		require.EqualError(t, vault.Health(io.Discard, nil), vault.ErrMissingVaultClient)
	})

	t.Run("Healthy", func(t *testing.T) {
		client, err := v.Client()
		require.NoError(t, err)
		require.NotNil(t, client)
		require.NoError(t, vault.Health(io.Discard, client))
	})
}
