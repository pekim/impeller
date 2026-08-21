package impeller

import (
	_ "embed"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

//go:embed boxes.ttf
var font []byte // a nice small font, that won't bloat the git repo

func TestMappingOnRelease(t *testing.T) {
	assert.NoError(t, Init())

	userData := 42
	released := false

	mapping := &Mapping{
		Data:   unsafe.Pointer(&font[0]),
		Length: uint64(len(font)),
		OnRelease: NewCallback(func(user_data unsafe.Pointer) {
			assert.Equal(t, userData, *(*int)(user_data))
			released = true
		}),
	}

	typographyContext := TypographyContextNew()
	assert.NotNil(t, typographyContext.handle)
	typographyContext.RegisterFont(mapping, unsafe.Pointer(&userData), "somefont")
	typographyContext.Release()
	assert.True(t, released)
}
