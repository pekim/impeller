package impeller

import (
	"testing"

	// "github.com/pekim/gl-purego/gl/v3.3-core/gl"
	// "github.com/pekim/gl-purego/glfw"
	"github.com/stretchr/testify/assert"
)

func TestMappingOnRelease(t *testing.T) {
	assert.NoError(t, Init())
	// assert.NoError(t, glfw.Initialize())
	// assert.True(t, glfw.Init())
	// // assert.NoError(t, gl.Initialize())

	// glfw.WindowHint(glfw.CONTEXT_VERSION_MAJOR, 3)
	// glfw.WindowHint(glfw.CONTEXT_VERSION_MINOR, 3)
	// glfw.WindowHint(glfw.OPENGL_PROFILE, glfw.OPENGL_CORE_PROFILE)
	// window := glfw.CreateWindow(800, 600, "test", 0, 0)
	// assert.NotZero(t, window)
	// window.MakeContextCurrent()

	// fmt.Println("glGetError", glfw.GetProcAddress("glGetError"))

	// context := ContextCreateOpenGLESNew(
	// 	GetVersion(),
	// 	NewProcAddressCallback(func(procName string, user_data unsafe.Pointer) {
	// 		fmt.Println("proc address callback :", procName, user_data)
	// 	}),
	// 	unsafe.Pointer(uintptr(42)),
	// )
	// fmt.Println("context", context)

	// typographyContext := TypographyContextNew()
	// fmt.Println("typographyContext", typographyContext)

	// data := []byte(" ")
	// mapping := &Mapping{
	// 	Data:   unsafe.Pointer(&data[0]),
	// 	Length: uint64(len(data)),
	// 	OnRelease: NewCallback(func(user_data unsafe.Pointer) {
	// 		fmt.Println("release", user_data)
	// 	}),
	// }

	// fragmentProgram := FragmentProgramNew(mapping, unsafe.Pointer(uintptr(42)))
	// fragmentProgram.Release()
}
