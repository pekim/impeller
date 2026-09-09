package main

import (
	"github.com/pekim/glfw"
	"github.com/pekim/impeller"
	"github.com/pekim/impeller/example/common"
)

func main() {
	var paint impeller.Paint

	setup := func(_window *glfw.Window, _draw common.RequestDraw) {
		paint = impeller.PaintNew()
	}
	teardown := func() {
		paint.Release()
	}

	draw := func(builder impeller.DisplayListBuilder, width int32, height int32) {
		// Clear the background to a white color.
		clearColour := impeller.Color{Red: 1.0, Green: 1.0, Blue: 1.0, Alpha: 1.0}
		paint.SetColor(&clearColour)
		builder.DrawPaint(paint)

		// Draw a red box, centred in the window.
		boxColour := impeller.Color{Red: 1.0, Green: 0.0, Blue: 0.0, Alpha: 1.0}
		paint.SetColor(&boxColour)
		boxSize := float32(100)
		boxRect := impeller.Rect{
			X:      float32(width/2) - boxSize/2,
			Y:      float32(height/2) - boxSize/2,
			Width:  boxSize,
			Height: boxSize,
		}
		builder.DrawRect(&boxRect, paint)
	}

	common.Run("simple", setup, teardown, draw)
}
