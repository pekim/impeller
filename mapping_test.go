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

	expectedUserData := unsafe.Pointer(uintptr(42))
	released := false
	mapping := &Mapping{
		Data:   unsafe.Pointer(&font[0]),
		Length: uint64(len(font)),
		OnRelease: NewCallback(func(user_data unsafe.Pointer) {
			assert.Equal(t, expectedUserData, user_data)
			released = true
		}),
	}

	typographyContext := TypographyContextNew()
	assert.NotNil(t, typographyContext.handle)
	typographyContext.RegisterFont(mapping, unsafe.Pointer(uintptr(42)), "somefont")
	typographyContext.Release()
	assert.True(t, released)
}

// "github.com/pekim/gl-purego/gl/v3.3-core/gl"
// "github.com/pekim/gl-purego/glfw"

//go :embed temp.spirv
// var glsl []byte

// assert.NoError(t, glfw.Initialize())
// assert.True(t, glfw.Init())
// // assert.NoError(t, gl.Initialize())

// glfw.WindowHint(glfw.CONTEXT_VERSION_MAJOR, 3)
// glfw.WindowHint(glfw.CONTEXT_VERSION_MINOR, 3)
// glfw.WindowHint(glfw.OPENGL_PROFILE, glfw.OPENGL_CORE_PROFILE)
// window := glfw.CreateWindow(800, 600, "test", 0, 0)
// assert.NotZero(t, window)
// window.MakeContextCurrent()

// context := ContextCreateOpenGLESNew(
// 	GetVersion(),
// 	NewProcAddressCallback(func(procName string, _user_data unsafe.Pointer) unsafe.Pointer {
// 		return unsafe.Pointer(glfw.GetProcAddress(procName))
// 	}),
// 	unsafe.Pointer(uintptr(42)),
// )
// assert.NotNil(t, context.handle)

// data := glsl
// mapping := &Mapping{
// 	Data:   unsafe.Pointer(&data[0]),
// 	Length: uint64(len(data)),
// 	OnRelease: NewCallback(func(user_data unsafe.Pointer) {
// 		fmt.Println("release", user_data)
// 	}),
// }

// fragmentProgram := FragmentProgramNew(mapping, unsafe.Pointer(uintptr(42)))
// assert.NotNil(t, fragmentProgram.handle)
// fragmentProgram.Release()
