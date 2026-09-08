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

/*
Create an OpenGL(ES) Impeller context.

@warning    Unlike other context types, the OpenGL ES context can only be
created, used, and collected on the calling thread. This
restriction may be lifted in the future once reactor workers are
exposed in the API. No other context types have threading
restrictions. Till reactor workers can be used, using the
context on a background thread will cause a stall of OpenGL
operations.

@param[in]  version      The version of the Impeller
standalone API. See `ImpellerGetVersion`. If the
specified here is not compatible with the version
of the library, context creation will fail and NULL
context returned from this call.
@param[in]  gl_proc_address_callback
The gl proc address callback. For instance,
`eglGetProcAddress`.
@param[in]  gl_proc_address_callback_user_data
The gl proc address callback user data baton. This
pointer is not interpreted by Impeller and will be
returned as user data in the proc address callback.
user data.

@return     The context or NULL if one cannot be created.
*/
func ContextCreateOpenGLESNew(version uint32, gl_proc_address_callback ProcAddressCallback, gl_proc_address_callback_user_data unsafe.Pointer) Context {
	var result Context
	_, err := ffi.CallFunction(
		cifImpellerContextCreateOpenGLESNew,
		funcImpellerContextCreateOpenGLESNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&version),
			unsafe.Pointer(&gl_proc_address_callback),
			unsafe.Pointer(&gl_proc_address_callback_user_data),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

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
func ContextCreateVulkanNew(version uint32, settings *ContextVulkanSettings) Context {
	var result Context
	_, err := ffi.CallFunction(
		cifImpellerContextCreateVulkanNew,
		funcImpellerContextCreateVulkanNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&version),
			unsafe.Pointer(&settings),
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
func (context Context) Retain() {
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
func (context Context) Release() {
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
func (context Context) GetVulkanInfo(out_vulkan_info *ContextVulkanInfo) Bool {
	var result Bool
	_, err := ffi.CallFunction(
		cifImpellerContextGetVulkanInfo,
		funcImpellerContextGetVulkanInfo,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&context),
			unsafe.Pointer(&out_vulkan_info),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
Create a new Vulkan swapchain using a VkSurfaceKHR instance.
Ownership of the surface is transferred over to Impeller. The
Vulkan instance the surface is created from must the same as the
context provided.

@param[in]  context             The context. Must be a Vulkan context whose
instance is the same used to create the
surface passed into the next argument.
@param      vulkan_surface_khr  The vulkan surface.

@return     The vulkan swapchain.
*/
func (context Context) VulkanSwapchainCreateNew(vulkan_surface_khr unsafe.Pointer) VulkanSwapchain {
	var result VulkanSwapchain
	_, err := ffi.CallFunction(
		cifImpellerVulkanSwapchainCreateNew,
		funcImpellerVulkanSwapchainCreateNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&context),
			unsafe.Pointer(&vulkan_surface_khr),
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

@param[in]  swapchain  The swapchain.
*/
func (swapchain VulkanSwapchain) Retain() {
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
func (swapchain VulkanSwapchain) Release() {
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
func (swapchain VulkanSwapchain) AcquireNextSurfaceNew() Surface {
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
func (context Context) SurfaceCreateWrappedFBONew(fbo uint64, format PixelFormat, size *ISize) Surface {
	var result Surface
	_, err := ffi.CallFunction(
		cifImpellerSurfaceCreateWrappedFBONew,
		funcImpellerSurfaceCreateWrappedFBONew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&context),
			unsafe.Pointer(&fbo),
			unsafe.Pointer(&format),
			unsafe.Pointer(&size),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
Create a surface by wrapping a Metal drawable. This is useful
during WSI when the drawable is the backing store of the Metal
layer being drawn to.

The Metal layer must be using the same device managed by the
underlying context.

@param[in]  context         The context. The Metal device managed by this
context must be the same used to create the
drawable that is being wrapped.
@param      metal_drawable  The drawable to wrap as a surface.

@return     The surface if one could be wrapped, NULL otherwise.
*/
func (context Context) SurfaceCreateWrappedMetalDrawableNew(metal_drawable unsafe.Pointer) Surface {
	var result Surface
	_, err := ffi.CallFunction(
		cifImpellerSurfaceCreateWrappedMetalDrawableNew,
		funcImpellerSurfaceCreateWrappedMetalDrawableNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&context),
			unsafe.Pointer(&metal_drawable),
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

@param[in]  surface  The surface.
*/
func (surface Surface) Retain() {
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
func (surface Surface) Release() {
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
func (surface Surface) DrawDisplayList(display_list DisplayList) Bool {
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
func (surface Surface) Present() Bool {
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
func (path Path) Retain() {
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
func (path Path) Release() {
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
func (path Path) GetBounds(out_bounds *Rect) {
	_, err := ffi.CallFunction(
		cifImpellerPathGetBounds,
		funcImpellerPathGetBounds,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&path),
			unsafe.Pointer(&out_bounds),
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
func (builder PathBuilder) Retain() {
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
func (builder PathBuilder) Release() {
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
func (builder PathBuilder) MoveTo(location *Point) {
	_, err := ffi.CallFunction(
		cifImpellerPathBuilderMoveTo,
		funcImpellerPathBuilderMoveTo,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&location),
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
func (builder PathBuilder) LineTo(location *Point) {
	_, err := ffi.CallFunction(
		cifImpellerPathBuilderLineTo,
		funcImpellerPathBuilderLineTo,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&location),
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
func (builder PathBuilder) QuadraticCurveTo(control_point *Point, end_point *Point) {
	_, err := ffi.CallFunction(
		cifImpellerPathBuilderQuadraticCurveTo,
		funcImpellerPathBuilderQuadraticCurveTo,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&control_point),
			unsafe.Pointer(&end_point),
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
func (builder PathBuilder) CubicCurveTo(control_point_1 *Point, control_point_2 *Point, end_point *Point) {
	_, err := ffi.CallFunction(
		cifImpellerPathBuilderCubicCurveTo,
		funcImpellerPathBuilderCubicCurveTo,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&control_point_1),
			unsafe.Pointer(&control_point_2),
			unsafe.Pointer(&end_point),
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
func (builder PathBuilder) AddRect(rect *Rect) {
	_, err := ffi.CallFunction(
		cifImpellerPathBuilderAddRect,
		funcImpellerPathBuilderAddRect,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&rect),
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
func (builder PathBuilder) AddArc(oval_bounds *Rect, start_angle_degrees float32, end_angle_degrees float32) {
	_, err := ffi.CallFunction(
		cifImpellerPathBuilderAddArc,
		funcImpellerPathBuilderAddArc,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&oval_bounds),
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
func (builder PathBuilder) AddOval(oval_bounds *Rect) {
	_, err := ffi.CallFunction(
		cifImpellerPathBuilderAddOval,
		funcImpellerPathBuilderAddOval,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&oval_bounds),
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
func (builder PathBuilder) AddRoundedRect(rect *Rect, rounding_radii *RoundingRadii) {
	_, err := ffi.CallFunction(
		cifImpellerPathBuilderAddRoundedRect,
		funcImpellerPathBuilderAddRoundedRect,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&rect),
			unsafe.Pointer(&rounding_radii),
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
func (builder PathBuilder) Close() {
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
func (builder PathBuilder) CopyPathNew(fill FillType) Path {
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
func (builder PathBuilder) TakePathNew(fill FillType) Path {
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
func (paint Paint) Retain() {
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
func (paint Paint) Release() {
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
func (paint Paint) SetColor(color *Color) {
	_, err := ffi.CallFunction(
		cifImpellerPaintSetColor,
		funcImpellerPaintSetColor,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paint),
			unsafe.Pointer(&color),
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
func (paint Paint) SetBlendMode(mode BlendMode) {
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
func (paint Paint) SetDrawStyle(style DrawStyle) {
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
func (paint Paint) SetStrokeCap(cap StrokeCap) {
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
func (paint Paint) SetStrokeJoin(join StrokeJoin) {
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
func (paint Paint) SetStrokeWidth(width float32) {
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
func (paint Paint) SetStrokeMiter(miter float32) {
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
func (paint Paint) SetColorFilter(color_filter ColorFilter) {
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
func (paint Paint) SetColorSource(color_source ColorSource) {
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
func (paint Paint) SetImageFilter(image_filter ImageFilter) {
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
func (paint Paint) SetMaskFilter(mask_filter MaskFilter) {
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

/*
Create a texture with decompressed bytes.

Impeller will do its best to perform the transfer of this data
to GPU memory with a minimal number of copies. Towards this
end, it may need to send this data to a different thread for
preparation and transfer. To facilitate this transfer, it is
recommended that the content mapping have a release callback
attach to it. When there is a release callback, Impeller assumes
that collection of the data can be deferred till texture upload
is done and can happen on a background thread. When there is no
release callback, Impeller may try to perform an eager copy of
the data if it needs to perform data preparation and transfer on
a background thread.

Whether an extra data copy actually occurs will always depend on
the rendering backend in use. But it is best practice to provide
a release callback and be resilient to the data being released
in a deferred manner on a background thread.

@warning    Do **not** supply compressed image data directly (PNG, JPEG,
etc...). This function only works with tightly packed
decompressed data.

@param[in]  context                        The context.
@param[in]  descriptor                     The texture descriptor.
@param[in]  contents                       The contents.
@param[in]  contents_on_release_user_data  The baton passes to the contents
release callback if one exists.

@return     The texture if one can be created using the provided data, NULL
otherwise.
*/
func (context Context) TextureCreateWithContentsNew(descriptor *TextureDescriptor, contents *Mapping, contents_on_release_user_data unsafe.Pointer) Texture {
	var result Texture
	_, err := ffi.CallFunction(
		cifImpellerTextureCreateWithContentsNew,
		funcImpellerTextureCreateWithContentsNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&context),
			unsafe.Pointer(&descriptor),
			unsafe.Pointer(&contents),
			unsafe.Pointer(&contents_on_release_user_data),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

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
func (context Context) TextureCreateWithOpenGLTextureHandleNew(descriptor *TextureDescriptor, handle uint64) Texture {
	var result Texture
	_, err := ffi.CallFunction(
		cifImpellerTextureCreateWithOpenGLTextureHandleNew,
		funcImpellerTextureCreateWithOpenGLTextureHandleNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&context),
			unsafe.Pointer(&descriptor),
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
func (texture Texture) Retain() {
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
func (texture Texture) Release() {
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
func (texture Texture) GetOpenGLHandle() uint64 {
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

/*
Create a new fragment program using data obtained by compiling a
GLSL shader with `impellerc`.

@warning    The data provided must be compiled by `impellerc`. Providing raw
GLSL strings will lead to a `nullptr` return. Impeller does not
compile shaders at runtime.

@param[in]  data                    The data compiled by `impellerc`.
@param      data_release_user_data  A baton passed back to the caller on the
invocation of the mappings release
callback. This call can happen on any
thread.

@return     The fragment program if one can be created, nullptr otherwise.
*/
func FragmentProgramNew(data *Mapping, data_release_user_data unsafe.Pointer) FragmentProgram {
	var result FragmentProgram
	_, err := ffi.CallFunction(
		cifImpellerFragmentProgramNew,
		funcImpellerFragmentProgramNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&data),
			unsafe.Pointer(&data_release_user_data),
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

@param[in]  fragment_program  The fragment program.
*/
func (fragment_program FragmentProgram) Retain() {
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
func (fragment_program FragmentProgram) Release() {
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
func (color_source ColorSource) Retain() {
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
func (color_source ColorSource) Release() {
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

/*
Create a color source that forms a linear gradient.

@param[in]  start_point     The start point.
@param[in]  end_point       The end point.
@param[in]  stop_count      The stop count.
@param[in]  colors          The colors.
@param[in]  stops           The stops.
@param[in]  tile_mode       The tile mode.
@param[in]  transformation  The transformation.

@return     The color source.
*/
func ColorSourceCreateLinearGradientNew(start_point *Point, end_point *Point, stop_count uint32, colors *Color, stops *float32, tile_mode TileMode, transformation *Matrix) ColorSource {
	var result ColorSource
	_, err := ffi.CallFunction(
		cifImpellerColorSourceCreateLinearGradientNew,
		funcImpellerColorSourceCreateLinearGradientNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&start_point),
			unsafe.Pointer(&end_point),
			unsafe.Pointer(&stop_count),
			unsafe.Pointer(&colors),
			unsafe.Pointer(&stops),
			unsafe.Pointer(&tile_mode),
			unsafe.Pointer(&transformation),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
Create a color source that forms a radial gradient.

@param[in]  center          The center.
@param[in]  radius          The radius.
@param[in]  stop_count      The stop count.
@param[in]  colors          The colors.
@param[in]  stops           The stops.
@param[in]  tile_mode       The tile mode.
@param[in]  transformation  The transformation.

@return     The color source.
*/
func ColorSourceCreateRadialGradientNew(center *Point, radius float32, stop_count uint32, colors *Color, stops *float32, tile_mode TileMode, transformation *Matrix) ColorSource {
	var result ColorSource
	_, err := ffi.CallFunction(
		cifImpellerColorSourceCreateRadialGradientNew,
		funcImpellerColorSourceCreateRadialGradientNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&center),
			unsafe.Pointer(&radius),
			unsafe.Pointer(&stop_count),
			unsafe.Pointer(&colors),
			unsafe.Pointer(&stops),
			unsafe.Pointer(&tile_mode),
			unsafe.Pointer(&transformation),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
Create a color source that forms a conical gradient.

@param[in]  start_center    The start center.
@param[in]  start_radius    The start radius.
@param[in]  end_center      The end center.
@param[in]  end_radius      The end radius.
@param[in]  stop_count      The stop count.
@param[in]  colors          The colors.
@param[in]  stops           The stops.
@param[in]  tile_mode       The tile mode.
@param[in]  transformation  The transformation.

@return     The color source.
*/
func ColorSourceCreateConicalGradientNew(start_center *Point, start_radius float32, end_center *Point, end_radius float32, stop_count uint32, colors *Color, stops *float32, tile_mode TileMode, transformation *Matrix) ColorSource {
	var result ColorSource
	_, err := ffi.CallFunction(
		cifImpellerColorSourceCreateConicalGradientNew,
		funcImpellerColorSourceCreateConicalGradientNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&start_center),
			unsafe.Pointer(&start_radius),
			unsafe.Pointer(&end_center),
			unsafe.Pointer(&end_radius),
			unsafe.Pointer(&stop_count),
			unsafe.Pointer(&colors),
			unsafe.Pointer(&stops),
			unsafe.Pointer(&tile_mode),
			unsafe.Pointer(&transformation),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
Create a color source that forms a sweep gradient.

@param[in]  center          The center.
@param[in]  start           The start.
@param[in]  end             The end.
@param[in]  stop_count      The stop count.
@param[in]  colors          The colors.
@param[in]  stops           The stops.
@param[in]  tile_mode       The tile mode.
@param[in]  transformation  The transformation.

@return     The color source.
*/
func ColorSourceCreateSweepGradientNew(center *Point, start float32, end float32, stop_count uint32, colors *Color, stops *float32, tile_mode TileMode, transformation *Matrix) ColorSource {
	var result ColorSource
	_, err := ffi.CallFunction(
		cifImpellerColorSourceCreateSweepGradientNew,
		funcImpellerColorSourceCreateSweepGradientNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&center),
			unsafe.Pointer(&start),
			unsafe.Pointer(&end),
			unsafe.Pointer(&stop_count),
			unsafe.Pointer(&colors),
			unsafe.Pointer(&stops),
			unsafe.Pointer(&tile_mode),
			unsafe.Pointer(&transformation),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
Create a color source that samples from an image.

@param[in]  image                 The image.
@param[in]  horizontal_tile_mode  The horizontal tile mode.
@param[in]  vertical_tile_mode    The vertical tile mode.
@param[in]  sampling              The sampling.
@param[in]  transformation        The transformation.

@return     The color source.
*/
func (image Texture) ColorSourceCreateImageNew(horizontal_tile_mode TileMode, vertical_tile_mode TileMode, sampling TextureSampling, transformation *Matrix) ColorSource {
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
			unsafe.Pointer(&transformation),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
Create a color source whose pixels are shaded by a fragment
program.

@see        https://docs.flutter.dev/ui/design/graphics/fragment-shaders

@param[in]  context            The context.
@param[in]  fragment_program   The fragment program.
@param      samplers           The samplers.
@param[in]  samplers_count     The samplers count.
@param[in]  data               The data (copied).
@param[in]  data_bytes_length  The data bytes length.

@return     The color source.
*/
func (context Context) ColorSourceCreateFragmentProgramNew(fragment_program FragmentProgram, samplers *Texture, samplers_count uint64, data string) ColorSource {
	var result ColorSource
	c_data := cString(data)
	data_bytes_length := len(data)
	_, err := ffi.CallFunction(
		cifImpellerColorSourceCreateFragmentProgramNew,
		funcImpellerColorSourceCreateFragmentProgramNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&context),
			unsafe.Pointer(&fragment_program),
			unsafe.Pointer(&samplers),
			unsafe.Pointer(&samplers_count),
			unsafe.Pointer(&c_data),
			unsafe.Pointer(&data_bytes_length),
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

@param[in]  color_filter  The color filter.
*/
func (color_filter ColorFilter) Retain() {
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
func (color_filter ColorFilter) Release() {
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
func ColorFilterCreateBlendNew(color *Color, blend_mode BlendMode) ColorFilter {
	var result ColorFilter
	_, err := ffi.CallFunction(
		cifImpellerColorFilterCreateBlendNew,
		funcImpellerColorFilterCreateBlendNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&color),
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
func ColorFilterCreateColorMatrixNew(color_matrix *ColorMatrix) ColorFilter {
	var result ColorFilter
	_, err := ffi.CallFunction(
		cifImpellerColorFilterCreateColorMatrixNew,
		funcImpellerColorFilterCreateColorMatrixNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&color_matrix),
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
func (mask_filter MaskFilter) Retain() {
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
func (mask_filter MaskFilter) Release() {
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
func (image_filter ImageFilter) Retain() {
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
func (image_filter ImageFilter) Release() {
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
func ImageFilterCreateMatrixNew(matrix *Matrix, sampling TextureSampling) ImageFilter {
	var result ImageFilter
	_, err := ffi.CallFunction(
		cifImpellerImageFilterCreateMatrixNew,
		funcImpellerImageFilterCreateMatrixNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&matrix),
			unsafe.Pointer(&sampling),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
Create an image filter where each pixel is shaded by a fragment
program.

@see        https://docs.flutter.dev/ui/design/graphics/fragment-shaders

@param[in]  context            The context.
@param[in]  fragment_program   The fragment program.
@param      samplers           The samplers.
@param[in]  samplers_count     The samplers count.
@param[in]  data               The data (copied).
@param[in]  data_bytes_length  The data bytes length.

@return     The image filter.
*/
func (context Context) ImageFilterCreateFragmentProgramNew(fragment_program FragmentProgram, samplers *Texture, samplers_count uint64, data string) ImageFilter {
	var result ImageFilter
	c_data := cString(data)
	data_bytes_length := len(data)
	_, err := ffi.CallFunction(
		cifImpellerImageFilterCreateFragmentProgramNew,
		funcImpellerImageFilterCreateFragmentProgramNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&context),
			unsafe.Pointer(&fragment_program),
			unsafe.Pointer(&samplers),
			unsafe.Pointer(&samplers_count),
			unsafe.Pointer(&c_data),
			unsafe.Pointer(&data_bytes_length),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
Creates a composed filter that when applied is identical to
subsequently applying the inner and then the outer filters.

	destination = outer_filter(inner_filter(source))

@param[in]  outer  The outer image filter.
@param[in]  inner  The inner image filter.

@return     The combined image filter.
*/
func (outer ImageFilter) CreateComposeNew(inner ImageFilter) ImageFilter {
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
func (display_list DisplayList) Retain() {
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
func (display_list DisplayList) Release() {
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
func DisplayListBuilderNew(cull_rect *Rect) DisplayListBuilder {
	var result DisplayListBuilder
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderNew,
		funcImpellerDisplayListBuilderNew,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&cull_rect),
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
func (builder DisplayListBuilder) Retain() {
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
func (builder DisplayListBuilder) Release() {
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
func (builder DisplayListBuilder) CreateDisplayListNew() DisplayList {
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
func (builder DisplayListBuilder) Save() {
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
func (builder DisplayListBuilder) SaveLayer(bounds *Rect, paint Paint, backdrop ImageFilter) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderSaveLayer,
		funcImpellerDisplayListBuilderSaveLayer,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&bounds),
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
func (builder DisplayListBuilder) Restore() {
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
func (builder DisplayListBuilder) Scale(x_scale float32, y_scale float32) {
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
func (builder DisplayListBuilder) Rotate(angle_degrees float32) {
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
func (builder DisplayListBuilder) Translate(x_translation float32, y_translation float32) {
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
func (builder DisplayListBuilder) Transform(transform *Matrix) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderTransform,
		funcImpellerDisplayListBuilderTransform,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&transform),
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
func (builder DisplayListBuilder) SetTransform(transform *Matrix) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderSetTransform,
		funcImpellerDisplayListBuilderSetTransform,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&transform),
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
func (builder DisplayListBuilder) GetTransform(out_transform *Matrix) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderGetTransform,
		funcImpellerDisplayListBuilderGetTransform,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&out_transform),
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
func (builder DisplayListBuilder) ResetTransform() {
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
func (builder DisplayListBuilder) GetSaveCount() uint32 {
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
func (builder DisplayListBuilder) RestoreToCount(count uint32) {
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
func (builder DisplayListBuilder) ClipRect(rect *Rect, op ClipOperation) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderClipRect,
		funcImpellerDisplayListBuilderClipRect,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&rect),
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
func (builder DisplayListBuilder) ClipOval(oval_bounds *Rect, op ClipOperation) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderClipOval,
		funcImpellerDisplayListBuilderClipOval,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&oval_bounds),
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
func (builder DisplayListBuilder) ClipRoundedRect(rect *Rect, radii *RoundingRadii, op ClipOperation) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderClipRoundedRect,
		funcImpellerDisplayListBuilderClipRoundedRect,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&rect),
			unsafe.Pointer(&radii),
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
func (builder DisplayListBuilder) ClipPath(path Path, op ClipOperation) {
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
func (builder DisplayListBuilder) DrawPaint(paint Paint) {
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
func (builder DisplayListBuilder) DrawLine(from *Point, to *Point, paint Paint) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderDrawLine,
		funcImpellerDisplayListBuilderDrawLine,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&from),
			unsafe.Pointer(&to),
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
func (builder DisplayListBuilder) DrawDashedLine(from *Point, to *Point, on_length float32, off_length float32, paint Paint) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderDrawDashedLine,
		funcImpellerDisplayListBuilderDrawDashedLine,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&from),
			unsafe.Pointer(&to),
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
func (builder DisplayListBuilder) DrawRect(rect *Rect, paint Paint) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderDrawRect,
		funcImpellerDisplayListBuilderDrawRect,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&rect),
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
func (builder DisplayListBuilder) DrawOval(oval_bounds *Rect, paint Paint) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderDrawOval,
		funcImpellerDisplayListBuilderDrawOval,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&oval_bounds),
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
func (builder DisplayListBuilder) DrawRoundedRect(rect *Rect, radii *RoundingRadii, paint Paint) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderDrawRoundedRect,
		funcImpellerDisplayListBuilderDrawRoundedRect,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&rect),
			unsafe.Pointer(&radii),
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
func (builder DisplayListBuilder) DrawRoundedRectDifference(outer_rect *Rect, outer_radii *RoundingRadii, inner_rect *Rect, inner_radii *RoundingRadii, paint Paint) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderDrawRoundedRectDifference,
		funcImpellerDisplayListBuilderDrawRoundedRectDifference,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&outer_rect),
			unsafe.Pointer(&outer_radii),
			unsafe.Pointer(&inner_rect),
			unsafe.Pointer(&inner_radii),
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
func (builder DisplayListBuilder) DrawPath(path Path, paint Paint) {
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
func (builder DisplayListBuilder) DrawDisplayList(display_list DisplayList, opacity float32) {
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
func (builder DisplayListBuilder) DrawParagraph(paragraph Paragraph, point *Point) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderDrawParagraph,
		funcImpellerDisplayListBuilderDrawParagraph,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&paragraph),
			unsafe.Pointer(&point),
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
func (builder DisplayListBuilder) DrawShadow(path Path, color *Color, elevation float32, occluder_is_transparent Bool, device_pixel_ratio float32) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderDrawShadow,
		funcImpellerDisplayListBuilderDrawShadow,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&path),
			unsafe.Pointer(&color),
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
func (builder DisplayListBuilder) DrawTexture(texture Texture, point *Point, sampling TextureSampling, paint Paint) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderDrawTexture,
		funcImpellerDisplayListBuilderDrawTexture,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&texture),
			unsafe.Pointer(&point),
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
func (builder DisplayListBuilder) DrawTextureRect(texture Texture, src_rect *Rect, dst_rect *Rect, sampling TextureSampling, paint Paint) {
	_, err := ffi.CallFunction(
		cifImpellerDisplayListBuilderDrawTextureRect,
		funcImpellerDisplayListBuilderDrawTextureRect,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&builder),
			unsafe.Pointer(&texture),
			unsafe.Pointer(&src_rect),
			unsafe.Pointer(&dst_rect),
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
func (context TypographyContext) Retain() {
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
func (context TypographyContext) Release() {
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

/*
Register a custom font.

The following font formats are supported:
* OpenType font collections (.ttc extension)
* TrueType fonts: (.ttf extension)
* OpenType fonts: (.otf extension)

@warning    Web Open Font Formats (.woff and .woff2 extensions) are **not**
supported.

The font data is specified as a mapping. It is possible for the
release callback of the mapping to not be called even past the
destruction of the typography context. Care must be taken to not
collect the mapping till the release callback is invoked by
Impeller.

The family alias name can be NULL. In such cases, the font
family specified in paragraph styles must match the family that
is specified in the font data.

If the family name alias is not NULL, that family name must be
used in the paragraph style to reference glyphs from this font
instead of the one encoded in the font itself.

Multiple fonts (with glyphs for different styles) can be
specified with the same family.

@see        `ImpellerParagraphStyleSetFontFamily`

@param[in]  context                        The context.
@param[in]  contents                       The contents.
@param[in]  contents_on_release_user_data  The user data baton to be passed
to the contents release callback.
@param[in]  family_name_alias              The family name alias or NULL if
the one specified in the font
data is to be used.

@return     If the font could be successfully registered.
*/
func (context TypographyContext) RegisterFont(contents *Mapping, contents_on_release_user_data unsafe.Pointer, family_name_alias string) Bool {
	var result Bool
	c_family_name_alias := cString(family_name_alias)
	_, err := ffi.CallFunction(
		cifImpellerTypographyContextRegisterFont,
		funcImpellerTypographyContextRegisterFont,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&context),
			unsafe.Pointer(&contents),
			unsafe.Pointer(&contents_on_release_user_data),
			unsafe.Pointer(&c_family_name_alias),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

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
func (paragraph_style ParagraphStyle) Retain() {
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
func (paragraph_style ParagraphStyle) Release() {
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
func (paragraph_style ParagraphStyle) SetForeground(paint Paint) {
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
func (paragraph_style ParagraphStyle) SetBackground(paint Paint) {
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
func (paragraph_style ParagraphStyle) SetFontWeight(weight FontWeight) {
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
func (paragraph_style ParagraphStyle) SetFontStyle(style FontStyle) {
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

/*
Set the font family.

@param[in]  paragraph_style  The paragraph style.
@param[in]  family_name      The family name.
*/
func (paragraph_style ParagraphStyle) SetFontFamily(family_name string) {
	c_family_name := cString(family_name)
	_, err := ffi.CallFunction(
		cifImpellerParagraphStyleSetFontFamily,
		funcImpellerParagraphStyleSetFontFamily,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph_style),
			unsafe.Pointer(&c_family_name),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Set the font size.

@param[in]  paragraph_style  The paragraph style.
@param[in]  size             The size.
*/
func (paragraph_style ParagraphStyle) SetFontSize(size float32) {
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
func (paragraph_style ParagraphStyle) SetHeight(height float32) {
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
func (paragraph_style ParagraphStyle) SetTextAlignment(align TextAlignment) {
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
func (paragraph_style ParagraphStyle) SetTextDirection(direction TextDirection) {
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
func (paragraph_style ParagraphStyle) SetTextDecoration(decoration *TextDecoration) {
	_, err := ffi.CallFunction(
		cifImpellerParagraphStyleSetTextDecoration,
		funcImpellerParagraphStyleSetTextDecoration,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph_style),
			unsafe.Pointer(&decoration),
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
func (paragraph_style ParagraphStyle) SetMaxLines(max_lines uint32) {
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

/*
Set the paragraph locale.

@param[in]  paragraph_style  The paragraph style.
@param[in]  locale           The locale.
*/
func (paragraph_style ParagraphStyle) SetLocale(locale string) {
	c_locale := cString(locale)
	_, err := ffi.CallFunction(
		cifImpellerParagraphStyleSetLocale,
		funcImpellerParagraphStyleSetLocale,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph_style),
			unsafe.Pointer(&c_locale),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Set the UTF-8 string to use as the ellipsis. Pass `nullptr` to
clear the setting to default.

@param[in]  paragraph_style  The paragraph style.
@param[in]  data             The ellipsis string UTF-8 data, or null.
*/
func (paragraph_style ParagraphStyle) SetEllipsis(ellipsis string) {
	c_ellipsis := cString(ellipsis)
	_, err := ffi.CallFunction(
		cifImpellerParagraphStyleSetEllipsis,
		funcImpellerParagraphStyleSetEllipsis,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph_style),
			unsafe.Pointer(&c_ellipsis),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Create a new paragraph builder.

@param[in]  context  The context.

@return     The paragraph builder.
*/
func (context TypographyContext) ParagraphBuilderNew() ParagraphBuilder {
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
func (paragraph_builder ParagraphBuilder) Retain() {
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
func (paragraph_builder ParagraphBuilder) Release() {
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
func (paragraph_builder ParagraphBuilder) PushStyle(style ParagraphStyle) {
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
func (paragraph_builder ParagraphBuilder) PopStyle() {
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

/*
Add UTF-8 encoded text to the paragraph. The text will be styled
according to the paragraph style already on top of the paragraph
style stack.

@param[in]  paragraph_builder  The paragraph builder.
@param[in]  data               The data.
@param[in]  length             The length.
*/
func (paragraph_builder ParagraphBuilder) AddText(data string) {
	c_data := cString(data)
	length := len(data)
	_, err := ffi.CallFunction(
		cifImpellerParagraphBuilderAddText,
		funcImpellerParagraphBuilderAddText,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph_builder),
			unsafe.Pointer(&c_data),
			unsafe.Pointer(&length),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Layout and build a new paragraph using the specified width. The
resulting paragraph is immutable. The paragraph builder must be
discarded and a new one created to build more paragraphs.

@param[in]  paragraph_builder  The paragraph builder.
@param[in]  width              The paragraph width.

@return     The paragraph if one can be created, NULL otherwise.
*/
func (paragraph_builder ParagraphBuilder) BuildParagraphNew(width float32) Paragraph {
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
func (paragraph Paragraph) Retain() {
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
func (paragraph Paragraph) Release() {
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
func (paragraph Paragraph) GetMaxWidth() float32 {
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
func (paragraph Paragraph) GetHeight() float32 {
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
func (paragraph Paragraph) GetLongestLineWidth() float32 {
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
func (paragraph Paragraph) GetMinIntrinsicWidth() float32 {
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
func (paragraph Paragraph) GetMaxIntrinsicWidth() float32 {
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
func (paragraph Paragraph) GetIdeographicBaseline() float32 {
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
func (paragraph Paragraph) GetAlphabeticBaseline() float32 {
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
func (paragraph Paragraph) GetLineCount() uint32 {
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
func (paragraph Paragraph) GetWordBoundary(code_unit_index uint64, out_range *Range) {
	_, err := ffi.CallFunction(
		cifImpellerParagraphGetWordBoundary,
		funcImpellerParagraphGetWordBoundary,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&paragraph),
			unsafe.Pointer(&code_unit_index),
			unsafe.Pointer(&out_range),
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
func (paragraph Paragraph) GetLineMetrics() LineMetrics {
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
func (paragraph Paragraph) CreateGlyphInfoAtCodeUnitIndexNew(code_unit_index uint64) GlyphInfo {
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
func (paragraph Paragraph) CreateGlyphInfoAtParagraphCoordinatesNew(x float64, y float64) GlyphInfo {
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
func (line_metrics LineMetrics) Retain() {
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
func (line_metrics LineMetrics) Release() {
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
func (metrics LineMetrics) GetUnscaledAscent(line uint64) float64 {
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
func (metrics LineMetrics) GetAscent(line uint64) float64 {
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
func (metrics LineMetrics) GetDescent(line uint64) float64 {
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
func (metrics LineMetrics) GetBaseline(line uint64) float64 {
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
func (metrics LineMetrics) IsHardbreak(line uint64) Bool {
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
func (metrics LineMetrics) GetWidth(line uint64) float64 {
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
func (metrics LineMetrics) GetHeight(line uint64) float64 {
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
func (metrics LineMetrics) GetLeft(line uint64) float64 {
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
func (metrics LineMetrics) GetCodeUnitStartIndex(line uint64) uint64 {
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
func (metrics LineMetrics) GetCodeUnitEndIndex(line uint64) uint64 {
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
func (metrics LineMetrics) GetCodeUnitEndIndexExcludingWhitespace(line uint64) uint64 {
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
func (metrics LineMetrics) GetCodeUnitEndIndexIncludingNewline(line uint64) uint64 {
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
func (glyph_info GlyphInfo) Retain() {
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
func (glyph_info GlyphInfo) Release() {
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
func (glyph_info GlyphInfo) GetGraphemeClusterCodeUnitRangeBegin() uint64 {
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
func (glyph_info GlyphInfo) GetGraphemeClusterCodeUnitRangeEnd() uint64 {
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
func (glyph_info GlyphInfo) GetGraphemeClusterBounds(out_bounds *Rect) {
	_, err := ffi.CallFunction(
		cifImpellerGlyphInfoGetGraphemeClusterBounds,
		funcImpellerGlyphInfoGetGraphemeClusterBounds,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&glyph_info),
			unsafe.Pointer(&out_bounds),
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
func (glyph_info GlyphInfo) IsEllipsis() Bool {
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
func (glyph_info GlyphInfo) GetTextDirection() TextDirection {
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
