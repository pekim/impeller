// This is a generated file. DO NOT EDIT.

package impeller

import (
	"unsafe"

	ffi "github.com/go-webgpu/goffi/ffi"
	types "github.com/go-webgpu/goffi/types"
)

var ImpellerGetVersion unsafe.Pointer
var ImpellerGetVersion_cif = &types.CallInterface{}
var ImpellerContextCreateOpenGLESNew unsafe.Pointer
var ImpellerContextCreateOpenGLESNew_cif = &types.CallInterface{}
var ImpellerContextCreateMetalNew unsafe.Pointer
var ImpellerContextCreateMetalNew_cif = &types.CallInterface{}
var ImpellerContextCreateVulkanNew unsafe.Pointer
var ImpellerContextCreateVulkanNew_cif = &types.CallInterface{}
var ImpellerContextRetain unsafe.Pointer
var ImpellerContextRetain_cif = &types.CallInterface{}
var ImpellerContextRelease unsafe.Pointer
var ImpellerContextRelease_cif = &types.CallInterface{}
var ImpellerContextGetVulkanInfo unsafe.Pointer
var ImpellerContextGetVulkanInfo_cif = &types.CallInterface{}
var ImpellerVulkanSwapchainCreateNew unsafe.Pointer
var ImpellerVulkanSwapchainCreateNew_cif = &types.CallInterface{}
var ImpellerVulkanSwapchainRetain unsafe.Pointer
var ImpellerVulkanSwapchainRetain_cif = &types.CallInterface{}
var ImpellerVulkanSwapchainRelease unsafe.Pointer
var ImpellerVulkanSwapchainRelease_cif = &types.CallInterface{}
var ImpellerVulkanSwapchainAcquireNextSurfaceNew unsafe.Pointer
var ImpellerVulkanSwapchainAcquireNextSurfaceNew_cif = &types.CallInterface{}
var ImpellerSurfaceCreateWrappedFBONew unsafe.Pointer
var ImpellerSurfaceCreateWrappedFBONew_cif = &types.CallInterface{}
var ImpellerSurfaceCreateWrappedMetalDrawableNew unsafe.Pointer
var ImpellerSurfaceCreateWrappedMetalDrawableNew_cif = &types.CallInterface{}
var ImpellerSurfaceRetain unsafe.Pointer
var ImpellerSurfaceRetain_cif = &types.CallInterface{}
var ImpellerSurfaceRelease unsafe.Pointer
var ImpellerSurfaceRelease_cif = &types.CallInterface{}
var ImpellerSurfaceDrawDisplayList unsafe.Pointer
var ImpellerSurfaceDrawDisplayList_cif = &types.CallInterface{}
var ImpellerSurfacePresent unsafe.Pointer
var ImpellerSurfacePresent_cif = &types.CallInterface{}
var ImpellerPathRetain unsafe.Pointer
var ImpellerPathRetain_cif = &types.CallInterface{}
var ImpellerPathRelease unsafe.Pointer
var ImpellerPathRelease_cif = &types.CallInterface{}
var ImpellerPathGetBounds unsafe.Pointer
var ImpellerPathGetBounds_cif = &types.CallInterface{}
var ImpellerPathBuilderNew unsafe.Pointer
var ImpellerPathBuilderNew_cif = &types.CallInterface{}
var ImpellerPathBuilderRetain unsafe.Pointer
var ImpellerPathBuilderRetain_cif = &types.CallInterface{}
var ImpellerPathBuilderRelease unsafe.Pointer
var ImpellerPathBuilderRelease_cif = &types.CallInterface{}
var ImpellerPathBuilderMoveTo unsafe.Pointer
var ImpellerPathBuilderMoveTo_cif = &types.CallInterface{}
var ImpellerPathBuilderLineTo unsafe.Pointer
var ImpellerPathBuilderLineTo_cif = &types.CallInterface{}
var ImpellerPathBuilderQuadraticCurveTo unsafe.Pointer
var ImpellerPathBuilderQuadraticCurveTo_cif = &types.CallInterface{}
var ImpellerPathBuilderCubicCurveTo unsafe.Pointer
var ImpellerPathBuilderCubicCurveTo_cif = &types.CallInterface{}
var ImpellerPathBuilderAddRect unsafe.Pointer
var ImpellerPathBuilderAddRect_cif = &types.CallInterface{}
var ImpellerPathBuilderAddArc unsafe.Pointer
var ImpellerPathBuilderAddArc_cif = &types.CallInterface{}
var ImpellerPathBuilderAddOval unsafe.Pointer
var ImpellerPathBuilderAddOval_cif = &types.CallInterface{}
var ImpellerPathBuilderAddRoundedRect unsafe.Pointer
var ImpellerPathBuilderAddRoundedRect_cif = &types.CallInterface{}
var ImpellerPathBuilderClose unsafe.Pointer
var ImpellerPathBuilderClose_cif = &types.CallInterface{}
var ImpellerPathBuilderCopyPathNew unsafe.Pointer
var ImpellerPathBuilderCopyPathNew_cif = &types.CallInterface{}
var ImpellerPathBuilderTakePathNew unsafe.Pointer
var ImpellerPathBuilderTakePathNew_cif = &types.CallInterface{}
var ImpellerPaintNew unsafe.Pointer
var ImpellerPaintNew_cif = &types.CallInterface{}
var ImpellerPaintRetain unsafe.Pointer
var ImpellerPaintRetain_cif = &types.CallInterface{}
var ImpellerPaintRelease unsafe.Pointer
var ImpellerPaintRelease_cif = &types.CallInterface{}
var ImpellerPaintSetColor unsafe.Pointer
var ImpellerPaintSetColor_cif = &types.CallInterface{}
var ImpellerPaintSetBlendMode unsafe.Pointer
var ImpellerPaintSetBlendMode_cif = &types.CallInterface{}
var ImpellerPaintSetDrawStyle unsafe.Pointer
var ImpellerPaintSetDrawStyle_cif = &types.CallInterface{}
var ImpellerPaintSetStrokeCap unsafe.Pointer
var ImpellerPaintSetStrokeCap_cif = &types.CallInterface{}
var ImpellerPaintSetStrokeJoin unsafe.Pointer
var ImpellerPaintSetStrokeJoin_cif = &types.CallInterface{}
var ImpellerPaintSetStrokeWidth unsafe.Pointer
var ImpellerPaintSetStrokeWidth_cif = &types.CallInterface{}
var ImpellerPaintSetStrokeMiter unsafe.Pointer
var ImpellerPaintSetStrokeMiter_cif = &types.CallInterface{}
var ImpellerPaintSetColorFilter unsafe.Pointer
var ImpellerPaintSetColorFilter_cif = &types.CallInterface{}
var ImpellerPaintSetColorSource unsafe.Pointer
var ImpellerPaintSetColorSource_cif = &types.CallInterface{}
var ImpellerPaintSetImageFilter unsafe.Pointer
var ImpellerPaintSetImageFilter_cif = &types.CallInterface{}
var ImpellerPaintSetMaskFilter unsafe.Pointer
var ImpellerPaintSetMaskFilter_cif = &types.CallInterface{}
var ImpellerTextureCreateWithContentsNew unsafe.Pointer
var ImpellerTextureCreateWithContentsNew_cif = &types.CallInterface{}
var ImpellerTextureCreateWithOpenGLTextureHandleNew unsafe.Pointer
var ImpellerTextureCreateWithOpenGLTextureHandleNew_cif = &types.CallInterface{}
var ImpellerTextureRetain unsafe.Pointer
var ImpellerTextureRetain_cif = &types.CallInterface{}
var ImpellerTextureRelease unsafe.Pointer
var ImpellerTextureRelease_cif = &types.CallInterface{}
var ImpellerTextureGetOpenGLHandle unsafe.Pointer
var ImpellerTextureGetOpenGLHandle_cif = &types.CallInterface{}
var ImpellerFragmentProgramNew unsafe.Pointer
var ImpellerFragmentProgramNew_cif = &types.CallInterface{}
var ImpellerFragmentProgramRetain unsafe.Pointer
var ImpellerFragmentProgramRetain_cif = &types.CallInterface{}
var ImpellerFragmentProgramRelease unsafe.Pointer
var ImpellerFragmentProgramRelease_cif = &types.CallInterface{}
var ImpellerColorSourceRetain unsafe.Pointer
var ImpellerColorSourceRetain_cif = &types.CallInterface{}
var ImpellerColorSourceRelease unsafe.Pointer
var ImpellerColorSourceRelease_cif = &types.CallInterface{}
var ImpellerColorSourceCreateLinearGradientNew unsafe.Pointer
var ImpellerColorSourceCreateLinearGradientNew_cif = &types.CallInterface{}
var ImpellerColorSourceCreateRadialGradientNew unsafe.Pointer
var ImpellerColorSourceCreateRadialGradientNew_cif = &types.CallInterface{}
var ImpellerColorSourceCreateConicalGradientNew unsafe.Pointer
var ImpellerColorSourceCreateConicalGradientNew_cif = &types.CallInterface{}
var ImpellerColorSourceCreateSweepGradientNew unsafe.Pointer
var ImpellerColorSourceCreateSweepGradientNew_cif = &types.CallInterface{}
var ImpellerColorSourceCreateImageNew unsafe.Pointer
var ImpellerColorSourceCreateImageNew_cif = &types.CallInterface{}
var ImpellerColorSourceCreateFragmentProgramNew unsafe.Pointer
var ImpellerColorSourceCreateFragmentProgramNew_cif = &types.CallInterface{}
var ImpellerColorFilterRetain unsafe.Pointer
var ImpellerColorFilterRetain_cif = &types.CallInterface{}
var ImpellerColorFilterRelease unsafe.Pointer
var ImpellerColorFilterRelease_cif = &types.CallInterface{}
var ImpellerColorFilterCreateBlendNew unsafe.Pointer
var ImpellerColorFilterCreateBlendNew_cif = &types.CallInterface{}
var ImpellerColorFilterCreateColorMatrixNew unsafe.Pointer
var ImpellerColorFilterCreateColorMatrixNew_cif = &types.CallInterface{}
var ImpellerMaskFilterRetain unsafe.Pointer
var ImpellerMaskFilterRetain_cif = &types.CallInterface{}
var ImpellerMaskFilterRelease unsafe.Pointer
var ImpellerMaskFilterRelease_cif = &types.CallInterface{}
var ImpellerMaskFilterCreateBlurNew unsafe.Pointer
var ImpellerMaskFilterCreateBlurNew_cif = &types.CallInterface{}
var ImpellerImageFilterRetain unsafe.Pointer
var ImpellerImageFilterRetain_cif = &types.CallInterface{}
var ImpellerImageFilterRelease unsafe.Pointer
var ImpellerImageFilterRelease_cif = &types.CallInterface{}
var ImpellerImageFilterCreateBlurNew unsafe.Pointer
var ImpellerImageFilterCreateBlurNew_cif = &types.CallInterface{}
var ImpellerImageFilterCreateDilateNew unsafe.Pointer
var ImpellerImageFilterCreateDilateNew_cif = &types.CallInterface{}
var ImpellerImageFilterCreateErodeNew unsafe.Pointer
var ImpellerImageFilterCreateErodeNew_cif = &types.CallInterface{}
var ImpellerImageFilterCreateMatrixNew unsafe.Pointer
var ImpellerImageFilterCreateMatrixNew_cif = &types.CallInterface{}
var ImpellerImageFilterCreateFragmentProgramNew unsafe.Pointer
var ImpellerImageFilterCreateFragmentProgramNew_cif = &types.CallInterface{}
var ImpellerImageFilterCreateComposeNew unsafe.Pointer
var ImpellerImageFilterCreateComposeNew_cif = &types.CallInterface{}
var ImpellerDisplayListRetain unsafe.Pointer
var ImpellerDisplayListRetain_cif = &types.CallInterface{}
var ImpellerDisplayListRelease unsafe.Pointer
var ImpellerDisplayListRelease_cif = &types.CallInterface{}
var ImpellerDisplayListBuilderNew unsafe.Pointer
var ImpellerDisplayListBuilderNew_cif = &types.CallInterface{}
var ImpellerDisplayListBuilderRetain unsafe.Pointer
var ImpellerDisplayListBuilderRetain_cif = &types.CallInterface{}
var ImpellerDisplayListBuilderRelease unsafe.Pointer
var ImpellerDisplayListBuilderRelease_cif = &types.CallInterface{}
var ImpellerDisplayListBuilderCreateDisplayListNew unsafe.Pointer
var ImpellerDisplayListBuilderCreateDisplayListNew_cif = &types.CallInterface{}
var ImpellerDisplayListBuilderSave unsafe.Pointer
var ImpellerDisplayListBuilderSave_cif = &types.CallInterface{}
var ImpellerDisplayListBuilderSaveLayer unsafe.Pointer
var ImpellerDisplayListBuilderSaveLayer_cif = &types.CallInterface{}
var ImpellerDisplayListBuilderRestore unsafe.Pointer
var ImpellerDisplayListBuilderRestore_cif = &types.CallInterface{}
var ImpellerDisplayListBuilderScale unsafe.Pointer
var ImpellerDisplayListBuilderScale_cif = &types.CallInterface{}
var ImpellerDisplayListBuilderRotate unsafe.Pointer
var ImpellerDisplayListBuilderRotate_cif = &types.CallInterface{}
var ImpellerDisplayListBuilderTranslate unsafe.Pointer
var ImpellerDisplayListBuilderTranslate_cif = &types.CallInterface{}
var ImpellerDisplayListBuilderTransform unsafe.Pointer
var ImpellerDisplayListBuilderTransform_cif = &types.CallInterface{}
var ImpellerDisplayListBuilderSetTransform unsafe.Pointer
var ImpellerDisplayListBuilderSetTransform_cif = &types.CallInterface{}
var ImpellerDisplayListBuilderGetTransform unsafe.Pointer
var ImpellerDisplayListBuilderGetTransform_cif = &types.CallInterface{}
var ImpellerDisplayListBuilderResetTransform unsafe.Pointer
var ImpellerDisplayListBuilderResetTransform_cif = &types.CallInterface{}
var ImpellerDisplayListBuilderGetSaveCount unsafe.Pointer
var ImpellerDisplayListBuilderGetSaveCount_cif = &types.CallInterface{}
var ImpellerDisplayListBuilderRestoreToCount unsafe.Pointer
var ImpellerDisplayListBuilderRestoreToCount_cif = &types.CallInterface{}
var ImpellerDisplayListBuilderClipRect unsafe.Pointer
var ImpellerDisplayListBuilderClipRect_cif = &types.CallInterface{}
var ImpellerDisplayListBuilderClipOval unsafe.Pointer
var ImpellerDisplayListBuilderClipOval_cif = &types.CallInterface{}
var ImpellerDisplayListBuilderClipRoundedRect unsafe.Pointer
var ImpellerDisplayListBuilderClipRoundedRect_cif = &types.CallInterface{}
var ImpellerDisplayListBuilderClipPath unsafe.Pointer
var ImpellerDisplayListBuilderClipPath_cif = &types.CallInterface{}
var ImpellerDisplayListBuilderDrawPaint unsafe.Pointer
var ImpellerDisplayListBuilderDrawPaint_cif = &types.CallInterface{}
var ImpellerDisplayListBuilderDrawLine unsafe.Pointer
var ImpellerDisplayListBuilderDrawLine_cif = &types.CallInterface{}
var ImpellerDisplayListBuilderDrawDashedLine unsafe.Pointer
var ImpellerDisplayListBuilderDrawDashedLine_cif = &types.CallInterface{}
var ImpellerDisplayListBuilderDrawRect unsafe.Pointer
var ImpellerDisplayListBuilderDrawRect_cif = &types.CallInterface{}
var ImpellerDisplayListBuilderDrawOval unsafe.Pointer
var ImpellerDisplayListBuilderDrawOval_cif = &types.CallInterface{}
var ImpellerDisplayListBuilderDrawRoundedRect unsafe.Pointer
var ImpellerDisplayListBuilderDrawRoundedRect_cif = &types.CallInterface{}
var ImpellerDisplayListBuilderDrawRoundedRectDifference unsafe.Pointer
var ImpellerDisplayListBuilderDrawRoundedRectDifference_cif = &types.CallInterface{}
var ImpellerDisplayListBuilderDrawPath unsafe.Pointer
var ImpellerDisplayListBuilderDrawPath_cif = &types.CallInterface{}
var ImpellerDisplayListBuilderDrawDisplayList unsafe.Pointer
var ImpellerDisplayListBuilderDrawDisplayList_cif = &types.CallInterface{}
var ImpellerDisplayListBuilderDrawParagraph unsafe.Pointer
var ImpellerDisplayListBuilderDrawParagraph_cif = &types.CallInterface{}
var ImpellerDisplayListBuilderDrawShadow unsafe.Pointer
var ImpellerDisplayListBuilderDrawShadow_cif = &types.CallInterface{}
var ImpellerDisplayListBuilderDrawTexture unsafe.Pointer
var ImpellerDisplayListBuilderDrawTexture_cif = &types.CallInterface{}
var ImpellerDisplayListBuilderDrawTextureRect unsafe.Pointer
var ImpellerDisplayListBuilderDrawTextureRect_cif = &types.CallInterface{}
var ImpellerTypographyContextNew unsafe.Pointer
var ImpellerTypographyContextNew_cif = &types.CallInterface{}
var ImpellerTypographyContextRetain unsafe.Pointer
var ImpellerTypographyContextRetain_cif = &types.CallInterface{}
var ImpellerTypographyContextRelease unsafe.Pointer
var ImpellerTypographyContextRelease_cif = &types.CallInterface{}
var ImpellerTypographyContextRegisterFont unsafe.Pointer
var ImpellerTypographyContextRegisterFont_cif = &types.CallInterface{}
var ImpellerParagraphStyleNew unsafe.Pointer
var ImpellerParagraphStyleNew_cif = &types.CallInterface{}
var ImpellerParagraphStyleRetain unsafe.Pointer
var ImpellerParagraphStyleRetain_cif = &types.CallInterface{}
var ImpellerParagraphStyleRelease unsafe.Pointer
var ImpellerParagraphStyleRelease_cif = &types.CallInterface{}
var ImpellerParagraphStyleSetForeground unsafe.Pointer
var ImpellerParagraphStyleSetForeground_cif = &types.CallInterface{}
var ImpellerParagraphStyleSetBackground unsafe.Pointer
var ImpellerParagraphStyleSetBackground_cif = &types.CallInterface{}
var ImpellerParagraphStyleSetFontWeight unsafe.Pointer
var ImpellerParagraphStyleSetFontWeight_cif = &types.CallInterface{}
var ImpellerParagraphStyleSetFontStyle unsafe.Pointer
var ImpellerParagraphStyleSetFontStyle_cif = &types.CallInterface{}
var ImpellerParagraphStyleSetFontFamily unsafe.Pointer
var ImpellerParagraphStyleSetFontFamily_cif = &types.CallInterface{}
var ImpellerParagraphStyleSetFontSize unsafe.Pointer
var ImpellerParagraphStyleSetFontSize_cif = &types.CallInterface{}
var ImpellerParagraphStyleSetHeight unsafe.Pointer
var ImpellerParagraphStyleSetHeight_cif = &types.CallInterface{}
var ImpellerParagraphStyleSetTextAlignment unsafe.Pointer
var ImpellerParagraphStyleSetTextAlignment_cif = &types.CallInterface{}
var ImpellerParagraphStyleSetTextDirection unsafe.Pointer
var ImpellerParagraphStyleSetTextDirection_cif = &types.CallInterface{}
var ImpellerParagraphStyleSetTextDecoration unsafe.Pointer
var ImpellerParagraphStyleSetTextDecoration_cif = &types.CallInterface{}
var ImpellerParagraphStyleSetMaxLines unsafe.Pointer
var ImpellerParagraphStyleSetMaxLines_cif = &types.CallInterface{}
var ImpellerParagraphStyleSetLocale unsafe.Pointer
var ImpellerParagraphStyleSetLocale_cif = &types.CallInterface{}
var ImpellerParagraphStyleSetEllipsis unsafe.Pointer
var ImpellerParagraphStyleSetEllipsis_cif = &types.CallInterface{}
var ImpellerParagraphBuilderNew unsafe.Pointer
var ImpellerParagraphBuilderNew_cif = &types.CallInterface{}
var ImpellerParagraphBuilderRetain unsafe.Pointer
var ImpellerParagraphBuilderRetain_cif = &types.CallInterface{}
var ImpellerParagraphBuilderRelease unsafe.Pointer
var ImpellerParagraphBuilderRelease_cif = &types.CallInterface{}
var ImpellerParagraphBuilderPushStyle unsafe.Pointer
var ImpellerParagraphBuilderPushStyle_cif = &types.CallInterface{}
var ImpellerParagraphBuilderPopStyle unsafe.Pointer
var ImpellerParagraphBuilderPopStyle_cif = &types.CallInterface{}
var ImpellerParagraphBuilderAddText unsafe.Pointer
var ImpellerParagraphBuilderAddText_cif = &types.CallInterface{}
var ImpellerParagraphBuilderBuildParagraphNew unsafe.Pointer
var ImpellerParagraphBuilderBuildParagraphNew_cif = &types.CallInterface{}
var ImpellerParagraphRetain unsafe.Pointer
var ImpellerParagraphRetain_cif = &types.CallInterface{}
var ImpellerParagraphRelease unsafe.Pointer
var ImpellerParagraphRelease_cif = &types.CallInterface{}
var ImpellerParagraphGetMaxWidth unsafe.Pointer
var ImpellerParagraphGetMaxWidth_cif = &types.CallInterface{}
var ImpellerParagraphGetHeight unsafe.Pointer
var ImpellerParagraphGetHeight_cif = &types.CallInterface{}
var ImpellerParagraphGetLongestLineWidth unsafe.Pointer
var ImpellerParagraphGetLongestLineWidth_cif = &types.CallInterface{}
var ImpellerParagraphGetMinIntrinsicWidth unsafe.Pointer
var ImpellerParagraphGetMinIntrinsicWidth_cif = &types.CallInterface{}
var ImpellerParagraphGetMaxIntrinsicWidth unsafe.Pointer
var ImpellerParagraphGetMaxIntrinsicWidth_cif = &types.CallInterface{}
var ImpellerParagraphGetIdeographicBaseline unsafe.Pointer
var ImpellerParagraphGetIdeographicBaseline_cif = &types.CallInterface{}
var ImpellerParagraphGetAlphabeticBaseline unsafe.Pointer
var ImpellerParagraphGetAlphabeticBaseline_cif = &types.CallInterface{}
var ImpellerParagraphGetLineCount unsafe.Pointer
var ImpellerParagraphGetLineCount_cif = &types.CallInterface{}
var ImpellerParagraphGetWordBoundary unsafe.Pointer
var ImpellerParagraphGetWordBoundary_cif = &types.CallInterface{}
var ImpellerParagraphGetLineMetrics unsafe.Pointer
var ImpellerParagraphGetLineMetrics_cif = &types.CallInterface{}
var ImpellerParagraphCreateGlyphInfoAtCodeUnitIndexNew unsafe.Pointer
var ImpellerParagraphCreateGlyphInfoAtCodeUnitIndexNew_cif = &types.CallInterface{}
var ImpellerParagraphCreateGlyphInfoAtParagraphCoordinatesNew unsafe.Pointer
var ImpellerParagraphCreateGlyphInfoAtParagraphCoordinatesNew_cif = &types.CallInterface{}
var ImpellerLineMetricsRetain unsafe.Pointer
var ImpellerLineMetricsRetain_cif = &types.CallInterface{}
var ImpellerLineMetricsRelease unsafe.Pointer
var ImpellerLineMetricsRelease_cif = &types.CallInterface{}
var ImpellerLineMetricsGetUnscaledAscent unsafe.Pointer
var ImpellerLineMetricsGetUnscaledAscent_cif = &types.CallInterface{}
var ImpellerLineMetricsGetAscent unsafe.Pointer
var ImpellerLineMetricsGetAscent_cif = &types.CallInterface{}
var ImpellerLineMetricsGetDescent unsafe.Pointer
var ImpellerLineMetricsGetDescent_cif = &types.CallInterface{}
var ImpellerLineMetricsGetBaseline unsafe.Pointer
var ImpellerLineMetricsGetBaseline_cif = &types.CallInterface{}
var ImpellerLineMetricsIsHardbreak unsafe.Pointer
var ImpellerLineMetricsIsHardbreak_cif = &types.CallInterface{}
var ImpellerLineMetricsGetWidth unsafe.Pointer
var ImpellerLineMetricsGetWidth_cif = &types.CallInterface{}
var ImpellerLineMetricsGetHeight unsafe.Pointer
var ImpellerLineMetricsGetHeight_cif = &types.CallInterface{}
var ImpellerLineMetricsGetLeft unsafe.Pointer
var ImpellerLineMetricsGetLeft_cif = &types.CallInterface{}
var ImpellerLineMetricsGetCodeUnitStartIndex unsafe.Pointer
var ImpellerLineMetricsGetCodeUnitStartIndex_cif = &types.CallInterface{}
var ImpellerLineMetricsGetCodeUnitEndIndex unsafe.Pointer
var ImpellerLineMetricsGetCodeUnitEndIndex_cif = &types.CallInterface{}
var ImpellerLineMetricsGetCodeUnitEndIndexExcludingWhitespace unsafe.Pointer
var ImpellerLineMetricsGetCodeUnitEndIndexExcludingWhitespace_cif = &types.CallInterface{}
var ImpellerLineMetricsGetCodeUnitEndIndexIncludingNewline unsafe.Pointer
var ImpellerLineMetricsGetCodeUnitEndIndexIncludingNewline_cif = &types.CallInterface{}
var ImpellerGlyphInfoRetain unsafe.Pointer
var ImpellerGlyphInfoRetain_cif = &types.CallInterface{}
var ImpellerGlyphInfoRelease unsafe.Pointer
var ImpellerGlyphInfoRelease_cif = &types.CallInterface{}
var ImpellerGlyphInfoGetGraphemeClusterCodeUnitRangeBegin unsafe.Pointer
var ImpellerGlyphInfoGetGraphemeClusterCodeUnitRangeBegin_cif = &types.CallInterface{}
var ImpellerGlyphInfoGetGraphemeClusterCodeUnitRangeEnd unsafe.Pointer
var ImpellerGlyphInfoGetGraphemeClusterCodeUnitRangeEnd_cif = &types.CallInterface{}
var ImpellerGlyphInfoGetGraphemeClusterBounds unsafe.Pointer
var ImpellerGlyphInfoGetGraphemeClusterBounds_cif = &types.CallInterface{}
var ImpellerGlyphInfoIsEllipsis unsafe.Pointer
var ImpellerGlyphInfoIsEllipsis_cif = &types.CallInterface{}
var ImpellerGlyphInfoGetTextDirection unsafe.Pointer
var ImpellerGlyphInfoGetTextDirection_cif = &types.CallInterface{}
var initialised = false

func Init() error {
	if initialised {
		return nil
	}
	defer func() {
		initialised = true
	}()

	filepath, err := libraryFilepath()
	if err != nil {
		return err
	}

	handle, err := ffi.LoadLibrary(filepath)
	if err != nil {
		return err
	}

	ImpellerGetVersion, err = ffi.GetSymbol(handle, "ImpellerGetVersion")
	if err != nil {
		return err
	}
	ImpellerContextCreateOpenGLESNew, err = ffi.GetSymbol(handle, "ImpellerContextCreateOpenGLESNew")
	if err != nil {
		return err
	}
	ImpellerContextCreateMetalNew, err = ffi.GetSymbol(handle, "ImpellerContextCreateMetalNew")
	if err != nil {
		return err
	}
	ImpellerContextCreateVulkanNew, err = ffi.GetSymbol(handle, "ImpellerContextCreateVulkanNew")
	if err != nil {
		return err
	}
	ImpellerContextRetain, err = ffi.GetSymbol(handle, "ImpellerContextRetain")
	if err != nil {
		return err
	}
	ImpellerContextRelease, err = ffi.GetSymbol(handle, "ImpellerContextRelease")
	if err != nil {
		return err
	}
	ImpellerContextGetVulkanInfo, err = ffi.GetSymbol(handle, "ImpellerContextGetVulkanInfo")
	if err != nil {
		return err
	}
	ImpellerVulkanSwapchainCreateNew, err = ffi.GetSymbol(handle, "ImpellerVulkanSwapchainCreateNew")
	if err != nil {
		return err
	}
	ImpellerVulkanSwapchainRetain, err = ffi.GetSymbol(handle, "ImpellerVulkanSwapchainRetain")
	if err != nil {
		return err
	}
	ImpellerVulkanSwapchainRelease, err = ffi.GetSymbol(handle, "ImpellerVulkanSwapchainRelease")
	if err != nil {
		return err
	}
	ImpellerVulkanSwapchainAcquireNextSurfaceNew, err = ffi.GetSymbol(handle, "ImpellerVulkanSwapchainAcquireNextSurfaceNew")
	if err != nil {
		return err
	}
	ImpellerSurfaceCreateWrappedFBONew, err = ffi.GetSymbol(handle, "ImpellerSurfaceCreateWrappedFBONew")
	if err != nil {
		return err
	}
	ImpellerSurfaceCreateWrappedMetalDrawableNew, err = ffi.GetSymbol(handle, "ImpellerSurfaceCreateWrappedMetalDrawableNew")
	if err != nil {
		return err
	}
	ImpellerSurfaceRetain, err = ffi.GetSymbol(handle, "ImpellerSurfaceRetain")
	if err != nil {
		return err
	}
	ImpellerSurfaceRelease, err = ffi.GetSymbol(handle, "ImpellerSurfaceRelease")
	if err != nil {
		return err
	}
	ImpellerSurfaceDrawDisplayList, err = ffi.GetSymbol(handle, "ImpellerSurfaceDrawDisplayList")
	if err != nil {
		return err
	}
	ImpellerSurfacePresent, err = ffi.GetSymbol(handle, "ImpellerSurfacePresent")
	if err != nil {
		return err
	}
	ImpellerPathRetain, err = ffi.GetSymbol(handle, "ImpellerPathRetain")
	if err != nil {
		return err
	}
	ImpellerPathRelease, err = ffi.GetSymbol(handle, "ImpellerPathRelease")
	if err != nil {
		return err
	}
	ImpellerPathGetBounds, err = ffi.GetSymbol(handle, "ImpellerPathGetBounds")
	if err != nil {
		return err
	}
	ImpellerPathBuilderNew, err = ffi.GetSymbol(handle, "ImpellerPathBuilderNew")
	if err != nil {
		return err
	}
	ImpellerPathBuilderRetain, err = ffi.GetSymbol(handle, "ImpellerPathBuilderRetain")
	if err != nil {
		return err
	}
	ImpellerPathBuilderRelease, err = ffi.GetSymbol(handle, "ImpellerPathBuilderRelease")
	if err != nil {
		return err
	}
	ImpellerPathBuilderMoveTo, err = ffi.GetSymbol(handle, "ImpellerPathBuilderMoveTo")
	if err != nil {
		return err
	}
	ImpellerPathBuilderLineTo, err = ffi.GetSymbol(handle, "ImpellerPathBuilderLineTo")
	if err != nil {
		return err
	}
	ImpellerPathBuilderQuadraticCurveTo, err = ffi.GetSymbol(handle, "ImpellerPathBuilderQuadraticCurveTo")
	if err != nil {
		return err
	}
	ImpellerPathBuilderCubicCurveTo, err = ffi.GetSymbol(handle, "ImpellerPathBuilderCubicCurveTo")
	if err != nil {
		return err
	}
	ImpellerPathBuilderAddRect, err = ffi.GetSymbol(handle, "ImpellerPathBuilderAddRect")
	if err != nil {
		return err
	}
	ImpellerPathBuilderAddArc, err = ffi.GetSymbol(handle, "ImpellerPathBuilderAddArc")
	if err != nil {
		return err
	}
	ImpellerPathBuilderAddOval, err = ffi.GetSymbol(handle, "ImpellerPathBuilderAddOval")
	if err != nil {
		return err
	}
	ImpellerPathBuilderAddRoundedRect, err = ffi.GetSymbol(handle, "ImpellerPathBuilderAddRoundedRect")
	if err != nil {
		return err
	}
	ImpellerPathBuilderClose, err = ffi.GetSymbol(handle, "ImpellerPathBuilderClose")
	if err != nil {
		return err
	}
	ImpellerPathBuilderCopyPathNew, err = ffi.GetSymbol(handle, "ImpellerPathBuilderCopyPathNew")
	if err != nil {
		return err
	}
	ImpellerPathBuilderTakePathNew, err = ffi.GetSymbol(handle, "ImpellerPathBuilderTakePathNew")
	if err != nil {
		return err
	}
	ImpellerPaintNew, err = ffi.GetSymbol(handle, "ImpellerPaintNew")
	if err != nil {
		return err
	}
	ImpellerPaintRetain, err = ffi.GetSymbol(handle, "ImpellerPaintRetain")
	if err != nil {
		return err
	}
	ImpellerPaintRelease, err = ffi.GetSymbol(handle, "ImpellerPaintRelease")
	if err != nil {
		return err
	}
	ImpellerPaintSetColor, err = ffi.GetSymbol(handle, "ImpellerPaintSetColor")
	if err != nil {
		return err
	}
	ImpellerPaintSetBlendMode, err = ffi.GetSymbol(handle, "ImpellerPaintSetBlendMode")
	if err != nil {
		return err
	}
	ImpellerPaintSetDrawStyle, err = ffi.GetSymbol(handle, "ImpellerPaintSetDrawStyle")
	if err != nil {
		return err
	}
	ImpellerPaintSetStrokeCap, err = ffi.GetSymbol(handle, "ImpellerPaintSetStrokeCap")
	if err != nil {
		return err
	}
	ImpellerPaintSetStrokeJoin, err = ffi.GetSymbol(handle, "ImpellerPaintSetStrokeJoin")
	if err != nil {
		return err
	}
	ImpellerPaintSetStrokeWidth, err = ffi.GetSymbol(handle, "ImpellerPaintSetStrokeWidth")
	if err != nil {
		return err
	}
	ImpellerPaintSetStrokeMiter, err = ffi.GetSymbol(handle, "ImpellerPaintSetStrokeMiter")
	if err != nil {
		return err
	}
	ImpellerPaintSetColorFilter, err = ffi.GetSymbol(handle, "ImpellerPaintSetColorFilter")
	if err != nil {
		return err
	}
	ImpellerPaintSetColorSource, err = ffi.GetSymbol(handle, "ImpellerPaintSetColorSource")
	if err != nil {
		return err
	}
	ImpellerPaintSetImageFilter, err = ffi.GetSymbol(handle, "ImpellerPaintSetImageFilter")
	if err != nil {
		return err
	}
	ImpellerPaintSetMaskFilter, err = ffi.GetSymbol(handle, "ImpellerPaintSetMaskFilter")
	if err != nil {
		return err
	}
	ImpellerTextureCreateWithContentsNew, err = ffi.GetSymbol(handle, "ImpellerTextureCreateWithContentsNew")
	if err != nil {
		return err
	}
	ImpellerTextureCreateWithOpenGLTextureHandleNew, err = ffi.GetSymbol(handle, "ImpellerTextureCreateWithOpenGLTextureHandleNew")
	if err != nil {
		return err
	}
	ImpellerTextureRetain, err = ffi.GetSymbol(handle, "ImpellerTextureRetain")
	if err != nil {
		return err
	}
	ImpellerTextureRelease, err = ffi.GetSymbol(handle, "ImpellerTextureRelease")
	if err != nil {
		return err
	}
	ImpellerTextureGetOpenGLHandle, err = ffi.GetSymbol(handle, "ImpellerTextureGetOpenGLHandle")
	if err != nil {
		return err
	}
	ImpellerFragmentProgramNew, err = ffi.GetSymbol(handle, "ImpellerFragmentProgramNew")
	if err != nil {
		return err
	}
	ImpellerFragmentProgramRetain, err = ffi.GetSymbol(handle, "ImpellerFragmentProgramRetain")
	if err != nil {
		return err
	}
	ImpellerFragmentProgramRelease, err = ffi.GetSymbol(handle, "ImpellerFragmentProgramRelease")
	if err != nil {
		return err
	}
	ImpellerColorSourceRetain, err = ffi.GetSymbol(handle, "ImpellerColorSourceRetain")
	if err != nil {
		return err
	}
	ImpellerColorSourceRelease, err = ffi.GetSymbol(handle, "ImpellerColorSourceRelease")
	if err != nil {
		return err
	}
	ImpellerColorSourceCreateLinearGradientNew, err = ffi.GetSymbol(handle, "ImpellerColorSourceCreateLinearGradientNew")
	if err != nil {
		return err
	}
	ImpellerColorSourceCreateRadialGradientNew, err = ffi.GetSymbol(handle, "ImpellerColorSourceCreateRadialGradientNew")
	if err != nil {
		return err
	}
	ImpellerColorSourceCreateConicalGradientNew, err = ffi.GetSymbol(handle, "ImpellerColorSourceCreateConicalGradientNew")
	if err != nil {
		return err
	}
	ImpellerColorSourceCreateSweepGradientNew, err = ffi.GetSymbol(handle, "ImpellerColorSourceCreateSweepGradientNew")
	if err != nil {
		return err
	}
	ImpellerColorSourceCreateImageNew, err = ffi.GetSymbol(handle, "ImpellerColorSourceCreateImageNew")
	if err != nil {
		return err
	}
	ImpellerColorSourceCreateFragmentProgramNew, err = ffi.GetSymbol(handle, "ImpellerColorSourceCreateFragmentProgramNew")
	if err != nil {
		return err
	}
	ImpellerColorFilterRetain, err = ffi.GetSymbol(handle, "ImpellerColorFilterRetain")
	if err != nil {
		return err
	}
	ImpellerColorFilterRelease, err = ffi.GetSymbol(handle, "ImpellerColorFilterRelease")
	if err != nil {
		return err
	}
	ImpellerColorFilterCreateBlendNew, err = ffi.GetSymbol(handle, "ImpellerColorFilterCreateBlendNew")
	if err != nil {
		return err
	}
	ImpellerColorFilterCreateColorMatrixNew, err = ffi.GetSymbol(handle, "ImpellerColorFilterCreateColorMatrixNew")
	if err != nil {
		return err
	}
	ImpellerMaskFilterRetain, err = ffi.GetSymbol(handle, "ImpellerMaskFilterRetain")
	if err != nil {
		return err
	}
	ImpellerMaskFilterRelease, err = ffi.GetSymbol(handle, "ImpellerMaskFilterRelease")
	if err != nil {
		return err
	}
	ImpellerMaskFilterCreateBlurNew, err = ffi.GetSymbol(handle, "ImpellerMaskFilterCreateBlurNew")
	if err != nil {
		return err
	}
	ImpellerImageFilterRetain, err = ffi.GetSymbol(handle, "ImpellerImageFilterRetain")
	if err != nil {
		return err
	}
	ImpellerImageFilterRelease, err = ffi.GetSymbol(handle, "ImpellerImageFilterRelease")
	if err != nil {
		return err
	}
	ImpellerImageFilterCreateBlurNew, err = ffi.GetSymbol(handle, "ImpellerImageFilterCreateBlurNew")
	if err != nil {
		return err
	}
	ImpellerImageFilterCreateDilateNew, err = ffi.GetSymbol(handle, "ImpellerImageFilterCreateDilateNew")
	if err != nil {
		return err
	}
	ImpellerImageFilterCreateErodeNew, err = ffi.GetSymbol(handle, "ImpellerImageFilterCreateErodeNew")
	if err != nil {
		return err
	}
	ImpellerImageFilterCreateMatrixNew, err = ffi.GetSymbol(handle, "ImpellerImageFilterCreateMatrixNew")
	if err != nil {
		return err
	}
	ImpellerImageFilterCreateFragmentProgramNew, err = ffi.GetSymbol(handle, "ImpellerImageFilterCreateFragmentProgramNew")
	if err != nil {
		return err
	}
	ImpellerImageFilterCreateComposeNew, err = ffi.GetSymbol(handle, "ImpellerImageFilterCreateComposeNew")
	if err != nil {
		return err
	}
	ImpellerDisplayListRetain, err = ffi.GetSymbol(handle, "ImpellerDisplayListRetain")
	if err != nil {
		return err
	}
	ImpellerDisplayListRelease, err = ffi.GetSymbol(handle, "ImpellerDisplayListRelease")
	if err != nil {
		return err
	}
	ImpellerDisplayListBuilderNew, err = ffi.GetSymbol(handle, "ImpellerDisplayListBuilderNew")
	if err != nil {
		return err
	}
	ImpellerDisplayListBuilderRetain, err = ffi.GetSymbol(handle, "ImpellerDisplayListBuilderRetain")
	if err != nil {
		return err
	}
	ImpellerDisplayListBuilderRelease, err = ffi.GetSymbol(handle, "ImpellerDisplayListBuilderRelease")
	if err != nil {
		return err
	}
	ImpellerDisplayListBuilderCreateDisplayListNew, err = ffi.GetSymbol(handle, "ImpellerDisplayListBuilderCreateDisplayListNew")
	if err != nil {
		return err
	}
	ImpellerDisplayListBuilderSave, err = ffi.GetSymbol(handle, "ImpellerDisplayListBuilderSave")
	if err != nil {
		return err
	}
	ImpellerDisplayListBuilderSaveLayer, err = ffi.GetSymbol(handle, "ImpellerDisplayListBuilderSaveLayer")
	if err != nil {
		return err
	}
	ImpellerDisplayListBuilderRestore, err = ffi.GetSymbol(handle, "ImpellerDisplayListBuilderRestore")
	if err != nil {
		return err
	}
	ImpellerDisplayListBuilderScale, err = ffi.GetSymbol(handle, "ImpellerDisplayListBuilderScale")
	if err != nil {
		return err
	}
	ImpellerDisplayListBuilderRotate, err = ffi.GetSymbol(handle, "ImpellerDisplayListBuilderRotate")
	if err != nil {
		return err
	}
	ImpellerDisplayListBuilderTranslate, err = ffi.GetSymbol(handle, "ImpellerDisplayListBuilderTranslate")
	if err != nil {
		return err
	}
	ImpellerDisplayListBuilderTransform, err = ffi.GetSymbol(handle, "ImpellerDisplayListBuilderTransform")
	if err != nil {
		return err
	}
	ImpellerDisplayListBuilderSetTransform, err = ffi.GetSymbol(handle, "ImpellerDisplayListBuilderSetTransform")
	if err != nil {
		return err
	}
	ImpellerDisplayListBuilderGetTransform, err = ffi.GetSymbol(handle, "ImpellerDisplayListBuilderGetTransform")
	if err != nil {
		return err
	}
	ImpellerDisplayListBuilderResetTransform, err = ffi.GetSymbol(handle, "ImpellerDisplayListBuilderResetTransform")
	if err != nil {
		return err
	}
	ImpellerDisplayListBuilderGetSaveCount, err = ffi.GetSymbol(handle, "ImpellerDisplayListBuilderGetSaveCount")
	if err != nil {
		return err
	}
	ImpellerDisplayListBuilderRestoreToCount, err = ffi.GetSymbol(handle, "ImpellerDisplayListBuilderRestoreToCount")
	if err != nil {
		return err
	}
	ImpellerDisplayListBuilderClipRect, err = ffi.GetSymbol(handle, "ImpellerDisplayListBuilderClipRect")
	if err != nil {
		return err
	}
	ImpellerDisplayListBuilderClipOval, err = ffi.GetSymbol(handle, "ImpellerDisplayListBuilderClipOval")
	if err != nil {
		return err
	}
	ImpellerDisplayListBuilderClipRoundedRect, err = ffi.GetSymbol(handle, "ImpellerDisplayListBuilderClipRoundedRect")
	if err != nil {
		return err
	}
	ImpellerDisplayListBuilderClipPath, err = ffi.GetSymbol(handle, "ImpellerDisplayListBuilderClipPath")
	if err != nil {
		return err
	}
	ImpellerDisplayListBuilderDrawPaint, err = ffi.GetSymbol(handle, "ImpellerDisplayListBuilderDrawPaint")
	if err != nil {
		return err
	}
	ImpellerDisplayListBuilderDrawLine, err = ffi.GetSymbol(handle, "ImpellerDisplayListBuilderDrawLine")
	if err != nil {
		return err
	}
	ImpellerDisplayListBuilderDrawDashedLine, err = ffi.GetSymbol(handle, "ImpellerDisplayListBuilderDrawDashedLine")
	if err != nil {
		return err
	}
	ImpellerDisplayListBuilderDrawRect, err = ffi.GetSymbol(handle, "ImpellerDisplayListBuilderDrawRect")
	if err != nil {
		return err
	}
	ImpellerDisplayListBuilderDrawOval, err = ffi.GetSymbol(handle, "ImpellerDisplayListBuilderDrawOval")
	if err != nil {
		return err
	}
	ImpellerDisplayListBuilderDrawRoundedRect, err = ffi.GetSymbol(handle, "ImpellerDisplayListBuilderDrawRoundedRect")
	if err != nil {
		return err
	}
	ImpellerDisplayListBuilderDrawRoundedRectDifference, err = ffi.GetSymbol(handle, "ImpellerDisplayListBuilderDrawRoundedRectDifference")
	if err != nil {
		return err
	}
	ImpellerDisplayListBuilderDrawPath, err = ffi.GetSymbol(handle, "ImpellerDisplayListBuilderDrawPath")
	if err != nil {
		return err
	}
	ImpellerDisplayListBuilderDrawDisplayList, err = ffi.GetSymbol(handle, "ImpellerDisplayListBuilderDrawDisplayList")
	if err != nil {
		return err
	}
	ImpellerDisplayListBuilderDrawParagraph, err = ffi.GetSymbol(handle, "ImpellerDisplayListBuilderDrawParagraph")
	if err != nil {
		return err
	}
	ImpellerDisplayListBuilderDrawShadow, err = ffi.GetSymbol(handle, "ImpellerDisplayListBuilderDrawShadow")
	if err != nil {
		return err
	}
	ImpellerDisplayListBuilderDrawTexture, err = ffi.GetSymbol(handle, "ImpellerDisplayListBuilderDrawTexture")
	if err != nil {
		return err
	}
	ImpellerDisplayListBuilderDrawTextureRect, err = ffi.GetSymbol(handle, "ImpellerDisplayListBuilderDrawTextureRect")
	if err != nil {
		return err
	}
	ImpellerTypographyContextNew, err = ffi.GetSymbol(handle, "ImpellerTypographyContextNew")
	if err != nil {
		return err
	}
	ImpellerTypographyContextRetain, err = ffi.GetSymbol(handle, "ImpellerTypographyContextRetain")
	if err != nil {
		return err
	}
	ImpellerTypographyContextRelease, err = ffi.GetSymbol(handle, "ImpellerTypographyContextRelease")
	if err != nil {
		return err
	}
	ImpellerTypographyContextRegisterFont, err = ffi.GetSymbol(handle, "ImpellerTypographyContextRegisterFont")
	if err != nil {
		return err
	}
	ImpellerParagraphStyleNew, err = ffi.GetSymbol(handle, "ImpellerParagraphStyleNew")
	if err != nil {
		return err
	}
	ImpellerParagraphStyleRetain, err = ffi.GetSymbol(handle, "ImpellerParagraphStyleRetain")
	if err != nil {
		return err
	}
	ImpellerParagraphStyleRelease, err = ffi.GetSymbol(handle, "ImpellerParagraphStyleRelease")
	if err != nil {
		return err
	}
	ImpellerParagraphStyleSetForeground, err = ffi.GetSymbol(handle, "ImpellerParagraphStyleSetForeground")
	if err != nil {
		return err
	}
	ImpellerParagraphStyleSetBackground, err = ffi.GetSymbol(handle, "ImpellerParagraphStyleSetBackground")
	if err != nil {
		return err
	}
	ImpellerParagraphStyleSetFontWeight, err = ffi.GetSymbol(handle, "ImpellerParagraphStyleSetFontWeight")
	if err != nil {
		return err
	}
	ImpellerParagraphStyleSetFontStyle, err = ffi.GetSymbol(handle, "ImpellerParagraphStyleSetFontStyle")
	if err != nil {
		return err
	}
	ImpellerParagraphStyleSetFontFamily, err = ffi.GetSymbol(handle, "ImpellerParagraphStyleSetFontFamily")
	if err != nil {
		return err
	}
	ImpellerParagraphStyleSetFontSize, err = ffi.GetSymbol(handle, "ImpellerParagraphStyleSetFontSize")
	if err != nil {
		return err
	}
	ImpellerParagraphStyleSetHeight, err = ffi.GetSymbol(handle, "ImpellerParagraphStyleSetHeight")
	if err != nil {
		return err
	}
	ImpellerParagraphStyleSetTextAlignment, err = ffi.GetSymbol(handle, "ImpellerParagraphStyleSetTextAlignment")
	if err != nil {
		return err
	}
	ImpellerParagraphStyleSetTextDirection, err = ffi.GetSymbol(handle, "ImpellerParagraphStyleSetTextDirection")
	if err != nil {
		return err
	}
	ImpellerParagraphStyleSetTextDecoration, err = ffi.GetSymbol(handle, "ImpellerParagraphStyleSetTextDecoration")
	if err != nil {
		return err
	}
	ImpellerParagraphStyleSetMaxLines, err = ffi.GetSymbol(handle, "ImpellerParagraphStyleSetMaxLines")
	if err != nil {
		return err
	}
	ImpellerParagraphStyleSetLocale, err = ffi.GetSymbol(handle, "ImpellerParagraphStyleSetLocale")
	if err != nil {
		return err
	}
	ImpellerParagraphStyleSetEllipsis, err = ffi.GetSymbol(handle, "ImpellerParagraphStyleSetEllipsis")
	if err != nil {
		return err
	}
	ImpellerParagraphBuilderNew, err = ffi.GetSymbol(handle, "ImpellerParagraphBuilderNew")
	if err != nil {
		return err
	}
	ImpellerParagraphBuilderRetain, err = ffi.GetSymbol(handle, "ImpellerParagraphBuilderRetain")
	if err != nil {
		return err
	}
	ImpellerParagraphBuilderRelease, err = ffi.GetSymbol(handle, "ImpellerParagraphBuilderRelease")
	if err != nil {
		return err
	}
	ImpellerParagraphBuilderPushStyle, err = ffi.GetSymbol(handle, "ImpellerParagraphBuilderPushStyle")
	if err != nil {
		return err
	}
	ImpellerParagraphBuilderPopStyle, err = ffi.GetSymbol(handle, "ImpellerParagraphBuilderPopStyle")
	if err != nil {
		return err
	}
	ImpellerParagraphBuilderAddText, err = ffi.GetSymbol(handle, "ImpellerParagraphBuilderAddText")
	if err != nil {
		return err
	}
	ImpellerParagraphBuilderBuildParagraphNew, err = ffi.GetSymbol(handle, "ImpellerParagraphBuilderBuildParagraphNew")
	if err != nil {
		return err
	}
	ImpellerParagraphRetain, err = ffi.GetSymbol(handle, "ImpellerParagraphRetain")
	if err != nil {
		return err
	}
	ImpellerParagraphRelease, err = ffi.GetSymbol(handle, "ImpellerParagraphRelease")
	if err != nil {
		return err
	}
	ImpellerParagraphGetMaxWidth, err = ffi.GetSymbol(handle, "ImpellerParagraphGetMaxWidth")
	if err != nil {
		return err
	}
	ImpellerParagraphGetHeight, err = ffi.GetSymbol(handle, "ImpellerParagraphGetHeight")
	if err != nil {
		return err
	}
	ImpellerParagraphGetLongestLineWidth, err = ffi.GetSymbol(handle, "ImpellerParagraphGetLongestLineWidth")
	if err != nil {
		return err
	}
	ImpellerParagraphGetMinIntrinsicWidth, err = ffi.GetSymbol(handle, "ImpellerParagraphGetMinIntrinsicWidth")
	if err != nil {
		return err
	}
	ImpellerParagraphGetMaxIntrinsicWidth, err = ffi.GetSymbol(handle, "ImpellerParagraphGetMaxIntrinsicWidth")
	if err != nil {
		return err
	}
	ImpellerParagraphGetIdeographicBaseline, err = ffi.GetSymbol(handle, "ImpellerParagraphGetIdeographicBaseline")
	if err != nil {
		return err
	}
	ImpellerParagraphGetAlphabeticBaseline, err = ffi.GetSymbol(handle, "ImpellerParagraphGetAlphabeticBaseline")
	if err != nil {
		return err
	}
	ImpellerParagraphGetLineCount, err = ffi.GetSymbol(handle, "ImpellerParagraphGetLineCount")
	if err != nil {
		return err
	}
	ImpellerParagraphGetWordBoundary, err = ffi.GetSymbol(handle, "ImpellerParagraphGetWordBoundary")
	if err != nil {
		return err
	}
	ImpellerParagraphGetLineMetrics, err = ffi.GetSymbol(handle, "ImpellerParagraphGetLineMetrics")
	if err != nil {
		return err
	}
	ImpellerParagraphCreateGlyphInfoAtCodeUnitIndexNew, err = ffi.GetSymbol(handle, "ImpellerParagraphCreateGlyphInfoAtCodeUnitIndexNew")
	if err != nil {
		return err
	}
	ImpellerParagraphCreateGlyphInfoAtParagraphCoordinatesNew, err = ffi.GetSymbol(handle, "ImpellerParagraphCreateGlyphInfoAtParagraphCoordinatesNew")
	if err != nil {
		return err
	}
	ImpellerLineMetricsRetain, err = ffi.GetSymbol(handle, "ImpellerLineMetricsRetain")
	if err != nil {
		return err
	}
	ImpellerLineMetricsRelease, err = ffi.GetSymbol(handle, "ImpellerLineMetricsRelease")
	if err != nil {
		return err
	}
	ImpellerLineMetricsGetUnscaledAscent, err = ffi.GetSymbol(handle, "ImpellerLineMetricsGetUnscaledAscent")
	if err != nil {
		return err
	}
	ImpellerLineMetricsGetAscent, err = ffi.GetSymbol(handle, "ImpellerLineMetricsGetAscent")
	if err != nil {
		return err
	}
	ImpellerLineMetricsGetDescent, err = ffi.GetSymbol(handle, "ImpellerLineMetricsGetDescent")
	if err != nil {
		return err
	}
	ImpellerLineMetricsGetBaseline, err = ffi.GetSymbol(handle, "ImpellerLineMetricsGetBaseline")
	if err != nil {
		return err
	}
	ImpellerLineMetricsIsHardbreak, err = ffi.GetSymbol(handle, "ImpellerLineMetricsIsHardbreak")
	if err != nil {
		return err
	}
	ImpellerLineMetricsGetWidth, err = ffi.GetSymbol(handle, "ImpellerLineMetricsGetWidth")
	if err != nil {
		return err
	}
	ImpellerLineMetricsGetHeight, err = ffi.GetSymbol(handle, "ImpellerLineMetricsGetHeight")
	if err != nil {
		return err
	}
	ImpellerLineMetricsGetLeft, err = ffi.GetSymbol(handle, "ImpellerLineMetricsGetLeft")
	if err != nil {
		return err
	}
	ImpellerLineMetricsGetCodeUnitStartIndex, err = ffi.GetSymbol(handle, "ImpellerLineMetricsGetCodeUnitStartIndex")
	if err != nil {
		return err
	}
	ImpellerLineMetricsGetCodeUnitEndIndex, err = ffi.GetSymbol(handle, "ImpellerLineMetricsGetCodeUnitEndIndex")
	if err != nil {
		return err
	}
	ImpellerLineMetricsGetCodeUnitEndIndexExcludingWhitespace, err = ffi.GetSymbol(handle, "ImpellerLineMetricsGetCodeUnitEndIndexExcludingWhitespace")
	if err != nil {
		return err
	}
	ImpellerLineMetricsGetCodeUnitEndIndexIncludingNewline, err = ffi.GetSymbol(handle, "ImpellerLineMetricsGetCodeUnitEndIndexIncludingNewline")
	if err != nil {
		return err
	}
	ImpellerGlyphInfoRetain, err = ffi.GetSymbol(handle, "ImpellerGlyphInfoRetain")
	if err != nil {
		return err
	}
	ImpellerGlyphInfoRelease, err = ffi.GetSymbol(handle, "ImpellerGlyphInfoRelease")
	if err != nil {
		return err
	}
	ImpellerGlyphInfoGetGraphemeClusterCodeUnitRangeBegin, err = ffi.GetSymbol(handle, "ImpellerGlyphInfoGetGraphemeClusterCodeUnitRangeBegin")
	if err != nil {
		return err
	}
	ImpellerGlyphInfoGetGraphemeClusterCodeUnitRangeEnd, err = ffi.GetSymbol(handle, "ImpellerGlyphInfoGetGraphemeClusterCodeUnitRangeEnd")
	if err != nil {
		return err
	}
	ImpellerGlyphInfoGetGraphemeClusterBounds, err = ffi.GetSymbol(handle, "ImpellerGlyphInfoGetGraphemeClusterBounds")
	if err != nil {
		return err
	}
	ImpellerGlyphInfoIsEllipsis, err = ffi.GetSymbol(handle, "ImpellerGlyphInfoIsEllipsis")
	if err != nil {
		return err
	}
	ImpellerGlyphInfoGetTextDirection, err = ffi.GetSymbol(handle, "ImpellerGlyphInfoGetTextDirection")
	if err != nil {
		return err
	}

	err = ffi.PrepareCallInterface(
		ImpellerGetVersion_cif,
		types.DefaultCall,
		types.UInt32TypeDescriptor,
		[]*types.TypeDescriptor{})
	if err != nil {
		return err
	}

	return nil
}
