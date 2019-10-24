package vault_test

import (
	"github.com/ilijamt/envwrap"
	"github.com/ilijamt/vht/internal/vault"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient(t *testing.T) {
	t.Run("No environment variables", func(t *testing.T) {
		env := envwrap.NewStorage()
		defer env.ReleaseAll()
		env.Store("VAULT_ADDR", "")
		env.Store("VAULT_TOKEN", "")
		client, err := vault.Client()
		require.EqualError(t, err, vault.ErrMissingVaultAddrOrCredentials)
		require.Nil(t, client)
	})

	t.Run("With environment variables", func(t *testing.T) {
		client, err := vault.Client()
		require.NoError(t, err)
		require.NotNil(t, client)
		secret, err := client.Sys().Health()
		require.NoError(t, err)
		require.NotNil(t, secret)
		require.False(t, secret.Sealed)
	})
}
