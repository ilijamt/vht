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
	t.Run("Invalid vault client", func(t *testing.T) {
		paths, err := vault.Tree("", nil, 10)
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
			require.NoError(t, err)
			_, err = client.Logical().Write(fmt.Sprintf("secret/data/%s/%s/%s", uroot.String(), "root", uuid1.String()), map[string]interface{}{"data": Data{
				Test: uuid1.String(),
				Time: time.Now(),
			}})
			_, err = client.Logical().Write(fmt.Sprintf("secret/data/%s/root/%s/%s", uroot.String(), uuid2.String(), uuid3.String()), map[string]interface{}{"data": Data{
				Test: uuid3.String(),
				Time: time.Now(),
			}})
		}

		paths, err := vault.Tree(fmt.Sprintf("secret/metadata/%s", uroot.String()), client, 10)
		require.NoError(t, err)
		require.NotEmpty(t, paths)
		for _, path := range paths {
			fmt.Println(path)
		}
		require.Len(t, paths, 31)
		require.NoError(t, vault.DeletePaths(paths, client, ioutil.Discard))
	})
}
