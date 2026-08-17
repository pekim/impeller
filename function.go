// This is a generated file. DO NOT EDIT.

package impeller

import (
	"unsafe"

	ffi "github.com/go-webgpu/goffi/ffi"
)

/*
Get the version of Impeller standalone API. This is the API that
will be accepted for validity checks when provided to the
context creation methods.

The current version of the API  is denoted by the
`IMPELLER_VERSION` macro. This version must be passed to APIs
that create top-level objects like graphics contexts.
Construction of the context may fail if the API version expected
by the caller is not supported by the library.

Since there are no API stability guarantees today, passing a
version that is different to the one returned by
`ImpellerGetVersion` will always fail.

@see        `ImpellerContextCreateOpenGLESNew`

@return     The version of the standalone API.
*/
func GetVersion() uint32 {
	var result uint32
	_, err := ffi.CallFunction(
		cifImpellerGetVersion,
		funcImpellerGetVersion,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{},
	)
	if err != nil {
		panic(err)
	}
	return result
}

// UNSUPPORTED :: ContextCreateOpenGLESNew  param count = 3

// UNSUPPORTED :: ContextCreateMetalNew  param count = 1

// UNSUPPORTED :: ContextCreateVulkanNew  param count = 2

// UNSUPPORTED :: ContextRetain  param count = 1

// UNSUPPORTED :: ContextRelease  param count = 1

// UNSUPPORTED :: ContextGetVulkanInfo  param count = 2

// UNSUPPORTED :: VulkanSwapchainCreateNew  param count = 2

// UNSUPPORTED :: VulkanSwapchainRetain  param count = 1

// UNSUPPORTED :: VulkanSwapchainRelease  param count = 1

// UNSUPPORTED :: VulkanSwapchainAcquireNextSurfaceNew  param count = 1

// UNSUPPORTED :: SurfaceCreateWrappedFBONew  param count = 4

// UNSUPPORTED :: SurfaceCreateWrappedMetalDrawableNew  param count = 2

// UNSUPPORTED :: SurfaceRetain  param count = 1

// UNSUPPORTED :: SurfaceRelease  param count = 1

// UNSUPPORTED :: SurfaceDrawDisplayList  param count = 2

// UNSUPPORTED :: SurfacePresent  param count = 1

// UNSUPPORTED :: PathRetain  param count = 1

// UNSUPPORTED :: PathRelease  param count = 1

// UNSUPPORTED :: PathGetBounds  param count = 2

/*
Create a new path builder. Paths themselves are immutable.
A builder builds these immutable paths.

@return     The path builder.
*/
func PathBuilderNew() PathBuilder {
	var result PathBuilder
	_, err := ffi.CallFunction(
		cifImpellerPathBuilderNew,
		funcImpellerPathBuilderNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{},
	)
	if err != nil {
		panic(err)
	}
	return result
}

// UNSUPPORTED :: PathBuilderRetain  param count = 1

// UNSUPPORTED :: PathBuilderRelease  param count = 1

// UNSUPPORTED :: PathBuilderMoveTo  param count = 2

// UNSUPPORTED :: PathBuilderLineTo  param count = 2

// UNSUPPORTED :: PathBuilderQuadraticCurveTo  param count = 3

// UNSUPPORTED :: PathBuilderCubicCurveTo  param count = 4

// UNSUPPORTED :: PathBuilderAddRect  param count = 2

// UNSUPPORTED :: PathBuilderAddArc  param count = 4

// UNSUPPORTED :: PathBuilderAddOval  param count = 2

// UNSUPPORTED :: PathBuilderAddRoundedRect  param count = 3

// UNSUPPORTED :: PathBuilderClose  param count = 1

// UNSUPPORTED :: PathBuilderCopyPathNew  param count = 2

// UNSUPPORTED :: PathBuilderTakePathNew  param count = 2

/*
Create a new paint with default values.

@return     The impeller paint.
*/
func PaintNew() Paint {
	var result Paint
	_, err := ffi.CallFunction(
		cifImpellerPaintNew,
		funcImpellerPaintNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{},
	)
	if err != nil {
		panic(err)
	}
	return result
}

// UNSUPPORTED :: PaintRetain  param count = 1

// UNSUPPORTED :: PaintRelease  param count = 1

// UNSUPPORTED :: PaintSetColor  param count = 2

// UNSUPPORTED :: PaintSetBlendMode  param count = 2

// UNSUPPORTED :: PaintSetDrawStyle  param count = 2

// UNSUPPORTED :: PaintSetStrokeCap  param count = 2

// UNSUPPORTED :: PaintSetStrokeJoin  param count = 2

// UNSUPPORTED :: PaintSetStrokeWidth  param count = 2

// UNSUPPORTED :: PaintSetStrokeMiter  param count = 2

// UNSUPPORTED :: PaintSetColorFilter  param count = 2

// UNSUPPORTED :: PaintSetColorSource  param count = 2

// UNSUPPORTED :: PaintSetImageFilter  param count = 2

// UNSUPPORTED :: PaintSetMaskFilter  param count = 2

// UNSUPPORTED :: TextureCreateWithContentsNew  param count = 4

// UNSUPPORTED :: TextureCreateWithOpenGLTextureHandleNew  param count = 3

// UNSUPPORTED :: TextureRetain  param count = 1

// UNSUPPORTED :: TextureRelease  param count = 1

// UNSUPPORTED :: TextureGetOpenGLHandle  param count = 1

// UNSUPPORTED :: FragmentProgramNew  param count = 2

// UNSUPPORTED :: FragmentProgramRetain  param count = 1

// UNSUPPORTED :: FragmentProgramRelease  param count = 1

// UNSUPPORTED :: ColorSourceRetain  param count = 1

// UNSUPPORTED :: ColorSourceRelease  param count = 1

// UNSUPPORTED :: ColorSourceCreateLinearGradientNew  param count = 7

// UNSUPPORTED :: ColorSourceCreateRadialGradientNew  param count = 7

// UNSUPPORTED :: ColorSourceCreateConicalGradientNew  param count = 9

// UNSUPPORTED :: ColorSourceCreateSweepGradientNew  param count = 8

// UNSUPPORTED :: ColorSourceCreateImageNew  param count = 5

// UNSUPPORTED :: ColorSourceCreateFragmentProgramNew  param count = 6

// UNSUPPORTED :: ColorFilterRetain  param count = 1

// UNSUPPORTED :: ColorFilterRelease  param count = 1

// UNSUPPORTED :: ColorFilterCreateBlendNew  param count = 2

// UNSUPPORTED :: ColorFilterCreateColorMatrixNew  param count = 1

// UNSUPPORTED :: MaskFilterRetain  param count = 1

// UNSUPPORTED :: MaskFilterRelease  param count = 1

// UNSUPPORTED :: MaskFilterCreateBlurNew  param count = 2

// UNSUPPORTED :: ImageFilterRetain  param count = 1

// UNSUPPORTED :: ImageFilterRelease  param count = 1

// UNSUPPORTED :: ImageFilterCreateBlurNew  param count = 3

// UNSUPPORTED :: ImageFilterCreateDilateNew  param count = 2

// UNSUPPORTED :: ImageFilterCreateErodeNew  param count = 2

// UNSUPPORTED :: ImageFilterCreateMatrixNew  param count = 2

// UNSUPPORTED :: ImageFilterCreateFragmentProgramNew  param count = 6

// UNSUPPORTED :: ImageFilterCreateComposeNew  param count = 2

// UNSUPPORTED :: DisplayListRetain  param count = 1

// UNSUPPORTED :: DisplayListRelease  param count = 1

// UNSUPPORTED :: DisplayListBuilderNew  param count = 1

// UNSUPPORTED :: DisplayListBuilderRetain  param count = 1

// UNSUPPORTED :: DisplayListBuilderRelease  param count = 1

// UNSUPPORTED :: DisplayListBuilderCreateDisplayListNew  param count = 1

// UNSUPPORTED :: DisplayListBuilderSave  param count = 1

// UNSUPPORTED :: DisplayListBuilderSaveLayer  param count = 4

// UNSUPPORTED :: DisplayListBuilderRestore  param count = 1

// UNSUPPORTED :: DisplayListBuilderScale  param count = 3

// UNSUPPORTED :: DisplayListBuilderRotate  param count = 2

// UNSUPPORTED :: DisplayListBuilderTranslate  param count = 3

// UNSUPPORTED :: DisplayListBuilderTransform  param count = 2

// UNSUPPORTED :: DisplayListBuilderSetTransform  param count = 2

// UNSUPPORTED :: DisplayListBuilderGetTransform  param count = 2

// UNSUPPORTED :: DisplayListBuilderResetTransform  param count = 1

// UNSUPPORTED :: DisplayListBuilderGetSaveCount  param count = 1

// UNSUPPORTED :: DisplayListBuilderRestoreToCount  param count = 2

// UNSUPPORTED :: DisplayListBuilderClipRect  param count = 3

// UNSUPPORTED :: DisplayListBuilderClipOval  param count = 3

// UNSUPPORTED :: DisplayListBuilderClipRoundedRect  param count = 4

// UNSUPPORTED :: DisplayListBuilderClipPath  param count = 3

// UNSUPPORTED :: DisplayListBuilderDrawPaint  param count = 2

// UNSUPPORTED :: DisplayListBuilderDrawLine  param count = 4

// UNSUPPORTED :: DisplayListBuilderDrawDashedLine  param count = 6

// UNSUPPORTED :: DisplayListBuilderDrawRect  param count = 3

// UNSUPPORTED :: DisplayListBuilderDrawOval  param count = 3

// UNSUPPORTED :: DisplayListBuilderDrawRoundedRect  param count = 4

// UNSUPPORTED :: DisplayListBuilderDrawRoundedRectDifference  param count = 6

// UNSUPPORTED :: DisplayListBuilderDrawPath  param count = 3

// UNSUPPORTED :: DisplayListBuilderDrawDisplayList  param count = 3

// UNSUPPORTED :: DisplayListBuilderDrawParagraph  param count = 3

// UNSUPPORTED :: DisplayListBuilderDrawShadow  param count = 6

// UNSUPPORTED :: DisplayListBuilderDrawTexture  param count = 5

// UNSUPPORTED :: DisplayListBuilderDrawTextureRect  param count = 6

/*
Create a new typography contents.

@return     The typography context.
*/
func TypographyContextNew() TypographyContext {
	var result TypographyContext
	_, err := ffi.CallFunction(
		cifImpellerTypographyContextNew,
		funcImpellerTypographyContextNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{},
	)
	if err != nil {
		panic(err)
	}
	return result
}

// UNSUPPORTED :: TypographyContextRetain  param count = 1

// UNSUPPORTED :: TypographyContextRelease  param count = 1

// UNSUPPORTED :: TypographyContextRegisterFont  param count = 4

/*
Create a new paragraph style.

@return     The paragraph style.
*/
func ParagraphStyleNew() ParagraphStyle {
	var result ParagraphStyle
	_, err := ffi.CallFunction(
		cifImpellerParagraphStyleNew,
		funcImpellerParagraphStyleNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{},
	)
	if err != nil {
		panic(err)
	}
	return result
}

// UNSUPPORTED :: ParagraphStyleRetain  param count = 1

// UNSUPPORTED :: ParagraphStyleRelease  param count = 1

// UNSUPPORTED :: ParagraphStyleSetForeground  param count = 2

// UNSUPPORTED :: ParagraphStyleSetBackground  param count = 2

// UNSUPPORTED :: ParagraphStyleSetFontWeight  param count = 2

// UNSUPPORTED :: ParagraphStyleSetFontStyle  param count = 2

// UNSUPPORTED :: ParagraphStyleSetFontFamily  param count = 2

// UNSUPPORTED :: ParagraphStyleSetFontSize  param count = 2

// UNSUPPORTED :: ParagraphStyleSetHeight  param count = 2

// UNSUPPORTED :: ParagraphStyleSetTextAlignment  param count = 2

// UNSUPPORTED :: ParagraphStyleSetTextDirection  param count = 2

// UNSUPPORTED :: ParagraphStyleSetTextDecoration  param count = 2

// UNSUPPORTED :: ParagraphStyleSetMaxLines  param count = 2

// UNSUPPORTED :: ParagraphStyleSetLocale  param count = 2

// UNSUPPORTED :: ParagraphStyleSetEllipsis  param count = 2

// UNSUPPORTED :: ParagraphBuilderNew  param count = 1

// UNSUPPORTED :: ParagraphBuilderRetain  param count = 1

// UNSUPPORTED :: ParagraphBuilderRelease  param count = 1

// UNSUPPORTED :: ParagraphBuilderPushStyle  param count = 2

// UNSUPPORTED :: ParagraphBuilderPopStyle  param count = 1

// UNSUPPORTED :: ParagraphBuilderAddText  param count = 3

// UNSUPPORTED :: ParagraphBuilderBuildParagraphNew  param count = 2

// UNSUPPORTED :: ParagraphRetain  param count = 1

// UNSUPPORTED :: ParagraphRelease  param count = 1

// UNSUPPORTED :: ParagraphGetMaxWidth  param count = 1

// UNSUPPORTED :: ParagraphGetHeight  param count = 1

// UNSUPPORTED :: ParagraphGetLongestLineWidth  param count = 1

// UNSUPPORTED :: ParagraphGetMinIntrinsicWidth  param count = 1

// UNSUPPORTED :: ParagraphGetMaxIntrinsicWidth  param count = 1

// UNSUPPORTED :: ParagraphGetIdeographicBaseline  param count = 1

// UNSUPPORTED :: ParagraphGetAlphabeticBaseline  param count = 1

// UNSUPPORTED :: ParagraphGetLineCount  param count = 1

// UNSUPPORTED :: ParagraphGetWordBoundary  param count = 3

// UNSUPPORTED :: ParagraphGetLineMetrics  param count = 1

// UNSUPPORTED :: ParagraphCreateGlyphInfoAtCodeUnitIndexNew  param count = 2

// UNSUPPORTED :: ParagraphCreateGlyphInfoAtParagraphCoordinatesNew  param count = 3

// UNSUPPORTED :: LineMetricsRetain  param count = 1

// UNSUPPORTED :: LineMetricsRelease  param count = 1

// UNSUPPORTED :: LineMetricsGetUnscaledAscent  param count = 2

// UNSUPPORTED :: LineMetricsGetAscent  param count = 2

// UNSUPPORTED :: LineMetricsGetDescent  param count = 2

// UNSUPPORTED :: LineMetricsGetBaseline  param count = 2

// UNSUPPORTED :: LineMetricsIsHardbreak  param count = 2

// UNSUPPORTED :: LineMetricsGetWidth  param count = 2

// UNSUPPORTED :: LineMetricsGetHeight  param count = 2

// UNSUPPORTED :: LineMetricsGetLeft  param count = 2

// UNSUPPORTED :: LineMetricsGetCodeUnitStartIndex  param count = 2

// UNSUPPORTED :: LineMetricsGetCodeUnitEndIndex  param count = 2

// UNSUPPORTED :: LineMetricsGetCodeUnitEndIndexExcludingWhitespace  param count = 2

// UNSUPPORTED :: LineMetricsGetCodeUnitEndIndexIncludingNewline  param count = 2

// UNSUPPORTED :: GlyphInfoRetain  param count = 1

// UNSUPPORTED :: GlyphInfoRelease  param count = 1

// UNSUPPORTED :: GlyphInfoGetGraphemeClusterCodeUnitRangeBegin  param count = 1

// UNSUPPORTED :: GlyphInfoGetGraphemeClusterCodeUnitRangeEnd  param count = 1

// UNSUPPORTED :: GlyphInfoGetGraphemeClusterBounds  param count = 2

// UNSUPPORTED :: GlyphInfoIsEllipsis  param count = 1

// UNSUPPORTED :: GlyphInfoGetTextDirection  param count = 1
