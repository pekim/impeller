package impeller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVersion(t *testing.T) {
	assert.NoError(t, Init())
	assert.Equal(t, uint32(0x20404000), GetVersion())
}

func TestNoArgsConstructor(t *testing.T) {
	assert.NoError(t, Init())
	assert.NotNil(t, ParagraphStyleNew())
}

func TestFloatArg(t *testing.T) {
	assert.NoError(t, Init())
	assert.NotNil(t, ImageFilterCreateDilateNew(5, 10))
}

func TestHandleArg(t *testing.T) {
	assert.NoError(t, Init())

	paint := PaintNew()
	assert.NotNil(t, paint)
	PaintSetStrokeMiter(paint, 4)
}
