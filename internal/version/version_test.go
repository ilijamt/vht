package version_test

import (
	"bufio"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ilijamt/vht/internal/version"
)

func TestPrintVersion(t *testing.T) {
	wr := bufio.NewWriter(io.Discard)
	version.PrintVersion(wr)
	assert.Greater(t, wr.Buffered(), 0)
	assert.NoError(t, wr.Flush())
}
