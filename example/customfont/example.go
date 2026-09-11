package main

import (
	_ "embed"

	"github.com/pekim/glfw"
	"github.com/pekim/impeller"
	"github.com/pekim/impeller/example/common"
)

//go:embed DejaVuSans.ttf
var dejaVuSansTTF []byte

//go:embed NotoColorEmoji.ttf
var notoColorEmojiTTF []byte

func main() {
	var typographyContext impeller.TypographyContext
	padding := float32(20)
	dejaVuSans := "DejaVuSans"
	notoColorEmoji := "NotoColorEmoji"

	black := impeller.Color{Red: 0.0, Green: 0.0, Blue: 0.0, Alpha: 1.0}
	white := impeller.Color{Red: 1.0, Green: 1.0, Blue: 1.0, Alpha: 1.0}
	var blackPaint impeller.Paint
	var whitePaint impeller.Paint

	registerFont := func(name string, data []byte) {
		fontMapping := impeller.Mapping{
			Data:   &data[0],
			Length: uint64(len(data)),
		}
		if !typographyContext.RegisterFont(&fontMapping, nil, name).Bool() {
			panic("failed to register font")
		}
	}

	setup := func(_window *glfw.Window, _context impeller.Context, _draw common.RequestDraw) {
		typographyContext = impeller.TypographyContextNew()

		registerFont(dejaVuSans, dejaVuSansTTF)
		registerFont(notoColorEmoji, notoColorEmojiTTF)

		blackPaint = impeller.PaintNew()
		blackPaint.SetColor(&black)
		whitePaint = impeller.PaintNew()
		whitePaint.SetColor(&white)
	}

	teardown := func() {
		whitePaint.Release()
	}

	draw := func(builder impeller.DisplayListBuilder, windowWidth int32, _windowHeight int32) {
		drawText := func(text string, fontFamily string, y float32) (float64, float64) {
			style := impeller.ParagraphStyleNew()
			style.SetFontSize(24)
			style.SetFontFamily(fontFamily)
			style.SetForeground(blackPaint)

			paraBuilder := typographyContext.ParagraphBuilderNew()
			paraBuilder.PushStyle(style)
			paraBuilder.AddText(text)
			para := paraBuilder.BuildParagraphNew(float32(windowWidth) - (2 * padding))
			builder.DrawParagraph(para, &impeller.Point{X: padding, Y: y})

			metrics := para.GetLineMetrics()
			return metrics.GetWidth(0), metrics.GetHeight(0)
		}

		// Clear the background to a white color.
		builder.DrawPaint(whitePaint)

		// Draw lines of text with different fonts.
		y := padding
		_, textHeight1 := drawText("This is with the default font family.", "", y)
		y += float32(textHeight1) + padding
		_, textHeight2 := drawText("This is with the "+dejaVuSans+" font family", dejaVuSans, y)
		y += float32(textHeight2) + padding
		drawText("😀 🤠 💯", notoColorEmoji, y)
	}

	common.Run("custom font", setup, teardown, draw)
}
