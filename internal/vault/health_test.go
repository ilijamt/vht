package vault_test

import (
	"github.com/ilijamt/vht/internal/vault"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func TestHealth(t *testing.T) {
	t.Run("Invalid vault client", func(t *testing.T) {
		require.EqualError(t, vault.Health(ioutil.Discard, nil), vault.ErrMissingVaultClient)
	})

	t.Run("Healthy", func(t *testing.T) {
		client, err := vault.Client()
		require.NoError(t, err)
		require.NotNil(t, client)
		require.NoError(t, vault.Health(ioutil.Discard, client))
	})
}
