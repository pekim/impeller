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
type Draw func(builder impeller.DisplayListBuilder, width int32, height int32)
type Setup func(context impeller.Context, draw RequestDraw)
type Teardown func()

type app struct {
	window            *glfw.Window
	framebufferWidth  glfw.Int
	framebufferHeight glfw.Int
	surface           *impeller.Surface
	statsStyle        impeller.ParagraphStyle
	context           impeller.Context
	builder           impeller.DisplayListBuilder
	typographyContext impeller.TypographyContext
	shouldDraw        bool
	lastFrameDuration time.Duration
}

func (app *app) init() {
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
		}),
	)

	black := impeller.Color{Red: 0.0, Green: 0.0, Blue: 0.0, Alpha: 1.0}
	blackPaint := impeller.PaintNew()
	blackPaint.SetColor(&black)

	app.statsStyle = impeller.ParagraphStyleNew()
	app.statsStyle.SetFontSize(20)
	app.statsStyle.SetForeground(blackPaint)

	app.typographyContext = impeller.TypographyContextNew()
	app.builder = impeller.DisplayListBuilderNew(nil)

	app.shouldDraw = true
}

func (app *app) createWindow(title string) {
	glfw.WindowHint(glfw.CONTEXT_CREATION_API, glfw.EGL_CONTEXT_API)
	app.window = glfw.CreateWindow(800, 600, fmt.Sprintf("Impeller Example - %s", title), nil, nil)
	app.window.MakeContextCurrent()
}

func (app *app) createContext() {
	app.context = impeller.ContextCreateOpenGLESNew(
		impeller.GetVersion(),
		impeller.NewProcAddressCallback(func(procName string, _user_data unsafe.Pointer) unsafe.Pointer {
			return unsafe.Pointer(glfw.GetProcAddress(procName))
		}),
		nil,
	)
}

func (app *app) onKey(_window *glfw.Window, key, _scancode, _action, mods glfw.Int) {
	// Quit on CTRL+Q or ESCAPE pressed.
	if (mods == glfw.MOD_CONTROL && key == glfw.KEY_Q) || (mods == 0 && key == glfw.KEY_ESCAPE) {
		app.window.SetShouldClose(glfw.TRUE)
	}
}

func (app *app) onFramebufferSize(_window *glfw.Window, _width, _height glfw.Int) {
	app.framebufferWidth, app.framebufferHeight = app.window.GetFramebufferSize()

	if app.surface != nil {
		app.surface.Release()
	}
	surfaceSize := impeller.ISize{
		Width:  int64(app.framebufferWidth),
		Height: int64(app.framebufferHeight),
	}
	surface := app.context.SurfaceCreateWrappedFBONew(0, impeller.PixelFormatRGBA8888, &surfaceSize)
	app.surface = &surface

	app.shouldDraw = true
}

func (app app) drawStats() {
	var memstats runtime.MemStats
	runtime.ReadMemStats(&memstats)

	paraBuilder := app.typographyContext.ParagraphBuilderNew()
	paraBuilder.PushStyle(app.statsStyle)

	secondsMultiplier := 1_000_000
	durationUnits := "µs"
	if app.lastFrameDuration >= time.Millisecond {
		secondsMultiplier = 1_000
		durationUnits = "ms"
	}
	paraBuilder.AddText(fmt.Sprintf(
		"Heap alloc = %1.1fM  Frame draw = %1.1f%s",
		float64(memstats.HeapAlloc)/(1_024*1_024),
		app.lastFrameDuration.Seconds()*float64(secondsMultiplier),
		durationUnits,
	))
	paraBuilder.PopStyle()
	padding := float32(20)
	para := paraBuilder.BuildParagraphNew(float32(app.framebufferWidth) - (2 * padding))
	paraHeight := para.GetHeight()
	app.builder.DrawParagraph(para, &impeller.Point{X: padding, Y: float32(app.framebufferHeight) - paraHeight - padding})
	para.Release()
	paraBuilder.Release()
}

func (app *app) setShouldDraw() {
	app.shouldDraw = true
}

func (app *app) runEventLoop(draw Draw) {
	for app.window.ShouldClose() == glfw.FALSE {
		if app.shouldDraw {
			start := time.Now()
			draw(app.builder, int32(app.framebufferWidth), int32(app.framebufferHeight))
			app.drawStats()

			dl := app.builder.CreateDisplayListNew()
			app.surface.DrawDisplayList(dl)
			dl.Release()

			app.window.SwapBuffers()
			app.lastFrameDuration = time.Since(start)
			app.shouldDraw = false
		}

		glfw.WaitEvents()
	}
}

func (app *app) teardown() {
	app.builder.Release()
	app.surface.Release()
	app.context.Release()
	app.window.Destroy()
	glfw.Terminate()
}

func Run(title string, setup Setup, teardown Teardown, draw Draw) {
	var app app
	app.init()
	app.createWindow(title)
	app.createContext()
	app.window.SetKeyCallback(glfw.KeyCallbackNew(app.onKey))
	app.window.SetFramebufferSizeCallback(glfw.FramebuffersizeCallbackNew(app.onFramebufferSize))
	app.onFramebufferSize(nil, 0, 0)

	setup(app.context, app.setShouldDraw)

	app.window.Show()
	app.runEventLoop(draw)

	teardown()
	app.teardown()
}
