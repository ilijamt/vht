package vault_test

import (
	"github.com/ilijamt/vht/internal/vault"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsV2(t *testing.T) {
	t.Run("Invalid vault client", func(t *testing.T) {
		yes, err := vault.IsV2("secret", nil)
		require.Errorf(t, err, vault.ErrMissingVaultClient)
		require.False(t, yes)
	})

	t.Run("No mount point", func(t *testing.T) {
		client, err := vault.Client()
		require.NoError(t, err)
		require.NotNil(t, client)
		yes, err := vault.IsV2("", client)
		require.Errorf(t, err, vault.ErrMissingPath)
		require.False(t, yes)
	})

	t.Run("KV is version 2", func(t *testing.T) {
		client, err := vault.Client()
		require.NoError(t, err)
		require.NotNil(t, client)
		yes, err := vault.IsV2("secret/", client)
		require.NoError(t, err)
		require.True(t, yes)
	})

	t.Run("Non existing mount point", func(t *testing.T) {
		client, err := vault.Client()
		require.NoError(t, err)
		require.NotNil(t, client)
		yes, err := vault.IsV2("missing/", client)
		require.NoError(t, err)
		require.False(t, yes)
	})

}

func TestIsKV(t *testing.T) {
	t.Run("Invalid vault client", func(t *testing.T) {
		yes, err := vault.IsKV("secret", nil)
		require.Errorf(t, err, vault.ErrMissingVaultClient)
		require.False(t, yes)
	})

	t.Run("No mount point", func(t *testing.T) {
		client, err := vault.Client()
		require.NoError(t, err)
		require.NotNil(t, client)
		yes, err := vault.IsKV("", client)
		require.Errorf(t, err, vault.ErrMissingPath)
		require.False(t, yes)
	})

	t.Run("Should be KV engine", func(t *testing.T) {
		client, err := vault.Client()
		require.NoError(t, err)
		require.NotNil(t, client)
		yes, err := vault.IsKV("secret/", client)
		require.NoError(t, err)
		require.True(t, yes)

	})

	t.Run("Should not be KV engine", func(t *testing.T) {
		client, err := vault.Client()
		require.NoError(t, err)
		require.NotNil(t, client)
		yes, err := vault.IsV2("cubbyhole/", client)
		require.NoError(t, err)
		require.False(t, yes)
	})

	t.Run("Non existing mount point", func(t *testing.T) {
		client, err := vault.Client()
		require.NoError(t, err)
		require.NotNil(t, client)
		yes, err := vault.IsKV("missing/", client)
		require.NoError(t, err)
		require.False(t, yes)
	})

}
