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

/*
Create a Vulkan context using the provided Vulkan Settings.

@param[in]  version   The version specified in the IMPELLER_VERSION macro.
@param[in]  settings  The Vulkan settings.

@return     The Vulkan context or NULL if one cannot be created.
*/
func ContextCreateVulkanNew(version uint32, settings ContextVulkanSettings) Context {
	var result Context
	_, err := ffi.CallFunction(
		cifImpellerContextCreateVulkanNew,
		funcImpellerContextCreateVulkanNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&version),
			unsafe.Pointer(new(&settings)),
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

/*
Get internal Vulkan handles managed by the given Vulkan context.
Ownership of the handles is still maintained by Impeller. This
accessor is just available so embedders can create resources
using the same device and instance as Impeller for interop.

@warning    If the context is not a Vulkan context, False is returned with
the [out] argument unaffected.

@param[in]  context          The context
@param[out]  out_vulkan_info  The out vulkan information

@return     If the Vulkan info could be fetched from the context.
*/
func ContextGetVulkanInfo(context Context, out_vulkan_info ContextVulkanInfo) Bool {
	var result Bool
	_, err := ffi.CallFunction(
		cifImpellerContextGetVulkanInfo,
		funcImpellerContextGetVulkanInfo,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&context),
			unsafe.Pointer(new(&out_vulkan_info)),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

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

/*
Create a new surface by wrapping an existing framebuffer object.
The framebuffer must be complete as determined by
`glCheckFramebufferStatus`. The framebuffer is still owned by
the caller and it must be collected once the surface is
collected.

@param[in]  context  The context.
@param[in]  fbo      The framebuffer object handle.
@param[in]  format   The format of the framebuffer.
@param[in]  size     The size of the framebuffer is texels.

@return     The surface if once can be created, NULL otherwise.
*/
func SurfaceCreateWrappedFBONew(context Context, fbo uint64, format PixelFormat, size ISize) Surface {
	var result Surface
	_, err := ffi.CallFunction(
		cifImpellerSurfaceCreateWrappedFBONew,
		funcImpellerSurfaceCreateWrappedFBONew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&context),
			unsafe.Pointer(&fbo),
			unsafe.Pointer(&format),
			unsafe.Pointer(new(&size)),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

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

/*
Get the bounds of the path.

The bounds are conservative. That is, they may be larger than
the actual shape of the path and could include the control
points and isolated calls to move the cursor.

@param[in]  path        The path
@param[out] out_bounds  The conservative bounds of the path.
*/
func PathGetBounds(path Path, out_bounds Rect) {
	_, err := ffi.CallFunction(
		cifImpellerPathGetBounds,
		funcImpellerPathGetBounds,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&path),
			unsafe.Pointer(new(&out_bounds)),
		},
	)
	if err != nil {
		panic(err)
	}
}

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

/*
Move the cursor to the specified location.

@param[in]  builder   The builder.
@param[in]  location  The location.
*/
func PathBuilderMoveTo(builder PathBuilder, location Point) {
	_, err := ffi.CallFunction(
		cifImpellerPathBuilderMoveTo,
		funcImpellerPathBuilderMoveTo,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(new(&location)),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Add a line segment from the current cursor location to the given
location. The cursor location is updated to be at the endpoint.

@param[in]  builder   The builder.
@param[in]  location  The location.
*/
func PathBuilderLineTo(builder PathBuilder, location Point) {
	_, err := ffi.CallFunction(
		cifImpellerPathBuilderLineTo,
		funcImpellerPathBuilderLineTo,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(new(&location)),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Add a quadratic curve from whose start point is the cursor to
the specified end point using the a single control point.

The new location of the cursor after this call is the end point.

@param[in]  builder        The builder.
@param[in]  control_point  The control point.
@param[in]  end_point      The end point.
*/
func PathBuilderQuadraticCurveTo(builder PathBuilder, control_point Point, end_point Point) {
	_, err := ffi.CallFunction(
		cifImpellerPathBuilderQuadraticCurveTo,
		funcImpellerPathBuilderQuadraticCurveTo,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(new(&control_point)),
			unsafe.Pointer(new(&end_point)),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Add a cubic curve whose start point is current cursor location
to the specified end point using the two specified control
points.

The new location of the cursor after this call is the end point
supplied.

@param[in]  builder          The builder
@param[in]  control_point_1  The control point 1
@param[in]  control_point_2  The control point 2
@param[in]  end_point        The end point
*/
func PathBuilderCubicCurveTo(builder PathBuilder, control_point_1 Point, control_point_2 Point, end_point Point) {
	_, err := ffi.CallFunction(
		cifImpellerPathBuilderCubicCurveTo,
		funcImpellerPathBuilderCubicCurveTo,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(new(&control_point_1)),
			unsafe.Pointer(new(&control_point_2)),
			unsafe.Pointer(new(&end_point)),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Adds a rectangle to the path.

@param[in]  builder  The builder.
@param[in]  rect     The rectangle.
*/
func PathBuilderAddRect(builder PathBuilder, rect Rect) {
	_, err := ffi.CallFunction(
		cifImpellerPathBuilderAddRect,
		funcImpellerPathBuilderAddRect,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(new(&rect)),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Add an arc to the path.

@param[in]  builder              The builder.
@param[in]  oval_bounds          The oval bounds.
@param[in]  start_angle_degrees  The start angle in degrees.
@param[in]  end_angle_degrees    The end angle in degrees.
*/
func PathBuilderAddArc(builder PathBuilder, oval_bounds Rect, start_angle_degrees float32, end_angle_degrees float32) {
	_, err := ffi.CallFunction(
		cifImpellerPathBuilderAddArc,
		funcImpellerPathBuilderAddArc,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(new(&oval_bounds)),
			unsafe.Pointer(&start_angle_degrees),
			unsafe.Pointer(&end_angle_degrees),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Add an oval to the path.

@param[in]  builder      The builder.
@param[in]  oval_bounds  The oval bounds.
*/
func PathBuilderAddOval(builder PathBuilder, oval_bounds Rect) {
	_, err := ffi.CallFunction(
		cifImpellerPathBuilderAddOval,
		funcImpellerPathBuilderAddOval,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(new(&oval_bounds)),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Add a rounded rect with potentially non-uniform radii to the
path.

@param[in]  builder         The builder.
@param[in]  rect            The rectangle.
@param[in]  rounding_radii  The rounding radii.
*/
func PathBuilderAddRoundedRect(builder PathBuilder, rect Rect, rounding_radii RoundingRadii) {
	_, err := ffi.CallFunction(
		cifImpellerPathBuilderAddRoundedRect,
		funcImpellerPathBuilderAddRoundedRect,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(new(&rect)),
			unsafe.Pointer(new(&rounding_radii)),
		},
	)
	if err != nil {
		panic(err)
	}
}

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

/*
Set the paint color.

@param[in]  paint  The paint.
@param[in]  color  The color.
*/
func PaintSetColor(paint Paint, color Color) {
	_, err := ffi.CallFunction(
		cifImpellerPaintSetColor,
		funcImpellerPaintSetColor,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paint),
			unsafe.Pointer(new(&color)),
		},
	)
	if err != nil {
		panic(err)
	}
}

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

// UNSUPPORTED TextureCreateWithContentsNew : param contents_on_release_user_data is "void *"

/*
Create a texture with an externally created OpenGL texture
handle.

Ownership of the handle is transferred over to Impeller after a
successful call to this method. Impeller is responsible for
calling glDeleteTextures on this handle. Do **not** collect this
handle yourself as this will lead to a double-free.

The handle must be created in the same context as the one used
by Impeller. If a different context is used, that context must
be in the same sharegroup as Impellers OpenGL context and all
synchronization of texture contents must already be complete.

If the context is not an OpenGL context, this call will always
fail.

@param[in]  context     The context
@param[in]  descriptor  The descriptor
@param[in]  handle      The handle

@return     The texture if one could be created by adopting the supplied
texture handle, NULL otherwise.
*/
func TextureCreateWithOpenGLTextureHandleNew(context Context, descriptor TextureDescriptor, handle uint64) Texture {
	var result Texture
	_, err := ffi.CallFunction(
		cifImpellerTextureCreateWithOpenGLTextureHandleNew,
		funcImpellerTextureCreateWithOpenGLTextureHandleNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&context),
			unsafe.Pointer(new(&descriptor)),
			unsafe.Pointer(&handle),
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

// UNSUPPORTED FragmentProgramNew : param data_release_user_data is "void *"

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

// UNSUPPORTED ColorSourceCreateLinearGradientNew : param stops is "const float *"

// UNSUPPORTED ColorSourceCreateRadialGradientNew : param stops is "const float *"

// UNSUPPORTED ColorSourceCreateConicalGradientNew : param stops is "const float *"

// UNSUPPORTED ColorSourceCreateSweepGradientNew : param stops is "const float *"

/*
Create a color source that samples from an image.

@param[in]  image                 The image.
@param[in]  horizontal_tile_mode  The horizontal tile mode.
@param[in]  vertical_tile_mode    The vertical tile mode.
@param[in]  sampling              The sampling.
@param[in]  transformation        The transformation.

@return     The color source.
*/
func ColorSourceCreateImageNew(image Texture, horizontal_tile_mode TileMode, vertical_tile_mode TileMode, sampling TextureSampling, transformation Matrix) ColorSource {
	var result ColorSource
	_, err := ffi.CallFunction(
		cifImpellerColorSourceCreateImageNew,
		funcImpellerColorSourceCreateImageNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&image),
			unsafe.Pointer(&horizontal_tile_mode),
			unsafe.Pointer(&vertical_tile_mode),
			unsafe.Pointer(&sampling),
			unsafe.Pointer(new(&transformation)),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

// UNSUPPORTED ColorSourceCreateFragmentProgramNew : param data is "const uint8_t *"

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

/*
Create a color filter that performs blending of pixel values
independently.

@param[in]  color       The color.
@param[in]  blend_mode  The blend mode.

@return     The color filter.
*/
func ColorFilterCreateBlendNew(color Color, blend_mode BlendMode) ColorFilter {
	var result ColorFilter
	_, err := ffi.CallFunction(
		cifImpellerColorFilterCreateBlendNew,
		funcImpellerColorFilterCreateBlendNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(new(&color)),
			unsafe.Pointer(&blend_mode),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
Create a color filter that transforms pixel color values
independently.

@param[in]  color_matrix  The color matrix.

@return     The color filter.
*/
func ColorFilterCreateColorMatrixNew(color_matrix ColorMatrix) ColorFilter {
	var result ColorFilter
	_, err := ffi.CallFunction(
		cifImpellerColorFilterCreateColorMatrixNew,
		funcImpellerColorFilterCreateColorMatrixNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(new(&color_matrix)),
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

/*
Creates an image filter that applies a transformation matrix to
the underlying image.

@param[in]  matrix    The transformation matrix.
@param[in]  sampling  The image sampling mode.

@return     The image filter.
*/
func ImageFilterCreateMatrixNew(matrix Matrix, sampling TextureSampling) ImageFilter {
	var result ImageFilter
	_, err := ffi.CallFunction(
		cifImpellerImageFilterCreateMatrixNew,
		funcImpellerImageFilterCreateMatrixNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(new(&matrix)),
			unsafe.Pointer(&sampling),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

// UNSUPPORTED ImageFilterCreateFragmentProgramNew : param data is "const uint8_t *"

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

/*
Create a new display list builder.

An optional cull rectangle may be specified. Impeller is allowed
to treat the contents outside this rectangle as being undefined.
This may aid performance optimizations.

@param[in]  cull_rect  The cull rectangle or NULL.

@return     The display list builder.
*/
func DisplayListBuilderNew(cull_rect Rect) DisplayListBuilder {
	var result DisplayListBuilder
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderNew,
		funcImpellerDisplayListBuilderNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(new(&cull_rect)),
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

/*
Stashes the current transformation and clip state onto a save
stack and creates and creates an offscreen layer onto which
subsequent rendering intent will be directed to.

On the balancing call to restore, the supplied paints filters
and blend modes will be used to composite the offscreen contents
back onto the display display list.

@param[in]  builder   The builder.
@param[in]  bounds    The bounds.
@param[in]  paint     The paint.
@param[in]  backdrop  The backdrop.
*/
func DisplayListBuilderSaveLayer(builder DisplayListBuilder, bounds Rect, paint Paint, backdrop ImageFilter) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderSaveLayer,
		funcImpellerDisplayListBuilderSaveLayer,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(new(&bounds)),
			unsafe.Pointer(&paint),
			unsafe.Pointer(&backdrop),
		},
	)
	if err != nil {
		panic(err)
	}
}

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

/*
Appends the the provided transformation to the transformation
already on the save stack.

@param[in]  builder    The builder.
@param[in]  transform  The transform to append.
*/
func DisplayListBuilderTransform(builder DisplayListBuilder, transform Matrix) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderTransform,
		funcImpellerDisplayListBuilderTransform,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(new(&transform)),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Clear the transformation on top of the save stack and replace it
with a new value.

@param[in]  builder    The builder.
@param[in]  transform  The new transform.
*/
func DisplayListBuilderSetTransform(builder DisplayListBuilder, transform Matrix) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderSetTransform,
		funcImpellerDisplayListBuilderSetTransform,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(new(&transform)),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Get the transformation currently built up on the top of the
transformation stack.

@param[in]  builder        The builder.
@param[out] out_transform  The transform.
*/
func DisplayListBuilderGetTransform(builder DisplayListBuilder, out_transform Matrix) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderGetTransform,
		funcImpellerDisplayListBuilderGetTransform,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(new(&out_transform)),
		},
	)
	if err != nil {
		panic(err)
	}
}

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

/*
Reduces the clip region to the intersection of the current clip
and the given rectangle taking into account the clip operation.

@param[in]  builder  The builder.
@param[in]  rect     The rectangle.
@param[in]  op       The operation.
*/
func DisplayListBuilderClipRect(builder DisplayListBuilder, rect Rect, op ClipOperation) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderClipRect,
		funcImpellerDisplayListBuilderClipRect,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(new(&rect)),
			unsafe.Pointer(&op),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Reduces the clip region to the intersection of the current clip
and the given oval taking into account the clip operation.

@param[in]  builder      The builder.
@param[in]  oval_bounds  The oval bounds.
@param[in]  op           The operation.
*/
func DisplayListBuilderClipOval(builder DisplayListBuilder, oval_bounds Rect, op ClipOperation) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderClipOval,
		funcImpellerDisplayListBuilderClipOval,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(new(&oval_bounds)),
			unsafe.Pointer(&op),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Reduces the clip region to the intersection of the current clip
and the given rounded rectangle taking into account the clip
operation.

@param[in]  builder  The builder.
@param[in]  rect     The rectangle.
@param[in]  radii    The radii.
@param[in]  op       The operation.
*/
func DisplayListBuilderClipRoundedRect(builder DisplayListBuilder, rect Rect, radii RoundingRadii, op ClipOperation) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderClipRoundedRect,
		funcImpellerDisplayListBuilderClipRoundedRect,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(new(&rect)),
			unsafe.Pointer(new(&radii)),
			unsafe.Pointer(&op),
		},
	)
	if err != nil {
		panic(err)
	}
}

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

/*
Draws a line segment.

@param[in]  builder  The builder.
@param[in]  from     The starting point of the line.
@param[in]  to       The end point of the line.
@param[in]  paint    The paint.
*/
func DisplayListBuilderDrawLine(builder DisplayListBuilder, from Point, to Point, paint Paint) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderDrawLine,
		funcImpellerDisplayListBuilderDrawLine,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(new(&from)),
			unsafe.Pointer(new(&to)),
			unsafe.Pointer(&paint),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Draws a dash line segment.

@param[in]  builder     The builder.
@param[in]  from        The starting point of the line.
@param[in]  to          The end point of the line.
@param[in]  on_length   On length.
@param[in]  off_length  Off length.
@param[in]  paint       The paint.
*/
func DisplayListBuilderDrawDashedLine(builder DisplayListBuilder, from Point, to Point, on_length float32, off_length float32, paint Paint) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderDrawDashedLine,
		funcImpellerDisplayListBuilderDrawDashedLine,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(new(&from)),
			unsafe.Pointer(new(&to)),
			unsafe.Pointer(&on_length),
			unsafe.Pointer(&off_length),
			unsafe.Pointer(&paint),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Draws a rectangle.

@param[in]  builder  The builder.
@param[in]  rect     The rectangle.
@param[in]  paint    The paint.
*/
func DisplayListBuilderDrawRect(builder DisplayListBuilder, rect Rect, paint Paint) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderDrawRect,
		funcImpellerDisplayListBuilderDrawRect,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(new(&rect)),
			unsafe.Pointer(&paint),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Draws an oval.

@param[in]  builder      The builder.
@param[in]  oval_bounds  The oval bounds.
@param[in]  paint        The paint.
*/
func DisplayListBuilderDrawOval(builder DisplayListBuilder, oval_bounds Rect, paint Paint) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderDrawOval,
		funcImpellerDisplayListBuilderDrawOval,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(new(&oval_bounds)),
			unsafe.Pointer(&paint),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Draws a rounded rect.

@param[in]  builder  The builder.
@param[in]  rect     The rectangle.
@param[in]  radii    The radii.
@param[in]  paint    The paint.
*/
func DisplayListBuilderDrawRoundedRect(builder DisplayListBuilder, rect Rect, radii RoundingRadii, paint Paint) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderDrawRoundedRect,
		funcImpellerDisplayListBuilderDrawRoundedRect,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(new(&rect)),
			unsafe.Pointer(new(&radii)),
			unsafe.Pointer(&paint),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Draws a shape that is the different between the specified
rectangles (each with configurable corner radii).

@param[in]  builder      The builder.
@param[in]  outer_rect   The outer rectangle.
@param[in]  outer_radii  The outer radii.
@param[in]  inner_rect   The inner rectangle.
@param[in]  inner_radii  The inner radii.
@param[in]  paint        The paint.
*/
func DisplayListBuilderDrawRoundedRectDifference(builder DisplayListBuilder, outer_rect Rect, outer_radii RoundingRadii, inner_rect Rect, inner_radii RoundingRadii, paint Paint) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderDrawRoundedRectDifference,
		funcImpellerDisplayListBuilderDrawRoundedRectDifference,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(new(&outer_rect)),
			unsafe.Pointer(new(&outer_radii)),
			unsafe.Pointer(new(&inner_rect)),
			unsafe.Pointer(new(&inner_radii)),
			unsafe.Pointer(&paint),
		},
	)
	if err != nil {
		panic(err)
	}
}

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

/*
Draw a paragraph at the specified point.

@param[in]  builder    The builder.
@param[in]  paragraph  The paragraph.
@param[in]  point      The point.
*/
func DisplayListBuilderDrawParagraph(builder DisplayListBuilder, paragraph Paragraph, point Point) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderDrawParagraph,
		funcImpellerDisplayListBuilderDrawParagraph,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&paragraph),
			unsafe.Pointer(new(&point)),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Draw a shadow for a Path given a material elevation. If the
occluding object is not opaque, additional hints (via the
`occluder_is_transparent` argument) must be provided to render
the shadow correctly.

@param[in]  builder    The builder.
@param[in]  path       The shadow path.
@param[in]  color      The shadow color.
@param[in]  elevation  The material elevation.
@param[in]  occluder_is_transparent
If the object casting the shadow is transparent.
@param[in]  device_pixel_ratio
The device pixel ratio.
*/
func DisplayListBuilderDrawShadow(builder DisplayListBuilder, path Path, color Color, elevation float32, occluder_is_transparent Bool, device_pixel_ratio float32) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderDrawShadow,
		funcImpellerDisplayListBuilderDrawShadow,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&path),
			unsafe.Pointer(new(&color)),
			unsafe.Pointer(&elevation),
			unsafe.Pointer(&occluder_is_transparent),
			unsafe.Pointer(&device_pixel_ratio),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Draw a texture at the specified point.

@param[in]  builder   The builder.
@param[in]  texture   The texture.
@param[in]  point     The point.
@param[in]  sampling  The sampling.
@param[in]  paint     The paint.
*/
func DisplayListBuilderDrawTexture(builder DisplayListBuilder, texture Texture, point Point, sampling TextureSampling, paint Paint) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderDrawTexture,
		funcImpellerDisplayListBuilderDrawTexture,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&texture),
			unsafe.Pointer(new(&point)),
			unsafe.Pointer(&sampling),
			unsafe.Pointer(&paint),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Draw a portion of texture at the specified location.

@param[in]  builder   The builder.
@param[in]  texture   The texture.
@param[in]  src_rect  The source rectangle.
@param[in]  dst_rect  The destination rectangle.
@param[in]  sampling  The sampling.
@param[in]  paint     The paint.
*/
func DisplayListBuilderDrawTextureRect(builder DisplayListBuilder, texture Texture, src_rect Rect, dst_rect Rect, sampling TextureSampling, paint Paint) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderDrawTextureRect,
		funcImpellerDisplayListBuilderDrawTextureRect,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&texture),
			unsafe.Pointer(new(&src_rect)),
			unsafe.Pointer(new(&dst_rect)),
			unsafe.Pointer(&sampling),
			unsafe.Pointer(&paint),
		},
	)
	if err != nil {
		panic(err)
	}
}

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

// UNSUPPORTED TypographyContextRegisterFont : param contents_on_release_user_data is "void *"

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

/*
Set one of more text decorations on the paragraph. Decorations
can be underlines, overlines, strikethroughs, etc.. The style of
decorations can be set as well (dashed, dotted, wavy, etc..)

@param[in]  ImpellerParagraphStyle  The paragraph style.
@param[in]  decoration              The text decoration.
*/
func ParagraphStyleSetTextDecoration(paragraph_style ParagraphStyle, decoration TextDecoration) {
	_, err := ffi.CallFunction(
		cifImpellerParagraphStyleSetTextDecoration,
		funcImpellerParagraphStyleSetTextDecoration,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph_style),
			unsafe.Pointer(new(&decoration)),
		},
	)
	if err != nil {
		panic(err)
	}
}

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

/*
Get the range into the UTF-16 code unit buffer that represents
the word at the specified caret location in the same buffer.

Word boundaries are defined more precisely in [Unicode Standard
Annex #29](http://www.unicode.org/reports/tr29/#Word_Boundaries)

@param[in]  paragraph        The paragraph
@param[in]  code_unit_index  The code unit index
@param[out]  code_unit_index The range.
*/
func ParagraphGetWordBoundary(paragraph Paragraph, code_unit_index uint64, out_range Range) {
	_, err := ffi.CallFunction(
		cifImpellerParagraphGetWordBoundary,
		funcImpellerParagraphGetWordBoundary,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph),
			unsafe.Pointer(&code_unit_index),
			unsafe.Pointer(new(&out_range)),
		},
	)
	if err != nil {
		panic(err)
	}
}

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

/*
Create a new instance of glyph info that can be queried for
information about the glyph closest to the specified coordinates
relative to the origin of the paragraph. The instance must be
freed using `ImpellerGlyphInfoRelease`.

@param[in]  paragraph  The paragraph.
@param[in]  x          The x coordinate relative to paragraph origin.
@param[in]  y          The x coordinate relative to paragraph origin.

@return     The glyph information.
*/
func ParagraphCreateGlyphInfoAtParagraphCoordinatesNew(paragraph Paragraph, x float64, y float64) GlyphInfo {
	var result GlyphInfo
	_, err := ffi.CallFunction(
		cifImpellerParagraphCreateGlyphInfoAtParagraphCoordinatesNew,
		funcImpellerParagraphCreateGlyphInfoAtParagraphCoordinatesNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph),
			unsafe.Pointer(&x),
			unsafe.Pointer(&y),
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

/*
The rise from the baseline as calculated from the font and style
for this line ignoring the height from the text style.

@param[in]  metrics  The metrics.
@param[in]  line     The line index (zero based).

@return     The unscaled ascent.
*/
func LineMetricsGetUnscaledAscent(metrics LineMetrics, line uint64) float64 {
	var result float64
	_, err := ffi.CallFunction(
		cifImpellerLineMetricsGetUnscaledAscent,
		funcImpellerLineMetricsGetUnscaledAscent,
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
The rise from the baseline as calculated from the font and style
for this line.

@param[in]  metrics  The metrics.
@param[in]  line     The line index (zero based).

@return     The ascent.
*/
func LineMetricsGetAscent(metrics LineMetrics, line uint64) float64 {
	var result float64
	_, err := ffi.CallFunction(
		cifImpellerLineMetricsGetAscent,
		funcImpellerLineMetricsGetAscent,
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
The drop from the baseline as calculated from the font and style
for this line.

@param[in]  metrics  The metrics.
@param[in]  line     The line index (zero based).

@return     The descent.
*/
func LineMetricsGetDescent(metrics LineMetrics, line uint64) float64 {
	var result float64
	_, err := ffi.CallFunction(
		cifImpellerLineMetricsGetDescent,
		funcImpellerLineMetricsGetDescent,
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
The y coordinate of the baseline for this line from the top of
the paragraph.

@param[in]  metrics  The metrics.
@param[in]  line     The line index (zero based).

@return     The baseline.
*/
func LineMetricsGetBaseline(metrics LineMetrics, line uint64) float64 {
	var result float64
	_, err := ffi.CallFunction(
		cifImpellerLineMetricsGetBaseline,
		funcImpellerLineMetricsGetBaseline,
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

/*
Width of the line from the left edge of the leftmost glyph to
the right edge of the rightmost glyph.

@param[in]  metrics  The metrics.
@param[in]  line     The line index (zero based).

@return     The width.
*/
func LineMetricsGetWidth(metrics LineMetrics, line uint64) float64 {
	var result float64
	_, err := ffi.CallFunction(
		cifImpellerLineMetricsGetWidth,
		funcImpellerLineMetricsGetWidth,
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
Total height of the line from the top edge to the bottom edge.

@param[in]  metrics  The metrics.
@param[in]  line     The line index (zero based).

@return     The height.
*/
func LineMetricsGetHeight(metrics LineMetrics, line uint64) float64 {
	var result float64
	_, err := ffi.CallFunction(
		cifImpellerLineMetricsGetHeight,
		funcImpellerLineMetricsGetHeight,
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
The x coordinate of left edge of the line.

@param[in]  metrics  The metrics.
@param[in]  line     The line index (zero based).

@return     The left edge coordinate.
*/
func LineMetricsGetLeft(metrics LineMetrics, line uint64) float64 {
	var result float64
	_, err := ffi.CallFunction(
		cifImpellerLineMetricsGetLeft,
		funcImpellerLineMetricsGetLeft,
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

/*
Fetch the bounds of the grapheme cluster for the glyph in the
coordinate space of the paragraph.

@param[in]  glyph_info  The glyph information.
@param[out] out_bounds  The grapheme cluster bounds.
*/
func GlyphInfoGetGraphemeClusterBounds(glyph_info GlyphInfo, out_bounds Rect) {
	_, err := ffi.CallFunction(
		cifImpellerGlyphInfoGetGraphemeClusterBounds,
		funcImpellerGlyphInfoGetGraphemeClusterBounds,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&glyph_info),
			unsafe.Pointer(new(&out_bounds)),
		},
	)
	if err != nil {
		panic(err)
	}
}

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

/*
@param[in]  glyph_info  The glyph information.

@return     The direction of the run that contains the glyph.
*/
func GlyphInfoGetTextDirection(glyph_info GlyphInfo) TextDirection {
	var result TextDirection
	_, err := ffi.CallFunction(
		cifImpellerGlyphInfoGetTextDirection,
		funcImpellerGlyphInfoGetTextDirection,
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
