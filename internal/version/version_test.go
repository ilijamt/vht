package version_test

import (
	"bufio"
	"github.com/ilijamt/vht/internal/version"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestPrintVersion(t *testing.T) {
	wr := bufio.NewWriter(io.Discard)
	version.PrintVersion(wr)
	assert.Equal(t, 71, wr.Buffered())
	assert.NoError(t, wr.Flush())
}
