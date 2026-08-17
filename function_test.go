package impeller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVersion(t *testing.T) {
	assert.NoError(t, Init())
	assert.Equal(t, uint32(0x20404000), GetVersion())
}
