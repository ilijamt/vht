package vault_test

import (
	"github.com/ilijamt/vht/internal/vault"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
	"time"
)

func TestDeletePaths(t *testing.T) {
	t.Run("Invalid vault client", func(t *testing.T) {
		require.Errorf(t, vault.DeletePaths([]string{}, nil, ioutil.Discard), vault.ErrMissingVaultClient)
	})

	t.Run("Empty paths", func(t *testing.T) {
		client, err := vault.Client()
		require.NoError(t, err)
		require.NotNil(t, client)
		require.NoError(t, vault.DeletePaths([]string{}, client, ioutil.Discard))
	})

	t.Run("Valid paths", func(t *testing.T) {
		client, err := vault.Client()
		require.NoError(t, err)
		require.NotNil(t, client)
		type Data struct {
			Test int
			Time time.Time
		}
		_, err = client.Logical().Write("secret/data/test/1", map[string]interface{}{"data": Data{
			Test: 1,
			Time: time.Now(),
		}})
		require.NoError(t, err)
		_, err = client.Logical().Write("secret/data/test/2", map[string]interface{}{"data": Data{
			Test: 2,
			Time: time.Now(),
		}})
		require.NoError(t, err)
		require.NoError(t, vault.DeletePaths([]string{
			"secret/metadata/test/1",
			"secret/metadata/test/2",
		}, client, ioutil.Discard))
	})
}
