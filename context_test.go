package impeller

import (
	"testing"
	"unsafe"

	"github.com/pekim/gl-purego/glfw"
	"github.com/stretchr/testify/assert"
)

func TestContext(t *testing.T) {
	assert.NoError(t, Init())

	assert.NoError(t, glfw.Initialise())
	assert.NotZero(t, glfw.Init())

	glfw.WindowHint(glfw.CONTEXT_CREATION_API, glfw.EGL_CONTEXT_API)
	window := glfw.CreateWindow(800, 600, "test", nil, nil)
	assert.NotZero(t, window)
	window.MakeContextCurrent()

	data := 42
	dataPtr := unsafe.Pointer(&data)
	context := ContextCreateOpenGLESNew(
		GetVersion(),
		NewProcAddressCallback(func(procName string, user_data unsafe.Pointer) unsafe.Pointer {
			assert.Equal(t, data, *(*int)(user_data))
			return unsafe.Pointer(glfw.GetProcAddress(procName))
		}),
		dataPtr,
	)
	assert.NotNil(t, context.handle)
}
