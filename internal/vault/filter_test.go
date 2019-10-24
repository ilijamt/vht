package vault_test

import (
	"github.com/ilijamt/vht/internal/vault"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestFilterDataPaths(t *testing.T) {

	t.Run("Empty", func(t *testing.T) {
		assert.Empty(t, vault.FilterOnlyDataPaths([]string{}))
	})

	t.Run("No path matches", func(t *testing.T) {
		r, _ := regexp.Compile("demo")
		assert.Empty(t, vault.FilterDataPaths([]string{
			"path/test/01",
			"path/test/02",
		}, r))
	})

	t.Run("Match found", func(t *testing.T) {
		r1, _ := regexp.Compile("test")
		assert.Len(t, vault.FilterDataPaths([]string{
			"path/test/01",
			"path/test/02",
		}, r1), 2)

		r2, _ := regexp.Compile("^test")
		assert.Empty(t, vault.FilterDataPaths([]string{
			"path/test/01",
			"path/test/02",
		}, r2))

		r3, _ := regexp.Compile("^path")
		assert.Len(t, vault.FilterDataPaths([]string{
			"path/test/01",
			"fpath/test/02",
		}, r3), 1)
	})

}
