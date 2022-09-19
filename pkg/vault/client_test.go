package vault_test

import (
	"github.com/ilijamt/vht/pkg/vault"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient(t *testing.T) {
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
