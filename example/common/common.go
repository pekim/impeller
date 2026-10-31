package common

import (
	"fmt"
	"runtime"
	"time"
	"unsafe"

	"github.com/pekim/go-glfw"
	"github.com/pekim/impeller"
)

// The approach used in this code is not necessarily a good one,
// or recommended. It's simply one that's good enough for the examples.

func init() {
	// This is important to ensure that only a single thread makes calls to
	// openGL apis.
	runtime.LockOSThread()
}

type RequestDraw func()

func Run(
	title string,
	setup func(window *glfw.Window, context impeller.Context, draw RequestDraw),
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

	typographyContext := impeller.TypographyContextNew()

	black := impeller.Color{Red: 0.0, Green: 0.0, Blue: 0.0, Alpha: 1.0}
	blackPaint := impeller.PaintNew()
	blackPaint.SetColor(&black)

	style := impeller.ParagraphStyleNew()
	style.SetFontSize(20)
	style.SetForeground(blackPaint)

	framebufferWidth, framebufferHeight := window.GetFramebufferSize()
	surfaceSize := impeller.ISize{
		Width:  int64(framebufferWidth),
		Height: int64(framebufferHeight),
	}
	surface := context.SurfaceCreateWrappedFBONew(0, impeller.PixelFormatRGBA8888, &surfaceSize)

	builder := impeller.DisplayListBuilderNew(nil)

	shouldDraw := true

	// Quit on CTRL+Q or ESCAPE pressed.
	window.SetKeyCallback(glfw.KeyCallbackNew(func(window *glfw.Window, key, _scancode, _action, mods glfw.Int) {
		if (mods == glfw.MOD_CONTROL && key == glfw.KEY_Q) || (mods == 0 && key == glfw.KEY_ESCAPE) {
			window.SetShouldClose(glfw.TRUE)
		}
	}))

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

	setup(window, context, setShouldDraw)

	window.Show()

	var lastFrameDuration time.Duration

	drawStats := func() {
		var memstats runtime.MemStats
		runtime.ReadMemStats(&memstats)

		paraBuilder := typographyContext.ParagraphBuilderNew()
		paraBuilder.PushStyle(style)

		secondsMultiplier := 1_000_000
		durationUnits := "µs"
		if lastFrameDuration >= time.Millisecond {
			secondsMultiplier = 1_000
			durationUnits = "ms"
		}
		paraBuilder.AddText(fmt.Sprintf(
			"Heap alloc = %1.1fM  Frame draw = %1.1f%s",
			float64(memstats.HeapAlloc)/(1_024*1_024),
			lastFrameDuration.Seconds()*float64(secondsMultiplier),
			durationUnits,
		))
		paraBuilder.PopStyle()
		padding := float32(20)
		para := paraBuilder.BuildParagraphNew(float32(framebufferWidth) - (2 * padding))
		paraHeight := para.GetHeight()
		builder.DrawParagraph(para, &impeller.Point{X: padding, Y: float32(framebufferHeight) - paraHeight - padding})
		para.Release()
		paraBuilder.Release()
	}

	for window.ShouldClose() == glfw.FALSE {
		if shouldDraw {
			start := time.Now()
			draw(builder, int32(framebufferWidth), int32(framebufferHeight))
			drawStats()

			dl := builder.CreateDisplayListNew()
			surface.DrawDisplayList(dl)
			dl.Release()

			window.SwapBuffers()
			lastFrameDuration = time.Since(start)
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
