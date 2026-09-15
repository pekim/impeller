package impeller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVersion(t *testing.T) {
	assert.NoError(t, Init())
	assert.Equal(t, uint32(0x20404000), GetVersion())
	assert.Equal(t, uint32(1), GetVersionMajor())
	assert.Equal(t, uint32(4), GetVersionMinor())
	assert.Equal(t, uint32(0), GetVersionPatch())
	assert.Equal(t, uint32(1), GetVersionVariant())
}

func TestNoArgsConstructor(t *testing.T) {
	assert.NoError(t, Init())
	assert.NotNil(t, ParagraphStyleNew().handle)
}

func TestFloatArg(t *testing.T) {
	assert.NoError(t, Init())
	assert.NotNil(t, ImageFilterCreateDilateNew(5, 10).handle)
}

func TestHandleArg(t *testing.T) {
	assert.NoError(t, Init())

	paint := PaintNew()
	assert.NotNil(t, paint)
	Paint.SetStrokeMiter(paint, 4)
}
