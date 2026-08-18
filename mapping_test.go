package impeller

import (
	"fmt"
	"testing"
	"unsafe"

	// "github.com/pekim/gl-purego/gl/v3.3-core/gl"
	"github.com/pekim/gl-purego/glfw"
	"github.com/stretchr/testify/assert"
)

func TestMappingOnRelease(t *testing.T) {
	assert.NoError(t, Init())
	assert.NoError(t, glfw.Initialize())
	assert.True(t, glfw.Init())
	// assert.NoError(t, gl.Initialize())

	glfw.WindowHint(glfw.CONTEXT_VERSION_MAJOR, 3)
	glfw.WindowHint(glfw.CONTEXT_VERSION_MINOR, 3)
	glfw.WindowHint(glfw.OPENGL_PROFILE, glfw.OPENGL_CORE_PROFILE)
	window := glfw.CreateWindow(800, 600, "test", 0, 0)
	assert.NotZero(t, window)
	window.MakeContextCurrent()

	context := ContextCreateOpenGLESNew(
		GetVersion(),
		NewProcAddressCallback(func(procName string, _user_data unsafe.Pointer) unsafe.Pointer {
			return unsafe.Pointer(glfw.GetProcAddress(procName))
		}),
		unsafe.Pointer(uintptr(42)),
	)
	assert.NotNil(t, context.handle)

	typographyContext := TypographyContextNew()
	assert.NotNil(t, typographyContext.handle)

	data := []byte(" ")
	mapping := &Mapping{
		Data:   unsafe.Pointer(&data[0]),
		Length: uint64(len(data)),
		OnRelease: NewCallback(func(user_data unsafe.Pointer) {
			fmt.Println("release", user_data)
		}),
	}

	fragmentProgram := FragmentProgramNew(mapping, unsafe.Pointer(uintptr(42)))
	assert.NotNil(t, fragmentProgram.handle)
	fragmentProgram.Release()
}
