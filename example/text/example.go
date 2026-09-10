package main

import (
	"github.com/pekim/glfw"
	"github.com/pekim/impeller"
	"github.com/pekim/impeller/example/common"
)

func main() {
	var typographyContext impeller.TypographyContext
	padding := float32(20)

	black := impeller.Color{Red: 0.0, Green: 0.0, Blue: 0.0, Alpha: 1.0}
	red := impeller.Color{Red: 7.0, Green: 0.0, Blue: 0.0, Alpha: 1.0}
	white := impeller.Color{Red: 1.0, Green: 1.0, Blue: 1.0, Alpha: 1.0}

	var blackPaint impeller.Paint
	var redPaint impeller.Paint
	var whitePaint impeller.Paint

	setup := func(_window *glfw.Window, _context impeller.Context, _draw common.RequestDraw) {
		typographyContext = impeller.TypographyContextNew()

		blackPaint = impeller.PaintNew()
		blackPaint.SetColor(&black)
		redPaint = impeller.PaintNew()
		redPaint.SetColor(&red)
		whitePaint = impeller.PaintNew()
		whitePaint.SetColor(&white)
	}
	teardown := func() {
		blackPaint.Release()
		redPaint.Release()
		whitePaint.Release()
	}

	draw := func(builder impeller.DisplayListBuilder, windowWidth int32, _windowHeight int32) {
		// Clear the background to a white color.
		builder.DrawPaint(whitePaint)

		// Draw some text.

		style := impeller.ParagraphStyleNew()
		style.SetFontSize(24)
		style.SetForeground(blackPaint)

		alternateStyle := impeller.ParagraphStyleNew()
		alternateStyle.SetFontSize(24)
		alternateStyle.SetFontStyle(impeller.FontStyleItalic)
		alternateStyle.SetTextDecoration(&impeller.TextDecoration{
			Types:               int32(impeller.TextDecorationTypeUnderline),
			Color:               black,
			Style:               impeller.TextDecorationStyleSolid,
			ThicknessMultiplier: 1,
		})
		alternateStyle.SetForeground(redPaint)

		paraBuilder := typographyContext.ParagraphBuilderNew()
		paraBuilder.PushStyle(style)
		paraBuilder.AddText("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute")
		paraBuilder.PopStyle()

		paraBuilder.PushStyle(alternateStyle)
		paraBuilder.AddText(" irure dolor")
		paraBuilder.PopStyle()

		paraBuilder.PushStyle(style)
		paraBuilder.AddText(" in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")
		paraBuilder.PopStyle()

		para := paraBuilder.BuildParagraphNew(float32(windowWidth) - (2 * padding))
		builder.DrawParagraph(para, &impeller.Point{X: padding, Y: padding})
	}

	common.Run("simple", setup, teardown, draw)
}
