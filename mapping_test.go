package impeller

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed boxes.ttf
var font []byte // a nice small font, that won't bloat the git repo

func TestMappingOnRelease(t *testing.T) {
	assert.NoError(t, Init())

	userData := uintptr(42)
	released := false

	mapping := &Mapping{
		Data:   &font[0],
		Length: uint64(len(font)),
		OnRelease: NewCallback(func(user_data uintptr) {
			assert.Equal(t, userData, user_data)
			released = true
		}),
	}

	typographyContext := TypographyContextNew()
	assert.NotNil(t, typographyContext.handle)
	typographyContext.RegisterFont(mapping, userData, "somefont")
	typographyContext.Release()
	assert.True(t, released)
}
