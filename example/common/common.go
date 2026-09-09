package common

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/pekim/glfw"
	"github.com/pekim/impeller"
)

func init() {
	// This is important to ensure that only a single thread makes calls to
	// openGL apis.
	runtime.LockOSThread()
}

type RequestDraw func()

func Run(
	title string,
	setup func(window *glfw.Window, draw RequestDraw),
	teardown func(),
	draw func(builder impeller.DisplayListBuilder, width int32, height int32),
) {
	if err := impeller.Init(); err != nil {
		panic(err)
	}

	if err := glfw.Initialise(); err != nil {
		panic(err)
	}
	glfw.Init()

	glfw.SetErrorCallback(glfw.ErrorCallbackNew(
		func(error_code glfw.Int, description string) {
			fmt.Printf("GLFW Error (%d): %s\n", error_code, description)
		}))

	glfw.WindowHint(glfw.CONTEXT_CREATION_API, glfw.EGL_CONTEXT_API)
	window := glfw.CreateWindow(800, 600, fmt.Sprintf("Impeller Example - %s", title), nil, nil)
	window.MakeContextCurrent()

	context := impeller.ContextCreateOpenGLESNew(
		impeller.GetVersion(),
		impeller.NewProcAddressCallback(func(procName string, _user_data unsafe.Pointer) unsafe.Pointer {
			return unsafe.Pointer(glfw.GetProcAddress(procName))
		}),
		nil,
	)

	framebufferWidth, framebufferHeight := window.GetFramebufferSize()
	surfaceSize := impeller.ISize{
		Width:  int64(framebufferWidth),
		Height: int64(framebufferHeight),
	}
	surface := context.SurfaceCreateWrappedFBONew(0, impeller.PixelFormatRGBA8888, &surfaceSize)

	builder := impeller.DisplayListBuilderNew(nil)

	shouldDraw := true

	window.SetFramebufferSizeCallback(glfw.FramebuffersizeCallbackNew(func(_window *glfw.Window, width, height glfw.Int) {
		framebufferWidth = width
		framebufferHeight = height

		surface.Release()
		surfaceSize := impeller.ISize{
			Width:  int64(framebufferWidth),
			Height: int64(framebufferHeight),
		}
		surface = context.SurfaceCreateWrappedFBONew(0, impeller.PixelFormatRGBA8888, &surfaceSize)

		shouldDraw = true
	}))

	setShouldDraw := func() {
		shouldDraw = true
	}

	setup(window, setShouldDraw)

	window.Show()

	for window.ShouldClose() == glfw.FALSE {
		if shouldDraw {
			draw(builder, int32(framebufferWidth), int32(framebufferHeight))

			dl := builder.CreateDisplayListNew()
			surface.DrawDisplayList(dl)
			dl.Release()

			window.SwapBuffers()
			shouldDraw = false
		}

		glfw.WaitEvents()
	}

	teardown()
	builder.Release()
	surface.Release()
	context.Release()
	window.Destroy()

	glfw.Terminate()
}
