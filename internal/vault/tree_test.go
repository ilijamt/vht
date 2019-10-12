package vault_test

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/ilijamt/vht/internal/vault"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
	"time"
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
		type Data struct {
			Test string
			Time time.Time
		}
		uroot := uuid.New()
		for i := 0; i < 10; i++ {
			uuid1 := uuid.New()
			uuid2 := uuid.New()
			uuid3 := uuid.New()
			uuid4 := uuid.New()
			_, err = client.Logical().Write(fmt.Sprintf("secret/data/%s/%s", uroot.String(), uuid4.String()), map[string]interface{}{"data": Data{
				Test: uuid1.String(),
				Time: time.Now(),
			}})
			_, err = client.Logical().Write(fmt.Sprintf("secret/data/%s/%s/%s", uroot.String(), "root", uuid1.String()), map[string]interface{}{"data": Data{
				Test: uuid1.String(),
				Time: time.Now(),
			}})
			_, err = client.Logical().Write(fmt.Sprintf("secret/data/%s/root/%s/%s", uroot.String(), uuid2.String(), uuid3.String()), map[string]interface{}{"data": Data{
				Test: uuid3.String(),
				Time: time.Now(),
			}})
		}

		paths, err := vault.Tree(fmt.Sprintf("secret/metadata/%s", uroot.String()), client, 2)
		require.NoError(t, err)
		require.NotEmpty(t, paths)
		require.Len(t, paths, 41)

		pathsSerial, err := vault.TreeSerial(fmt.Sprintf("secret/metadata/%s", uroot.String()), client)
		require.NoError(t, err)
		require.NotEmpty(t, pathsSerial)
		require.Len(t, pathsSerial, 41)
		require.Equal(t, len(pathsSerial), len(paths))
		require.NoError(t, vault.DeletePaths(paths, client, ioutil.Discard))
		require.ElementsMatch(t, pathsSerial, paths)
	})
}
