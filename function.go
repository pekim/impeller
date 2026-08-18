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

// UNSUPPORTED ContextCreateOpenGLESNew : param gl_proc_address_callback is "ImpellerProcAddressCallback"

/*
Create a Metal context using the system default Metal device.

@param[in]  version  The version specified in the IMPELLER_VERSION macro.

@return     The Metal context or NULL if one cannot be created.
*/
func ContextCreateMetalNew(version uint32) Context {
	var result Context
	_, err := ffi.CallFunction(
		cifImpellerContextCreateMetalNew,
		funcImpellerContextCreateMetalNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&version),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

// UNSUPPORTED ContextCreateVulkanNew : param settings is "const ImpellerContextVulkanSettings *"

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  context  The context.
*/
func ContextRetain(context Context) {
	_, err := ffi.CallFunction(
		cifImpellerContextRetain,
		funcImpellerContextRetain,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&context),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  context  The context.
*/
func ContextRelease(context Context) {
	_, err := ffi.CallFunction(
		cifImpellerContextRelease,
		funcImpellerContextRelease,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&context),
		},
	)
	if err != nil {
		panic(err)
	}
}

// UNSUPPORTED ContextGetVulkanInfo : param out_vulkan_info is "ImpellerContextVulkanInfo *"

// UNSUPPORTED VulkanSwapchainCreateNew : param vulkan_surface_khr is "void *"

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  swapchain  The swapchain.
*/
func VulkanSwapchainRetain(swapchain VulkanSwapchain) {
	_, err := ffi.CallFunction(
		cifImpellerVulkanSwapchainRetain,
		funcImpellerVulkanSwapchainRetain,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&swapchain),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  swapchain  The swapchain.
*/
func VulkanSwapchainRelease(swapchain VulkanSwapchain) {
	_, err := ffi.CallFunction(
		cifImpellerVulkanSwapchainRelease,
		funcImpellerVulkanSwapchainRelease,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&swapchain),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
A potentially blocking operation, acquires the next surface to
render to. Since this may block, surface acquisition must be
delayed for as long as possible to avoid an idle wait on the
CPU.

@param[in]  swapchain  The swapchain.

@return     The surface if one could be obtained, NULL otherwise.
*/
func VulkanSwapchainAcquireNextSurfaceNew(swapchain VulkanSwapchain) Surface {
	var result Surface
	_, err := ffi.CallFunction(
		cifImpellerVulkanSwapchainAcquireNextSurfaceNew,
		funcImpellerVulkanSwapchainAcquireNextSurfaceNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&swapchain),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

// UNSUPPORTED SurfaceCreateWrappedFBONew : param size is "const ImpellerISize *"

// UNSUPPORTED SurfaceCreateWrappedMetalDrawableNew : param metal_drawable is "void *"

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  surface  The surface.
*/
func SurfaceRetain(surface Surface) {
	_, err := ffi.CallFunction(
		cifImpellerSurfaceRetain,
		funcImpellerSurfaceRetain,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&surface),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  surface  The surface.
*/
func SurfaceRelease(surface Surface) {
	_, err := ffi.CallFunction(
		cifImpellerSurfaceRelease,
		funcImpellerSurfaceRelease,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&surface),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Draw a display list onto the surface. The same display list can
be drawn multiple times to different surfaces.

@warning    In the OpenGL backend, Impeller will not make an effort to
preserve the OpenGL state that is current in the context.
Embedders that perform additional OpenGL operations in the
context should expect the reset state after control transitions
back to them. Key state to watch out for would be the viewports,
stencil rects, test toggles, resource (texture, framebuffer,
buffer) bindings, etc...

@param[in]  surface       The surface to draw the display list to.
@param[in]  display_list  The display list to draw onto the surface.

@return     If the display list could be drawn onto the surface.
*/
func SurfaceDrawDisplayList(surface Surface, display_list DisplayList) Bool {
	var result Bool
	_, err := ffi.CallFunction(
		cifImpellerSurfaceDrawDisplayList,
		funcImpellerSurfaceDrawDisplayList,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&surface),
			unsafe.Pointer(&display_list),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
Present the surface to the underlying window system.

@param[in]  surface  The surface to present.

@return     True if the surface could be presented.
*/
func SurfacePresent(surface Surface) Bool {
	var result Bool
	_, err := ffi.CallFunction(
		cifImpellerSurfacePresent,
		funcImpellerSurfacePresent,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&surface),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  path  The path.
*/
func PathRetain(path Path) {
	_, err := ffi.CallFunction(
		cifImpellerPathRetain,
		funcImpellerPathRetain,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&path),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  path  The path.
*/
func PathRelease(path Path) {
	_, err := ffi.CallFunction(
		cifImpellerPathRelease,
		funcImpellerPathRelease,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&path),
		},
	)
	if err != nil {
		panic(err)
	}
}

// UNSUPPORTED PathGetBounds : param out_bounds is "ImpellerRect *"

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

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  builder  The builder.
*/
func PathBuilderRetain(builder PathBuilder) {
	_, err := ffi.CallFunction(
		cifImpellerPathBuilderRetain,
		funcImpellerPathBuilderRetain,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  builder  The builder.
*/
func PathBuilderRelease(builder PathBuilder) {
	_, err := ffi.CallFunction(
		cifImpellerPathBuilderRelease,
		funcImpellerPathBuilderRelease,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
		},
	)
	if err != nil {
		panic(err)
	}
}

// UNSUPPORTED PathBuilderMoveTo : param location is "const ImpellerPoint *"

// UNSUPPORTED PathBuilderLineTo : param location is "const ImpellerPoint *"

// UNSUPPORTED PathBuilderQuadraticCurveTo : param control_point is "const ImpellerPoint *"

// UNSUPPORTED PathBuilderCubicCurveTo : param control_point_1 is "const ImpellerPoint *"

// UNSUPPORTED PathBuilderAddRect : param rect is "const ImpellerRect *"

// UNSUPPORTED PathBuilderAddArc : param oval_bounds is "const ImpellerRect *"

// UNSUPPORTED PathBuilderAddOval : param oval_bounds is "const ImpellerRect *"

// UNSUPPORTED PathBuilderAddRoundedRect : param rect is "const ImpellerRect *"

/*
Close the path.

@param[in]  builder  The builder.
*/
func PathBuilderClose(builder PathBuilder) {
	_, err := ffi.CallFunction(
		cifImpellerPathBuilderClose,
		funcImpellerPathBuilderClose,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Create a new path by copying the existing built-up path. The
existing path can continue being added to.

@param[in]  builder  The builder.
@param[in]  fill     The fill.

@return     The impeller path.
*/
func PathBuilderCopyPathNew(builder PathBuilder, fill FillType) Path {
	var result Path
	_, err := ffi.CallFunction(
		cifImpellerPathBuilderCopyPathNew,
		funcImpellerPathBuilderCopyPathNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&fill),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
Create a new path using the existing built-up path. The existing
path builder now contains an empty path.

@param[in]  builder  The builder.
@param[in]  fill     The fill.

@return     The impeller path.
*/
func PathBuilderTakePathNew(builder PathBuilder, fill FillType) Path {
	var result Path
	_, err := ffi.CallFunction(
		cifImpellerPathBuilderTakePathNew,
		funcImpellerPathBuilderTakePathNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&fill),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

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

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  paint  The paint.
*/
func PaintRetain(paint Paint) {
	_, err := ffi.CallFunction(
		cifImpellerPaintRetain,
		funcImpellerPaintRetain,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paint),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  paint  The paint.
*/
func PaintRelease(paint Paint) {
	_, err := ffi.CallFunction(
		cifImpellerPaintRelease,
		funcImpellerPaintRelease,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paint),
		},
	)
	if err != nil {
		panic(err)
	}
}

// UNSUPPORTED PaintSetColor : param color is "const ImpellerColor *"

/*
Set the paint blend mode. The blend mode controls how the new
paints contents are mixed with the values already drawn using
previous draw calls.

@param[in]  paint  The paint.
@param[in]  mode   The mode.
*/
func PaintSetBlendMode(paint Paint, mode BlendMode) {
	_, err := ffi.CallFunction(
		cifImpellerPaintSetBlendMode,
		funcImpellerPaintSetBlendMode,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paint),
			unsafe.Pointer(&mode),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Set the paint draw style. The style controls if the closed
shapes are filled and/or stroked.

@param[in]  paint  The paint.
@param[in]  style  The style.
*/
func PaintSetDrawStyle(paint Paint, style DrawStyle) {
	_, err := ffi.CallFunction(
		cifImpellerPaintSetDrawStyle,
		funcImpellerPaintSetDrawStyle,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paint),
			unsafe.Pointer(&style),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Sets how strokes rendered using this paint are capped.

@param[in]  paint  The paint.
@param[in]  cap    The stroke cap style.
*/
func PaintSetStrokeCap(paint Paint, cap StrokeCap) {
	_, err := ffi.CallFunction(
		cifImpellerPaintSetStrokeCap,
		funcImpellerPaintSetStrokeCap,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paint),
			unsafe.Pointer(&cap),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Sets how strokes rendered using this paint are joined.

@param[in]  paint  The paint.
@param[in]  join   The join.
*/
func PaintSetStrokeJoin(paint Paint, join StrokeJoin) {
	_, err := ffi.CallFunction(
		cifImpellerPaintSetStrokeJoin,
		funcImpellerPaintSetStrokeJoin,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paint),
			unsafe.Pointer(&join),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Set the width of the strokes rendered using this paint.

@param[in]  paint  The paint.
@param[in]  width  The width.
*/
func PaintSetStrokeWidth(paint Paint, width float32) {
	_, err := ffi.CallFunction(
		cifImpellerPaintSetStrokeWidth,
		funcImpellerPaintSetStrokeWidth,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paint),
			unsafe.Pointer(&width),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Set the miter limit of the strokes rendered using this paint.

@param[in]  paint  The paint.
@param[in]  miter  The miter limit.
*/
func PaintSetStrokeMiter(paint Paint, miter float32) {
	_, err := ffi.CallFunction(
		cifImpellerPaintSetStrokeMiter,
		funcImpellerPaintSetStrokeMiter,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paint),
			unsafe.Pointer(&miter),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Set the color filter of the paint.

Color filters are functions that take two colors and mix them to
produce a single color. This color is then usually merged with
the destination during blending.

@param[in]  paint         The paint.
@param[in]  color_filter  The color filter.
*/
func PaintSetColorFilter(paint Paint, color_filter ColorFilter) {
	_, err := ffi.CallFunction(
		cifImpellerPaintSetColorFilter,
		funcImpellerPaintSetColorFilter,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paint),
			unsafe.Pointer(&color_filter),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Set the color source of the paint.

Color sources are functions that generate colors for each
texture element covered by a draw call.

@param[in]  paint         The paint.
@param[in]  color_source  The color source.
*/
func PaintSetColorSource(paint Paint, color_source ColorSource) {
	_, err := ffi.CallFunction(
		cifImpellerPaintSetColorSource,
		funcImpellerPaintSetColorSource,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paint),
			unsafe.Pointer(&color_source),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Set the image filter of a paint.

Image filters are functions that are applied to regions of a
texture to produce a single color.

@param[in]  paint         The paint.
@param[in]  image_filter  The image filter.
*/
func PaintSetImageFilter(paint Paint, image_filter ImageFilter) {
	_, err := ffi.CallFunction(
		cifImpellerPaintSetImageFilter,
		funcImpellerPaintSetImageFilter,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paint),
			unsafe.Pointer(&image_filter),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Set the mask filter of a paint.

@param[in]  paint        The paint.
@param[in]  mask_filter  The mask filter.
*/
func PaintSetMaskFilter(paint Paint, mask_filter MaskFilter) {
	_, err := ffi.CallFunction(
		cifImpellerPaintSetMaskFilter,
		funcImpellerPaintSetMaskFilter,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paint),
			unsafe.Pointer(&mask_filter),
		},
	)
	if err != nil {
		panic(err)
	}
}

// UNSUPPORTED TextureCreateWithContentsNew : param descriptor is "const ImpellerTextureDescriptor *"

// UNSUPPORTED TextureCreateWithOpenGLTextureHandleNew : param descriptor is "const ImpellerTextureDescriptor *"

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  texture  The texture.
*/
func TextureRetain(texture Texture) {
	_, err := ffi.CallFunction(
		cifImpellerTextureRetain,
		funcImpellerTextureRetain,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&texture),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  texture  The texture.
*/
func TextureRelease(texture Texture) {
	_, err := ffi.CallFunction(
		cifImpellerTextureRelease,
		funcImpellerTextureRelease,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&texture),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Get the OpenGL handle associated with this texture. If this is
not an OpenGL texture, this method will always return 0.

OpenGL handles are lazily created, this method will return
GL_NONE is no OpenGL handle is available. To ensure that this
call eagerly creates an OpenGL texture, call this on a thread
where Impeller knows there is an OpenGL context available.

@param[in]  texture  The texture.

@return     The OpenGL handle if one is available, GL_NONE otherwise.
*/
func TextureGetOpenGLHandle(texture Texture) uint64 {
	var result uint64
	_, err := ffi.CallFunction(
		cifImpellerTextureGetOpenGLHandle,
		funcImpellerTextureGetOpenGLHandle,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&texture),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

// UNSUPPORTED FragmentProgramNew : param data is "const ImpellerMapping *"

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  fragment_program  The fragment program.
*/
func FragmentProgramRetain(fragment_program FragmentProgram) {
	_, err := ffi.CallFunction(
		cifImpellerFragmentProgramRetain,
		funcImpellerFragmentProgramRetain,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&fragment_program),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  fragment_program  The fragment program.
*/
func FragmentProgramRelease(fragment_program FragmentProgram) {
	_, err := ffi.CallFunction(
		cifImpellerFragmentProgramRelease,
		funcImpellerFragmentProgramRelease,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&fragment_program),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  color_source  The color source.
*/
func ColorSourceRetain(color_source ColorSource) {
	_, err := ffi.CallFunction(
		cifImpellerColorSourceRetain,
		funcImpellerColorSourceRetain,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&color_source),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  color_source  The color source.
*/
func ColorSourceRelease(color_source ColorSource) {
	_, err := ffi.CallFunction(
		cifImpellerColorSourceRelease,
		funcImpellerColorSourceRelease,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&color_source),
		},
	)
	if err != nil {
		panic(err)
	}
}

// UNSUPPORTED ColorSourceCreateLinearGradientNew : param start_point is "const ImpellerPoint *"

// UNSUPPORTED ColorSourceCreateRadialGradientNew : param center is "const ImpellerPoint *"

// UNSUPPORTED ColorSourceCreateConicalGradientNew : param start_center is "const ImpellerPoint *"

// UNSUPPORTED ColorSourceCreateSweepGradientNew : param center is "const ImpellerPoint *"

// UNSUPPORTED ColorSourceCreateImageNew : param transformation is "const ImpellerMatrix *"

// UNSUPPORTED ColorSourceCreateFragmentProgramNew : param samplers is "ImpellerTexture  _Nonnull *"

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  color_filter  The color filter.
*/
func ColorFilterRetain(color_filter ColorFilter) {
	_, err := ffi.CallFunction(
		cifImpellerColorFilterRetain,
		funcImpellerColorFilterRetain,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&color_filter),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  color_filter  The color filter.
*/
func ColorFilterRelease(color_filter ColorFilter) {
	_, err := ffi.CallFunction(
		cifImpellerColorFilterRelease,
		funcImpellerColorFilterRelease,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&color_filter),
		},
	)
	if err != nil {
		panic(err)
	}
}

// UNSUPPORTED ColorFilterCreateBlendNew : param color is "const ImpellerColor *"

// UNSUPPORTED ColorFilterCreateColorMatrixNew : param color_matrix is "const ImpellerColorMatrix *"

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  mask_filter  The mask filter.
*/
func MaskFilterRetain(mask_filter MaskFilter) {
	_, err := ffi.CallFunction(
		cifImpellerMaskFilterRetain,
		funcImpellerMaskFilterRetain,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&mask_filter),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  mask_filter  The mask filter.
*/
func MaskFilterRelease(mask_filter MaskFilter) {
	_, err := ffi.CallFunction(
		cifImpellerMaskFilterRelease,
		funcImpellerMaskFilterRelease,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&mask_filter),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Create a mask filter that blurs contents in the masked shape.

@param[in]  style  The style.
@param[in]  sigma  The sigma.

@return     The mask filter.
*/
func MaskFilterCreateBlurNew(style BlurStyle, sigma float32) MaskFilter {
	var result MaskFilter
	_, err := ffi.CallFunction(
		cifImpellerMaskFilterCreateBlurNew,
		funcImpellerMaskFilterCreateBlurNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&style),
			unsafe.Pointer(&sigma),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  image_filter  The image filter.
*/
func ImageFilterRetain(image_filter ImageFilter) {
	_, err := ffi.CallFunction(
		cifImpellerImageFilterRetain,
		funcImpellerImageFilterRetain,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&image_filter),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  image_filter  The image filter.
*/
func ImageFilterRelease(image_filter ImageFilter) {
	_, err := ffi.CallFunction(
		cifImpellerImageFilterRelease,
		funcImpellerImageFilterRelease,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&image_filter),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Creates an image filter that applies a Gaussian blur.

The Gaussian blur applied may be an approximation for
performance.

@param[in]  x_sigma    The x sigma.
@param[in]  y_sigma    The y sigma.
@param[in]  tile_mode  The tile mode.

@return     The image filter.
*/
func ImageFilterCreateBlurNew(x_sigma float32, y_sigma float32, tile_mode TileMode) ImageFilter {
	var result ImageFilter
	_, err := ffi.CallFunction(
		cifImpellerImageFilterCreateBlurNew,
		funcImpellerImageFilterCreateBlurNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&x_sigma),
			unsafe.Pointer(&y_sigma),
			unsafe.Pointer(&tile_mode),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
Creates an image filter that enhances the per-channel pixel
values to the maximum value in a circle around the pixel.

@param[in]  x_radius  The x radius.
@param[in]  y_radius  The y radius.

@return     The image filter.
*/
func ImageFilterCreateDilateNew(x_radius float32, y_radius float32) ImageFilter {
	var result ImageFilter
	_, err := ffi.CallFunction(
		cifImpellerImageFilterCreateDilateNew,
		funcImpellerImageFilterCreateDilateNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&x_radius),
			unsafe.Pointer(&y_radius),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
Creates an image filter that dampens the per-channel pixel
values to the minimum value in a circle around the pixel.

@param[in]  x_radius  The x radius.
@param[in]  y_radius  The y radius.

@return     The image filter.
*/
func ImageFilterCreateErodeNew(x_radius float32, y_radius float32) ImageFilter {
	var result ImageFilter
	_, err := ffi.CallFunction(
		cifImpellerImageFilterCreateErodeNew,
		funcImpellerImageFilterCreateErodeNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&x_radius),
			unsafe.Pointer(&y_radius),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

// UNSUPPORTED ImageFilterCreateMatrixNew : param matrix is "const ImpellerMatrix *"

// UNSUPPORTED ImageFilterCreateFragmentProgramNew : param samplers is "ImpellerTexture  _Nonnull *"

/*
Creates a composed filter that when applied is identical to
subsequently applying the inner and then the outer filters.

	destination = outer_filter(inner_filter(source))

@param[in]  outer  The outer image filter.
@param[in]  inner  The inner image filter.

@return     The combined image filter.
*/
func ImageFilterCreateComposeNew(outer ImageFilter, inner ImageFilter) ImageFilter {
	var result ImageFilter
	_, err := ffi.CallFunction(
		cifImpellerImageFilterCreateComposeNew,
		funcImpellerImageFilterCreateComposeNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&outer),
			unsafe.Pointer(&inner),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  display_list  The display list.
*/
func DisplayListRetain(display_list DisplayList) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListRetain,
		funcImpellerDisplayListRetain,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&display_list),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  display_list  The display list.
*/
func DisplayListRelease(display_list DisplayList) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListRelease,
		funcImpellerDisplayListRelease,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&display_list),
		},
	)
	if err != nil {
		panic(err)
	}
}

// UNSUPPORTED DisplayListBuilderNew : param cull_rect is "const ImpellerRect *"

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  builder  The display list builder.
*/
func DisplayListBuilderRetain(builder DisplayListBuilder) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderRetain,
		funcImpellerDisplayListBuilderRetain,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  builder  The display list builder.
*/
func DisplayListBuilderRelease(builder DisplayListBuilder) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderRelease,
		funcImpellerDisplayListBuilderRelease,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Create a new display list using the rendering intent already
encoded in the builder. The builder is reset after this call.

@param[in]  builder  The builder.

@return     The display list.
*/
func DisplayListBuilderCreateDisplayListNew(builder DisplayListBuilder) DisplayList {
	var result DisplayList
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderCreateDisplayListNew,
		funcImpellerDisplayListBuilderCreateDisplayListNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
Stashes the current transformation and clip state onto a save
stack.

@param[in]  builder  The builder.
*/
func DisplayListBuilderSave(builder DisplayListBuilder) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderSave,
		funcImpellerDisplayListBuilderSave,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
		},
	)
	if err != nil {
		panic(err)
	}
}

// UNSUPPORTED DisplayListBuilderSaveLayer : param bounds is "const ImpellerRect *"

/*
Pops the last entry pushed onto the save stack using a call to
`ImpellerDisplayListBuilderSave` or
`ImpellerDisplayListBuilderSaveLayer`.

@param[in]  builder  The builder.
*/
func DisplayListBuilderRestore(builder DisplayListBuilder) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderRestore,
		funcImpellerDisplayListBuilderRestore,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Apply a scale to the transformation matrix currently on top of
the save stack.

@param[in]  builder  The builder.
@param[in]  x_scale  The x scale.
@param[in]  y_scale  The y scale.
*/
func DisplayListBuilderScale(builder DisplayListBuilder, x_scale float32, y_scale float32) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderScale,
		funcImpellerDisplayListBuilderScale,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&x_scale),
			unsafe.Pointer(&y_scale),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Apply a clockwise rotation to the transformation matrix
currently on top of the save stack.

@param[in]  builder        The builder.
@param[in]  angle_degrees  The angle in degrees.
*/
func DisplayListBuilderRotate(builder DisplayListBuilder, angle_degrees float32) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderRotate,
		funcImpellerDisplayListBuilderRotate,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&angle_degrees),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Apply a translation to the transformation matrix currently on
top of the save stack.

@param[in]  builder        The builder.
@param[in]  x_translation  The x translation.
@param[in]  y_translation  The y translation.
*/
func DisplayListBuilderTranslate(builder DisplayListBuilder, x_translation float32, y_translation float32) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderTranslate,
		funcImpellerDisplayListBuilderTranslate,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&x_translation),
			unsafe.Pointer(&y_translation),
		},
	)
	if err != nil {
		panic(err)
	}
}

// UNSUPPORTED DisplayListBuilderTransform : param transform is "const ImpellerMatrix *"

// UNSUPPORTED DisplayListBuilderSetTransform : param transform is "const ImpellerMatrix *"

// UNSUPPORTED DisplayListBuilderGetTransform : param out_transform is "ImpellerMatrix *"

/*
Reset the transformation on top of the transformation stack to
identity.

@param[in]  builder  The builder.
*/
func DisplayListBuilderResetTransform(builder DisplayListBuilder) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderResetTransform,
		funcImpellerDisplayListBuilderResetTransform,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Get the current size of the save stack.

@param[in]  builder  The builder.

@return     The save stack size.
*/
func DisplayListBuilderGetSaveCount(builder DisplayListBuilder) uint32 {
	var result uint32
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderGetSaveCount,
		funcImpellerDisplayListBuilderGetSaveCount,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
Effectively calls ImpellerDisplayListBuilderRestore till the
size of the save stack becomes a specified count.

@param[in]  builder  The builder.
@param[in]  count    The count.
*/
func DisplayListBuilderRestoreToCount(builder DisplayListBuilder, count uint32) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderRestoreToCount,
		funcImpellerDisplayListBuilderRestoreToCount,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&count),
		},
	)
	if err != nil {
		panic(err)
	}
}

// UNSUPPORTED DisplayListBuilderClipRect : param rect is "const ImpellerRect *"

// UNSUPPORTED DisplayListBuilderClipOval : param oval_bounds is "const ImpellerRect *"

// UNSUPPORTED DisplayListBuilderClipRoundedRect : param rect is "const ImpellerRect *"

/*
Reduces the clip region to the intersection of the current clip
and the given path taking into account the clip operation.

@param[in]  builder  The builder.
@param[in]  path     The path.
@param[in]  op       The operation.
*/
func DisplayListBuilderClipPath(builder DisplayListBuilder, path Path, op ClipOperation) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderClipPath,
		funcImpellerDisplayListBuilderClipPath,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&path),
			unsafe.Pointer(&op),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Fills the current clip with the specified paint.

@param[in]  builder  The builder.
@param[in]  paint    The paint.
*/
func DisplayListBuilderDrawPaint(builder DisplayListBuilder, paint Paint) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderDrawPaint,
		funcImpellerDisplayListBuilderDrawPaint,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&paint),
		},
	)
	if err != nil {
		panic(err)
	}
}

// UNSUPPORTED DisplayListBuilderDrawLine : param from is "const ImpellerPoint *"

// UNSUPPORTED DisplayListBuilderDrawDashedLine : param from is "const ImpellerPoint *"

// UNSUPPORTED DisplayListBuilderDrawRect : param rect is "const ImpellerRect *"

// UNSUPPORTED DisplayListBuilderDrawOval : param oval_bounds is "const ImpellerRect *"

// UNSUPPORTED DisplayListBuilderDrawRoundedRect : param rect is "const ImpellerRect *"

// UNSUPPORTED DisplayListBuilderDrawRoundedRectDifference : param outer_rect is "const ImpellerRect *"

/*
Draws the specified shape.

@param[in]  builder  The builder.
@param[in]  path     The path.
@param[in]  paint    The paint.
*/
func DisplayListBuilderDrawPath(builder DisplayListBuilder, path Path, paint Paint) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderDrawPath,
		funcImpellerDisplayListBuilderDrawPath,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&path),
			unsafe.Pointer(&paint),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Flattens the contents of another display list into the one
currently being built.

@param[in]  builder       The builder.
@param[in]  display_list  The display list.
@param[in]  opacity       The opacity.
*/
func DisplayListBuilderDrawDisplayList(builder DisplayListBuilder, display_list DisplayList, opacity float32) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderDrawDisplayList,
		funcImpellerDisplayListBuilderDrawDisplayList,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&display_list),
			unsafe.Pointer(&opacity),
		},
	)
	if err != nil {
		panic(err)
	}
}

// UNSUPPORTED DisplayListBuilderDrawParagraph : param point is "const ImpellerPoint *"

// UNSUPPORTED DisplayListBuilderDrawShadow : param color is "const ImpellerColor *"

// UNSUPPORTED DisplayListBuilderDrawTexture : param point is "const ImpellerPoint *"

// UNSUPPORTED DisplayListBuilderDrawTextureRect : param src_rect is "const ImpellerRect *"

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

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  context  The typography context.
*/
func TypographyContextRetain(context TypographyContext) {
	_, err := ffi.CallFunction(
		cifImpellerTypographyContextRetain,
		funcImpellerTypographyContextRetain,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&context),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  context  The typography context.
*/
func TypographyContextRelease(context TypographyContext) {
	_, err := ffi.CallFunction(
		cifImpellerTypographyContextRelease,
		funcImpellerTypographyContextRelease,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&context),
		},
	)
	if err != nil {
		panic(err)
	}
}

// UNSUPPORTED TypographyContextRegisterFont : param contents is "const ImpellerMapping *"

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

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  paragraph_style  The paragraph style.
*/
func ParagraphStyleRetain(paragraph_style ParagraphStyle) {
	_, err := ffi.CallFunction(
		cifImpellerParagraphStyleRetain,
		funcImpellerParagraphStyleRetain,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph_style),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  paragraph_style  The paragraph style.
*/
func ParagraphStyleRelease(paragraph_style ParagraphStyle) {
	_, err := ffi.CallFunction(
		cifImpellerParagraphStyleRelease,
		funcImpellerParagraphStyleRelease,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph_style),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Set the paint used to render the text glyph contents.

@param[in]  paragraph_style  The paragraph style.
@param[in]  paint            The paint.
*/
func ParagraphStyleSetForeground(paragraph_style ParagraphStyle, paint Paint) {
	_, err := ffi.CallFunction(
		cifImpellerParagraphStyleSetForeground,
		funcImpellerParagraphStyleSetForeground,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph_style),
			unsafe.Pointer(&paint),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Set the paint used to render the background of the text glyphs.

@param[in]  paragraph_style  The paragraph style.
@param[in]  paint            The paint.
*/
func ParagraphStyleSetBackground(paragraph_style ParagraphStyle, paint Paint) {
	_, err := ffi.CallFunction(
		cifImpellerParagraphStyleSetBackground,
		funcImpellerParagraphStyleSetBackground,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph_style),
			unsafe.Pointer(&paint),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Set the weight of the font to select when rendering glyphs.

@param[in]  paragraph_style  The paragraph style.
@param[in]  weight           The weight.
*/
func ParagraphStyleSetFontWeight(paragraph_style ParagraphStyle, weight FontWeight) {
	_, err := ffi.CallFunction(
		cifImpellerParagraphStyleSetFontWeight,
		funcImpellerParagraphStyleSetFontWeight,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph_style),
			unsafe.Pointer(&weight),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Set whether the glyphs should be bolded or italicized.

@param[in]  paragraph_style  The paragraph style.
@param[in]  style            The style.
*/
func ParagraphStyleSetFontStyle(paragraph_style ParagraphStyle, style FontStyle) {
	_, err := ffi.CallFunction(
		cifImpellerParagraphStyleSetFontStyle,
		funcImpellerParagraphStyleSetFontStyle,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph_style),
			unsafe.Pointer(&style),
		},
	)
	if err != nil {
		panic(err)
	}
}

// UNSUPPORTED ParagraphStyleSetFontFamily : param family_name is "const char *"

/*
Set the font size.

@param[in]  paragraph_style  The paragraph style.
@param[in]  size             The size.
*/
func ParagraphStyleSetFontSize(paragraph_style ParagraphStyle, size float32) {
	_, err := ffi.CallFunction(
		cifImpellerParagraphStyleSetFontSize,
		funcImpellerParagraphStyleSetFontSize,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph_style),
			unsafe.Pointer(&size),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
The height of the text as a multiple of text size.

When height is 0.0, the line height will be determined by the
font's metrics directly, which may differ from the font size.
Otherwise the line height of the text will be a multiple of font
size, and be exactly fontSize * height logical pixels tall.

@param[in]  paragraph_style  The paragraph style.
@param[in]  height           The height.
*/
func ParagraphStyleSetHeight(paragraph_style ParagraphStyle, height float32) {
	_, err := ffi.CallFunction(
		cifImpellerParagraphStyleSetHeight,
		funcImpellerParagraphStyleSetHeight,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph_style),
			unsafe.Pointer(&height),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Set the alignment of text within the paragraph.

@param[in]  paragraph_style  The paragraph style.
@param[in]  align            The align.
*/
func ParagraphStyleSetTextAlignment(paragraph_style ParagraphStyle, align TextAlignment) {
	_, err := ffi.CallFunction(
		cifImpellerParagraphStyleSetTextAlignment,
		funcImpellerParagraphStyleSetTextAlignment,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph_style),
			unsafe.Pointer(&align),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Set the directionality of the text within the paragraph.

@param[in]  paragraph_style  The paragraph style.
@param[in]  direction        The direction.
*/
func ParagraphStyleSetTextDirection(paragraph_style ParagraphStyle, direction TextDirection) {
	_, err := ffi.CallFunction(
		cifImpellerParagraphStyleSetTextDirection,
		funcImpellerParagraphStyleSetTextDirection,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph_style),
			unsafe.Pointer(&direction),
		},
	)
	if err != nil {
		panic(err)
	}
}

// UNSUPPORTED ParagraphStyleSetTextDecoration : param decoration is "const ImpellerTextDecoration *"

/*
Set the maximum line count within the paragraph.

@param[in]  paragraph_style  The paragraph style.
@param[in]  max_lines        The maximum lines.
*/
func ParagraphStyleSetMaxLines(paragraph_style ParagraphStyle, max_lines uint32) {
	_, err := ffi.CallFunction(
		cifImpellerParagraphStyleSetMaxLines,
		funcImpellerParagraphStyleSetMaxLines,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph_style),
			unsafe.Pointer(&max_lines),
		},
	)
	if err != nil {
		panic(err)
	}
}

// UNSUPPORTED ParagraphStyleSetLocale : param locale is "const char *"

// UNSUPPORTED ParagraphStyleSetEllipsis : param ellipsis is "const char *"

/*
Create a new paragraph builder.

@param[in]  context  The context.

@return     The paragraph builder.
*/
func ParagraphBuilderNew(context TypographyContext) ParagraphBuilder {
	var result ParagraphBuilder
	_, err := ffi.CallFunction(
		cifImpellerParagraphBuilderNew,
		funcImpellerParagraphBuilderNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&context),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  paragraph_builder  The paragraph builder.
*/
func ParagraphBuilderRetain(paragraph_builder ParagraphBuilder) {
	_, err := ffi.CallFunction(
		cifImpellerParagraphBuilderRetain,
		funcImpellerParagraphBuilderRetain,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph_builder),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  paragraph_builder  The paragraph_builder.
*/
func ParagraphBuilderRelease(paragraph_builder ParagraphBuilder) {
	_, err := ffi.CallFunction(
		cifImpellerParagraphBuilderRelease,
		funcImpellerParagraphBuilderRelease,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph_builder),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Push a new paragraph style onto the paragraph style stack
managed by the paragraph builder.

Not all paragraph styles can be combined. For instance, it does
not make sense to mix text alignment for different text runs
within a paragraph. In such cases, the preference of the the
first paragraph style on the style stack will take hold.

If text is pushed onto the paragraph builder without a style
previously pushed onto the stack, a default paragraph text style
will be used. This may not always be desirable because some
style element cannot be overridden. It is recommended that a
default paragraph style always be pushed onto the stack before
the addition of any text.

@param[in]  paragraph_builder  The paragraph builder.
@param[in]  style              The style.
*/
func ParagraphBuilderPushStyle(paragraph_builder ParagraphBuilder, style ParagraphStyle) {
	_, err := ffi.CallFunction(
		cifImpellerParagraphBuilderPushStyle,
		funcImpellerParagraphBuilderPushStyle,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph_builder),
			unsafe.Pointer(&style),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Pop a previously pushed paragraph style from the paragraph style
stack.

@param[in]  paragraph_builder  The paragraph builder.
*/
func ParagraphBuilderPopStyle(paragraph_builder ParagraphBuilder) {
	_, err := ffi.CallFunction(
		cifImpellerParagraphBuilderPopStyle,
		funcImpellerParagraphBuilderPopStyle,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph_builder),
		},
	)
	if err != nil {
		panic(err)
	}
}

// UNSUPPORTED ParagraphBuilderAddText : param data is "const uint8_t *"

/*
Layout and build a new paragraph using the specified width. The
resulting paragraph is immutable. The paragraph builder must be
discarded and a new one created to build more paragraphs.

@param[in]  paragraph_builder  The paragraph builder.
@param[in]  width              The paragraph width.

@return     The paragraph if one can be created, NULL otherwise.
*/
func ParagraphBuilderBuildParagraphNew(paragraph_builder ParagraphBuilder, width float32) Paragraph {
	var result Paragraph
	_, err := ffi.CallFunction(
		cifImpellerParagraphBuilderBuildParagraphNew,
		funcImpellerParagraphBuilderBuildParagraphNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph_builder),
			unsafe.Pointer(&width),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  paragraph  The paragraph.
*/
func ParagraphRetain(paragraph Paragraph) {
	_, err := ffi.CallFunction(
		cifImpellerParagraphRetain,
		funcImpellerParagraphRetain,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  paragraph  The paragraph.
*/
func ParagraphRelease(paragraph Paragraph) {
	_, err := ffi.CallFunction(
		cifImpellerParagraphRelease,
		funcImpellerParagraphRelease,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
@see        `ImpellerParagraphGetMinIntrinsicWidth`

@param[in]  paragraph  The paragraph.

@return     The width provided to the paragraph builder during the call to
layout. This is the maximum width any line in the laid out
paragraph can occupy. But, it is not necessarily the actual
width of the paragraph after layout.
*/
func ParagraphGetMaxWidth(paragraph Paragraph) float32 {
	var result float32
	_, err := ffi.CallFunction(
		cifImpellerParagraphGetMaxWidth,
		funcImpellerParagraphGetMaxWidth,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
@param[in]  paragraph  The paragraph.

@return     The height of the laid out paragraph. This is **not** a tight
bounding box and some glyphs may not reach the minimum location
they are allowed to reach.
*/
func ParagraphGetHeight(paragraph Paragraph) float32 {
	var result float32
	_, err := ffi.CallFunction(
		cifImpellerParagraphGetHeight,
		funcImpellerParagraphGetHeight,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
@param[in]  paragraph  The paragraph.

@return     The length of the longest line in the paragraph. This is the
horizontal distance between the left edge of the leftmost glyph
and the right edge of the rightmost glyph, in the longest line
in the paragraph.
*/
func ParagraphGetLongestLineWidth(paragraph Paragraph) float32 {
	var result float32
	_, err := ffi.CallFunction(
		cifImpellerParagraphGetLongestLineWidth,
		funcImpellerParagraphGetLongestLineWidth,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
@see        `ImpellerParagraphGetMaxWidth`

@param[in]  paragraph  The paragraph.

@return     The actual width of the longest line in the paragraph after
layout. This is expected to be less than or equal to
`ImpellerParagraphGetMaxWidth`.
*/
func ParagraphGetMinIntrinsicWidth(paragraph Paragraph) float32 {
	var result float32
	_, err := ffi.CallFunction(
		cifImpellerParagraphGetMinIntrinsicWidth,
		funcImpellerParagraphGetMinIntrinsicWidth,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
@param[in]  paragraph  The paragraph.

@return     The width of the paragraph without line breaking.
*/
func ParagraphGetMaxIntrinsicWidth(paragraph Paragraph) float32 {
	var result float32
	_, err := ffi.CallFunction(
		cifImpellerParagraphGetMaxIntrinsicWidth,
		funcImpellerParagraphGetMaxIntrinsicWidth,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
@param[in]  paragraph  The paragraph.

@return     The distance from the top of the paragraph to the ideographic
baseline of the first line when using ideographic fonts
(Japanese, Korean, etc...).
*/
func ParagraphGetIdeographicBaseline(paragraph Paragraph) float32 {
	var result float32
	_, err := ffi.CallFunction(
		cifImpellerParagraphGetIdeographicBaseline,
		funcImpellerParagraphGetIdeographicBaseline,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
@param[in]  paragraph  The paragraph.

@return     The distance from the top of the paragraph to the alphabetic
baseline of the first line when using alphabetic fonts (A-Z,
a-z, Greek, etc...).
*/
func ParagraphGetAlphabeticBaseline(paragraph Paragraph) float32 {
	var result float32
	_, err := ffi.CallFunction(
		cifImpellerParagraphGetAlphabeticBaseline,
		funcImpellerParagraphGetAlphabeticBaseline,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
@param[in]  paragraph  The paragraph.

@return     The number of lines visible in the paragraph after line
breaking.
*/
func ParagraphGetLineCount(paragraph Paragraph) uint32 {
	var result uint32
	_, err := ffi.CallFunction(
		cifImpellerParagraphGetLineCount,
		funcImpellerParagraphGetLineCount,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

// UNSUPPORTED ParagraphGetWordBoundary : param out_range is "ImpellerRange *"

/*
Get the line metrics of this laid out paragraph. Calculating the
line metrics is expensive. The first time line metrics are
requested, they will be cached along with the paragraph (which
is immutable).

@param[in]  paragraph  The paragraph.

@return     The line metrics.
*/
func ParagraphGetLineMetrics(paragraph Paragraph) LineMetrics {
	var result LineMetrics
	_, err := ffi.CallFunction(
		cifImpellerParagraphGetLineMetrics,
		funcImpellerParagraphGetLineMetrics,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
Create a new instance of glyph info that can be queried for
information about the glyph at the given UTF-16 code unit index.
The instance must be freed using `ImpellerGlyphInfoRelease`.

@param[in]  paragraph        The paragraph.
@param[in]  code_unit_index  The UTF-16 code unit index.

@return     The glyph information.
*/
func ParagraphCreateGlyphInfoAtCodeUnitIndexNew(paragraph Paragraph, code_unit_index uint64) GlyphInfo {
	var result GlyphInfo
	_, err := ffi.CallFunction(
		cifImpellerParagraphCreateGlyphInfoAtCodeUnitIndexNew,
		funcImpellerParagraphCreateGlyphInfoAtCodeUnitIndexNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph),
			unsafe.Pointer(&code_unit_index),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

// UNSUPPORTED ParagraphCreateGlyphInfoAtParagraphCoordinatesNew : param x is "double"

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  line_metrics  The line metrics.
*/
func LineMetricsRetain(line_metrics LineMetrics) {
	_, err := ffi.CallFunction(
		cifImpellerLineMetricsRetain,
		funcImpellerLineMetricsRetain,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&line_metrics),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  line_metrics  The line metrics.
*/
func LineMetricsRelease(line_metrics LineMetrics) {
	_, err := ffi.CallFunction(
		cifImpellerLineMetricsRelease,
		funcImpellerLineMetricsRelease,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&line_metrics),
		},
	)
	if err != nil {
		panic(err)
	}
}

// UNSUPPORTED LineMetricsGetUnscaledAscent : result type is "double"

// UNSUPPORTED LineMetricsGetAscent : result type is "double"

// UNSUPPORTED LineMetricsGetDescent : result type is "double"

// UNSUPPORTED LineMetricsGetBaseline : result type is "double"

/*
Used to determine if this line ends with an explicit line break
(e.g. '\n') or is the end of the paragraph.

@param[in]  metrics  The metrics.
@param[in]  line     The line index (zero based).

@return     True if the line is a hard break.
*/
func LineMetricsIsHardbreak(metrics LineMetrics, line uint64) Bool {
	var result Bool
	_, err := ffi.CallFunction(
		cifImpellerLineMetricsIsHardbreak,
		funcImpellerLineMetricsIsHardbreak,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&metrics),
			unsafe.Pointer(&line),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

// UNSUPPORTED LineMetricsGetWidth : result type is "double"

// UNSUPPORTED LineMetricsGetHeight : result type is "double"

// UNSUPPORTED LineMetricsGetLeft : result type is "double"

/*
Fetch the start index in the buffer of UTF-16 code units used to
represent the paragraph line.

@param[in]  metrics  The metrics.
@param[in]  line     The line index (zero based).

@return     The UTF-16 code units start index.
*/
func LineMetricsGetCodeUnitStartIndex(metrics LineMetrics, line uint64) uint64 {
	var result uint64
	_, err := ffi.CallFunction(
		cifImpellerLineMetricsGetCodeUnitStartIndex,
		funcImpellerLineMetricsGetCodeUnitStartIndex,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&metrics),
			unsafe.Pointer(&line),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
Fetch the end index in the buffer of UTF-16 code units used to
represent the paragraph line.

@param[in]  metrics  The metrics.
@param[in]  line     The line index (zero based).

@return     The UTF-16 code units end index.
*/
func LineMetricsGetCodeUnitEndIndex(metrics LineMetrics, line uint64) uint64 {
	var result uint64
	_, err := ffi.CallFunction(
		cifImpellerLineMetricsGetCodeUnitEndIndex,
		funcImpellerLineMetricsGetCodeUnitEndIndex,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&metrics),
			unsafe.Pointer(&line),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
Fetch the end index (excluding whitespace) in the buffer of
UTF-16 code units used to represent the paragraph line.

@param[in]  metrics  The metrics.
@param[in]  line     The line index (zero based).

@return     The UTF-16 code units end index excluding whitespace.
*/
func LineMetricsGetCodeUnitEndIndexExcludingWhitespace(metrics LineMetrics, line uint64) uint64 {
	var result uint64
	_, err := ffi.CallFunction(
		cifImpellerLineMetricsGetCodeUnitEndIndexExcludingWhitespace,
		funcImpellerLineMetricsGetCodeUnitEndIndexExcludingWhitespace,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&metrics),
			unsafe.Pointer(&line),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
Fetch the end index (including newlines) in the buffer of UTF-16
code units used to represent the paragraph line.

@param[in]  metrics  The metrics.
@param[in]  line     The line index (zero based).

@return     The UTF-16 code units end index including newlines.
*/
func LineMetricsGetCodeUnitEndIndexIncludingNewline(metrics LineMetrics, line uint64) uint64 {
	var result uint64
	_, err := ffi.CallFunction(
		cifImpellerLineMetricsGetCodeUnitEndIndexIncludingNewline,
		funcImpellerLineMetricsGetCodeUnitEndIndexIncludingNewline,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&metrics),
			unsafe.Pointer(&line),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  glyph_info  The glyph information.
*/
func GlyphInfoRetain(glyph_info GlyphInfo) {
	_, err := ffi.CallFunction(
		cifImpellerGlyphInfoRetain,
		funcImpellerGlyphInfoRetain,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&glyph_info),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  glyph_info  The glyph information.
*/
func GlyphInfoRelease(glyph_info GlyphInfo) {
	_, err := ffi.CallFunction(
		cifImpellerGlyphInfoRelease,
		funcImpellerGlyphInfoRelease,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&glyph_info),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Fetch the start index in the buffer of UTF-16 code units used to
represent the grapheme cluster for a glyph.

@param[in]  glyph_info  The glyph information.

@return     The UTF-16 code units start index.
*/
func GlyphInfoGetGraphemeClusterCodeUnitRangeBegin(glyph_info GlyphInfo) uint64 {
	var result uint64
	_, err := ffi.CallFunction(
		cifImpellerGlyphInfoGetGraphemeClusterCodeUnitRangeBegin,
		funcImpellerGlyphInfoGetGraphemeClusterCodeUnitRangeBegin,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&glyph_info),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
Fetch the end index in the buffer of UTF-16 code units used to
represent the grapheme cluster for a glyph.

@param[in]  glyph_info  The glyph information.

@return     The UTF-16 code units end index.
*/
func GlyphInfoGetGraphemeClusterCodeUnitRangeEnd(glyph_info GlyphInfo) uint64 {
	var result uint64
	_, err := ffi.CallFunction(
		cifImpellerGlyphInfoGetGraphemeClusterCodeUnitRangeEnd,
		funcImpellerGlyphInfoGetGraphemeClusterCodeUnitRangeEnd,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&glyph_info),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

// UNSUPPORTED GlyphInfoGetGraphemeClusterBounds : param out_bounds is "ImpellerRect *"

/*
@param[in]  glyph_info  The glyph information.

@return     True if the glyph represents an ellipsis. False otherwise.
*/
func GlyphInfoIsEllipsis(glyph_info GlyphInfo) Bool {
	var result Bool
	_, err := ffi.CallFunction(
		cifImpellerGlyphInfoIsEllipsis,
		funcImpellerGlyphInfoIsEllipsis,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&glyph_info),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

// UNSUPPORTED GlyphInfoGetTextDirection : result type is "ImpellerTextDirection"
