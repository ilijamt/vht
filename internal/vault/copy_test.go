package vault_test
//
//import (
//	"github.com/ilijamt/vht/internal/vault"
//	"github.com/stretchr/testify/require"
//	"io/ioutil"
//	"math/rand"
//	"os"
//	"regexp"
//	"testing"
//	"time"
//)
//
//func TestCopyPath(t *testing.T) {
//
//	t.Run("Invalid vault client", func(t *testing.T) {
//		deleted, err := vault.CopyPath("/path", "/new-path", 10, nil, ioutil.Discard)
//		require.Errorf(t, err, vault.ErrMissingVaultClient)
//		require.EqualValues(t, 0, deleted)
//	})
//
//	t.Run("Missing path", func(t *testing.T) {
//		client, err := vault.Client()
//		require.NoError(t, err)
//		require.NotNil(t, client)
//		deleted, err := vault.CopyPath("paths", "", 10, client, ioutil.Discard)
//		require.Errorf(t, err, vault.ErrMissingPath)
//		require.EqualValues(t, 0, deleted)
//		deleted, err = vault.CopyPath("", "paths", 10, client, ioutil.Discard)
//		require.Errorf(t, err, vault.ErrMissingPath)
//		require.EqualValues(t, 0, deleted)
//	})
//
//	t.Run("Copying path to a new path", func(t *testing.T) {
//		client, err := vault.Client()
//		require.NoError(t, err)
//		require.NotNil(t, client)
//		type Data struct {
//			Test int
//			Time time.Time
//		}
//
//		var write = func(path string) {
//			_, err = client.Logical().Write(path, map[string]interface{}{"data": Data{Test: rand.Intn(1000), Time: time.Now()}})
//			require.NoError(t, err)
//		}
//		write("secret/data/copy-path/path1/1")
//		write("secret/data/copy-path/path1/level1/2")
//		write("secret/data/copy-path/path1/level1/level2/3")
//		write("secret/data/copy-path/path1/level1/level2/level3/4")
//		write("secret/data/copy-path/path1/2")
//		write("secret/data/copy-path/path1/level2/3")
//		write("secret/data/copy-path/path1/level2/level3/4")
//		write("secret/data/copy-path/path1/level2/level3/level4/5")
//
//		paths, err := vault.Tree("secret/metadata/copy-path", client, 10)
//		require.NoError(t, err)
//		r, _ := regexp.Compile(".*")
//		filteredPaths := vault.FilterDataPaths(paths, r)
//		require.Len(t, filteredPaths, 8)
//
//		var written uint64
//		written, err = vault.CopyPath("secret/metadata/copy-path", "secret/data/copy-path-new", 10, client, os.Stdout)
//		require.NoError(t, err)
//		require.EqualValues(t, written, 8)
//
//		deleted, err := vault.DeletePaths([]string{
//			"secret/metadata/copy-path/path1/1",
//			"secret/metadata/copy-path/path1/level1/2",
//			"secret/metadata/copy-path/path1/level1/level2/3",
//			"secret/metadata/copy-path/path1/level1/level2/level3/4",
//			"secret/metadata/copy-path/path1/2",
//			"secret/metadata/copy-path/path1/level2/3",
//			"secret/metadata/copy-path/path1/level2/level3/4",
//			"secret/metadata/copy-path/path1/level2/level3/level4/5",
//		}, client, ioutil.Discard)
//
//		require.NoError(t, err)
//		require.EqualValues(t, 0, deleted)
//
//		deleted, err = vault.DeletePaths([]string{
//			"secret/metadata/copy-path-new/path1/1",
//			"secret/metadata/copy-path-new/path1/level1/2",
//			"secret/metadata/copy-path-new/path1/level1/level2/3",
//			"secret/metadata/copy-path-new/path1/level1/level2/level3/4",
//			"secret/metadata/copy-path-new/path1/2",
//			"secret/metadata/copy-path-new/path1/level2/3",
//			"secret/metadata/copy-path-new/path1/level2/level3/4",
//			"secret/metadata/copy-path-new/path1/level2/level3/level4/5",
//		}, client, ioutil.Discard)
//
//		require.NoError(t, err)
//		require.EqualValues(t, 8, deleted)
//	})
//
//}
