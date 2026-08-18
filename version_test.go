package impeller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVersionParts(t *testing.T) {
	assert.NoError(t, Init())

	assert.Equal(t, uint32(1), GetVersionVariant(), "variant")
	assert.Equal(t, uint32(1), GetVersionMajor(), "major")
	assert.Equal(t, uint32(4), GetVersionMinor(), "minor")
	assert.Equal(t, uint32(0), GetVersionPatch(), "patch")
}
