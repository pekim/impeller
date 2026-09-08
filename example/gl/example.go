package main

import (
	_ "embed"
	"fmt"
	"runtime"
	"unsafe"

	"github.com/pekim/glfw"
	"github.com/pekim/impeller"
)

//go:embed DejaVuSans.ttf
var dejaVuSansTTF []byte

func init() {
	// This is important to ensure that only a single thread makes calls to
	// openGL apis.
	runtime.LockOSThread()
}

func main() {
	err := glfw.Initialise()
	if err != nil {
		panic(err)
	}
	err = impeller.Init()
	if err != nil {
		panic(err)
	}

	glfw.Init()

	glfw.SetErrorCallback(glfw.ErrorCallbackNew(
		func(error_code glfw.Int, description string) {
			fmt.Printf("GLFW Error (%d): %s\n", error_code, description)
		}))

	glfw.WindowHint(glfw.CONTEXT_CREATION_API, glfw.EGL_CONTEXT_API)
	window := glfw.CreateWindow(800, 600, "Impeller Example (OpenGL)", nil, nil)

	framebufferWidth, framebufferHeight := window.GetFramebufferSize()

	window.MakeContextCurrent()

	context := impeller.ContextCreateOpenGLESNew(
		impeller.GetVersion(),
		impeller.NewProcAddressCallback(func(procName string, _user_data unsafe.Pointer) unsafe.Pointer {
			return unsafe.Pointer(glfw.GetProcAddress(procName))
		}),
		nil,
	)

	surfaceSize := impeller.ISize{
		Width:  int64(framebufferWidth),
		Height: int64(framebufferHeight),
	}
	surface := context.SurfaceCreateWrappedFBONew(0, impeller.PixelFormatRGBA8888, &surfaceSize)

	builder := impeller.DisplayListBuilderNew(nil)
	paint := impeller.PaintNew()

	// Clear the background to a white color.
	clearColour := impeller.Color{Red: 1.0, Green: 1.0, Blue: 1.0, Alpha: 1.0}
	paint.SetColor(&clearColour)
	builder.DrawPaint(paint)

	// Draw a red box.
	boxColour := impeller.Color{Red: 1.0, Green: 0.0, Blue: 0.0, Alpha: 1.0}
	paint.SetColor(&boxColour)
	boxRect := impeller.Rect{X: 10, Y: 10, Width: 100, Height: 100}
	builder.DrawRect(&boxRect, paint)

	typographyContext := impeller.TypographyContextNew()
	fontMapping := impeller.Mapping{
		Data:   unsafe.Pointer(&dejaVuSansTTF[0]),
		Length: uint64(len(dejaVuSansTTF)),
	}
	dejaVuSans := "DejaVuSans"
	fmt.Println(typographyContext.RegisterFont(&fontMapping, nil, dejaVuSans))

	paraStyle := impeller.ParagraphStyleNew()
	paraStyle.SetFontSize(1.25 * 16)
	textColour := impeller.Color{Red: 0.0, Green: 0.0, Blue: 0.0, Alpha: 1.0}
	paint.SetColor(&textColour)
	paraStyle.SetForeground(paint)
	paraBuilder := typographyContext.ParagraphBuilderNew()
	paraBuilder.PushStyle(paraStyle)
	paraBuilder.AddText("The quick dog jumped over the lazy dog's hind legs.")
	windowWidth, _ := window.GetFramebufferSize()
	para := paraBuilder.BuildParagraphNew(float32(windowWidth))
	builder.DrawParagraph(para, &impeller.Point{X: 10, Y: 150})

	paraStyle2 := impeller.ParagraphStyleNew()
	paraStyle2.SetFontFamily(dejaVuSans)
	paraStyle2.SetFontSize(1.25 * 16)
	paint.SetColor(&textColour)
	paraStyle2.SetForeground(paint)
	paraBuilder2 := typographyContext.ParagraphBuilderNew()
	paraBuilder2.PushStyle(paraStyle2)
	paraBuilder2.AddText("The quick dog jumped over the lazy dog's hind legs.")
	para2 := paraBuilder2.BuildParagraphNew(float32(windowWidth))
	builder.DrawParagraph(para2, &impeller.Point{X: 10, Y: 200})

	dl := builder.CreateDisplayListNew()

	paint.Release()
	builder.Release()

	// window.SetCursorPosCallback(glfw.CursorposCallbackNew(func(_window *glfw.Window, xpos, ypos glfw.Double) {
	// 	fmt.Println("cursor pos", xpos, ypos)
	// }))

	window.SetFramebufferSizeCallback(glfw.FramebuffersizeCallbackNew(func(_window *glfw.Window, width, height glfw.Int) {
		// fmt.Println("frame size", width, height)
		surface.Release()
		surfaceSize := impeller.ISize{
			Width:  int64(width),
			Height: int64(height),
		}
		surface = context.SurfaceCreateWrappedFBONew(0, impeller.PixelFormatRGBA8888, &surfaceSize)
	}))

	window.Show()

	for window.ShouldClose() == glfw.FALSE {
		glfw.WaitEvents()
		surface.DrawDisplayList(dl)
		window.SwapBuffers()
	}

	dl.Release()
	surface.Release()
	context.Release()
	window.Destroy()

	glfw.Terminate()
}
