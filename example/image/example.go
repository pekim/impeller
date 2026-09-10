package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	_ "image/png" // register png decoder
	"unsafe"

	"github.com/pekim/glfw"
	"github.com/pekim/impeller"
	"github.com/pekim/impeller/example/common"
)

//go:embed primate.png
var primatePNG []byte

func main() {
	img, _, err := image.Decode(bytes.NewReader(primatePNG))
	if err != nil {
		panic(fmt.Errorf("failed to decode image: %w", err))
	}
	var pix []uint8
	switch rgba := img.(type) {
	case *image.NRGBA:
		pix = rgba.Pix
	case *image.RGBA:
		pix = rgba.Pix
	default:
		panic(fmt.Errorf("unsupported image.Image type %T", img))
	}

	var imageTexture impeller.Texture
	var paint impeller.Paint

	setup := func(_window *glfw.Window, context impeller.Context, _draw common.RequestDraw) {
		paint = impeller.PaintNew()

		imageTexture = context.TextureCreateWithContentsNew(
			&impeller.TextureDescriptor{
				PixelFormat: impeller.PixelFormatRGBA8888,
				Size: impeller.ISize{
					Width:  int64(img.Bounds().Dx()),
					Height: int64(img.Bounds().Dy()),
				},
				MipCount: 1,
			},
			&impeller.Mapping{
				Data:   unsafe.Pointer(&pix[0]),
				Length: uint64(len(pix)),
			},
			nil,
		)
	}

	teardown := func() {
		paint.Release()
	}

	draw := func(builder impeller.DisplayListBuilder, _width int32, _height int32) {
		// Clear the background to a white color.
		clearColour := impeller.Color{Red: 1.0, Green: 1.0, Blue: 1.0, Alpha: 1.0}
		paint.SetColor(&clearColour)
		builder.DrawPaint(paint)

		// Draw the image texture.
		padding := float32(50)
		builder.DrawTexture(imageTexture, &impeller.Point{X: padding, Y: padding}, impeller.TextureSamplingLinear, paint)
		// Draw the image texture, scaled to 50%.
		builder.DrawTextureRect(
			imageTexture,
			&impeller.Rect{
				X:      0,
				Y:      0,
				Width:  float32(img.Bounds().Dx()),
				Height: float32(img.Bounds().Dy()),
			},
			&impeller.Rect{
				X:      padding + float32(img.Bounds().Dx()) + padding,
				Y:      50,
				Width:  float32(img.Bounds().Dx() / 2),
				Height: float32(img.Bounds().Dy() / 2),
			},
			impeller.TextureSamplingLinear,
			paint,
		)
	}

	common.Run("simple", setup, teardown, draw)
}

// package main

// import (
// 	"bytes"
// 	_ "embed"
// 	"fmt"
// 	"image"
// 	_ "image/png" // register png decoder
// 	"runtime"
// 	"unsafe"

// 	"github.com/pekim/glfw"
// 	"github.com/pekim/impeller"
// )

// //go:embed primate.png
// var primatePNG []byte

// func init() {
// 	// This is important to ensure that only a single thread makes calls to
// 	// openGL apis.
// 	runtime.LockOSThread()
// }

// func main() {
// 	err := glfw.Initialise()
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = impeller.Init()
// 	if err != nil {
// 		panic(err)
// 	}

// 	glfw.Init()

// 	glfw.SetErrorCallback(glfw.ErrorCallbackNew(
// 		func(error_code glfw.Int, description string) {
// 			fmt.Printf("GLFW Error (%d): %s\n", error_code, description)
// 		}))

// 	glfw.WindowHint(glfw.CONTEXT_CREATION_API, glfw.EGL_CONTEXT_API)
// 	window := glfw.CreateWindow(800, 600, "Impeller Example - image", nil, nil)

// 	window.MakeContextCurrent()

// 	context := impeller.ContextCreateOpenGLESNew(
// 		impeller.GetVersion(),
// 		impeller.NewProcAddressCallback(func(procName string, _user_data unsafe.Pointer) unsafe.Pointer {
// 			return unsafe.Pointer(glfw.GetProcAddress(procName))
// 		}),
// 		nil,
// 	)

// 	framebufferWidth, framebufferHeight := window.GetFramebufferSize()
// 	surfaceSize := impeller.ISize{
// 		Width:  int64(framebufferWidth),
// 		Height: int64(framebufferHeight),
// 	}
// 	surface := context.SurfaceCreateWrappedFBONew(0, impeller.PixelFormatRGBA8888, &surfaceSize)

// 	builder := impeller.DisplayListBuilderNew(nil)

// 	img, _, err := image.Decode(bytes.NewReader(primatePNG))
// 	if err != nil {
// 		panic(fmt.Errorf("failed to decode image: %w", err))
// 	}
// 	var pix []uint8
// 	// var stride int
// 	switch rgba := img.(type) {
// 	case *image.NRGBA:
// 		pix = rgba.Pix
// 		// stride = rgba.Stride
// 	case *image.RGBA:
// 		pix = rgba.Pix
// 		// stride = rgba.Stride
// 	default:
// 		panic(fmt.Errorf("unsupported image.Image type %T", img))
// 	}
// 	imageTexture := context.TextureCreateWithContentsNew(
// 		&impeller.TextureDescriptor{
// 			PixelFormat: impeller.PixelFormatRGBA8888,
// 			Size: impeller.ISize{
// 				Width:  int64(img.Bounds().Dx()),
// 				Height: int64(img.Bounds().Dy()),
// 			},
// 			MipCount: 1,
// 		},
// 		&impeller.Mapping{
// 			Data:   unsafe.Pointer(&pix[0]),
// 			Length: uint64(len(pix)),
// 		},
// 		nil,
// 	)

// 	// Clear the background to a white color.
// 	paint := impeller.PaintNew()
// 	clearColour := impeller.Color{Red: 1.0, Green: 1.0, Blue: 1.0, Alpha: 1.0}
// 	paint.SetColor(&clearColour)
// 	builder.DrawPaint(paint)

// 	// Draw the image texture.
// 	padding := float32(50)
// 	builder.DrawTexture(imageTexture, &impeller.Point{X: padding, Y: padding}, impeller.TextureSamplingLinear, paint)
// 	// Draw the image texture, scaled to 50%.
// 	builder.DrawTextureRect(
// 		imageTexture,
// 		&impeller.Rect{
// 			X:      0,
// 			Y:      0,
// 			Width:  float32(img.Bounds().Dx()),
// 			Height: float32(img.Bounds().Dy()),
// 		},
// 		&impeller.Rect{
// 			X:      padding + float32(img.Bounds().Dx()) + padding,
// 			Y:      50,
// 			Width:  float32(img.Bounds().Dx() / 2),
// 			Height: float32(img.Bounds().Dy() / 2),
// 		},
// 		impeller.TextureSamplingLinear,
// 		paint,
// 	)

// 	dl := builder.CreateDisplayListNew()

// 	paint.Release()
// 	builder.Release()

// 	draw := true

// 	window.SetFramebufferSizeCallback(glfw.FramebuffersizeCallbackNew(func(_window *glfw.Window, width, height glfw.Int) {
// 		surface.Release()
// 		surfaceSize := impeller.ISize{
// 			Width:  int64(width),
// 			Height: int64(height),
// 		}
// 		surface = context.SurfaceCreateWrappedFBONew(0, impeller.PixelFormatRGBA8888, &surfaceSize)
// 		draw = true
// 	}))

// 	window.Show()

// 	for window.ShouldClose() == glfw.FALSE {
// 		if draw {
// 			surface.DrawDisplayList(dl)
// 			window.SwapBuffers()
// 			draw = false
// 		}

// 		glfw.WaitEvents()
// 	}

// 	surface.Release()
// 	context.Release()
// 	window.Destroy()

// 	glfw.Terminate()
// }
