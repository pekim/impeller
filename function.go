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
// UNSUPPORTED :: ContextCreateOpenGLESNew  param count = 3

/*
Create a Metal context using the system default Metal device.

@param[in]  version  The version specified in the IMPELLER_VERSION macro.

@return     The Metal context or NULL if one cannot be created.
*/
// UNSUPPORTED :: ContextCreateMetalNew  param count = 1

/*
Create a Vulkan context using the provided Vulkan Settings.

@param[in]  version   The version specified in the IMPELLER_VERSION macro.
@param[in]  settings  The Vulkan settings.

@return     The Vulkan context or NULL if one cannot be created.
*/
// UNSUPPORTED :: ContextCreateVulkanNew  param count = 2

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  context  The context.
*/
// UNSUPPORTED :: ContextRetain  param count = 1

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  context  The context.
*/
// UNSUPPORTED :: ContextRelease  param count = 1

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
// UNSUPPORTED :: ContextGetVulkanInfo  param count = 2

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
// UNSUPPORTED :: VulkanSwapchainCreateNew  param count = 2

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  swapchain  The swapchain.
*/
// UNSUPPORTED :: VulkanSwapchainRetain  param count = 1

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  swapchain  The swapchain.
*/
// UNSUPPORTED :: VulkanSwapchainRelease  param count = 1

/*
A potentially blocking operation, acquires the next surface to
render to. Since this may block, surface acquisition must be
delayed for as long as possible to avoid an idle wait on the
CPU.

@param[in]  swapchain  The swapchain.

@return     The surface if one could be obtained, NULL otherwise.
*/
// UNSUPPORTED :: VulkanSwapchainAcquireNextSurfaceNew  param count = 1

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
// UNSUPPORTED :: SurfaceCreateWrappedFBONew  param count = 4

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
// UNSUPPORTED :: SurfaceCreateWrappedMetalDrawableNew  param count = 2

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  surface  The surface.
*/
// UNSUPPORTED :: SurfaceRetain  param count = 1

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  surface  The surface.
*/
// UNSUPPORTED :: SurfaceRelease  param count = 1

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
// UNSUPPORTED :: SurfaceDrawDisplayList  param count = 2

/*
Present the surface to the underlying window system.

@param[in]  surface  The surface to present.

@return     True if the surface could be presented.
*/
// UNSUPPORTED :: SurfacePresent  param count = 1

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  path  The path.
*/
// UNSUPPORTED :: PathRetain  param count = 1

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  path  The path.
*/
// UNSUPPORTED :: PathRelease  param count = 1

/*
Get the bounds of the path.

The bounds are conservative. That is, they may be larger than
the actual shape of the path and could include the control
points and isolated calls to move the cursor.

@param[in]  path        The path
@param[out] out_bounds  The conservative bounds of the path.
*/
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

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  builder  The builder.
*/
// UNSUPPORTED :: PathBuilderRetain  param count = 1

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  builder  The builder.
*/
// UNSUPPORTED :: PathBuilderRelease  param count = 1

/*
Move the cursor to the specified location.

@param[in]  builder   The builder.
@param[in]  location  The location.
*/
// UNSUPPORTED :: PathBuilderMoveTo  param count = 2

/*
Add a line segment from the current cursor location to the given
location. The cursor location is updated to be at the endpoint.

@param[in]  builder   The builder.
@param[in]  location  The location.
*/
// UNSUPPORTED :: PathBuilderLineTo  param count = 2

/*
Add a quadratic curve from whose start point is the cursor to
the specified end point using the a single control point.

The new location of the cursor after this call is the end point.

@param[in]  builder        The builder.
@param[in]  control_point  The control point.
@param[in]  end_point      The end point.
*/
// UNSUPPORTED :: PathBuilderQuadraticCurveTo  param count = 3

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
// UNSUPPORTED :: PathBuilderCubicCurveTo  param count = 4

/*
Adds a rectangle to the path.

@param[in]  builder  The builder.
@param[in]  rect     The rectangle.
*/
// UNSUPPORTED :: PathBuilderAddRect  param count = 2

/*
Add an arc to the path.

@param[in]  builder              The builder.
@param[in]  oval_bounds          The oval bounds.
@param[in]  start_angle_degrees  The start angle in degrees.
@param[in]  end_angle_degrees    The end angle in degrees.
*/
// UNSUPPORTED :: PathBuilderAddArc  param count = 4

/*
Add an oval to the path.

@param[in]  builder      The builder.
@param[in]  oval_bounds  The oval bounds.
*/
// UNSUPPORTED :: PathBuilderAddOval  param count = 2

/*
Add a rounded rect with potentially non-uniform radii to the
path.

@param[in]  builder         The builder.
@param[in]  rect            The rectangle.
@param[in]  rounding_radii  The rounding radii.
*/
// UNSUPPORTED :: PathBuilderAddRoundedRect  param count = 3

/*
Close the path.

@param[in]  builder  The builder.
*/
// UNSUPPORTED :: PathBuilderClose  param count = 1

/*
Create a new path by copying the existing built-up path. The
existing path can continue being added to.

@param[in]  builder  The builder.
@param[in]  fill     The fill.

@return     The impeller path.
*/
// UNSUPPORTED :: PathBuilderCopyPathNew  param count = 2

/*
Create a new path using the existing built-up path. The existing
path builder now contains an empty path.

@param[in]  builder  The builder.
@param[in]  fill     The fill.

@return     The impeller path.
*/
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

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  paint  The paint.
*/
// UNSUPPORTED :: PaintRetain  param count = 1

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  paint  The paint.
*/
// UNSUPPORTED :: PaintRelease  param count = 1

/*
Set the paint color.

@param[in]  paint  The paint.
@param[in]  color  The color.
*/
// UNSUPPORTED :: PaintSetColor  param count = 2

/*
Set the paint blend mode. The blend mode controls how the new
paints contents are mixed with the values already drawn using
previous draw calls.

@param[in]  paint  The paint.
@param[in]  mode   The mode.
*/
// UNSUPPORTED :: PaintSetBlendMode  param count = 2

/*
Set the paint draw style. The style controls if the closed
shapes are filled and/or stroked.

@param[in]  paint  The paint.
@param[in]  style  The style.
*/
// UNSUPPORTED :: PaintSetDrawStyle  param count = 2

/*
Sets how strokes rendered using this paint are capped.

@param[in]  paint  The paint.
@param[in]  cap    The stroke cap style.
*/
// UNSUPPORTED :: PaintSetStrokeCap  param count = 2

/*
Sets how strokes rendered using this paint are joined.

@param[in]  paint  The paint.
@param[in]  join   The join.
*/
// UNSUPPORTED :: PaintSetStrokeJoin  param count = 2

/*
Set the width of the strokes rendered using this paint.

@param[in]  paint  The paint.
@param[in]  width  The width.
*/
// UNSUPPORTED :: PaintSetStrokeWidth  param count = 2

/*
Set the miter limit of the strokes rendered using this paint.

@param[in]  paint  The paint.
@param[in]  miter  The miter limit.
*/
// UNSUPPORTED :: PaintSetStrokeMiter  param count = 2

/*
Set the color filter of the paint.

Color filters are functions that take two colors and mix them to
produce a single color. This color is then usually merged with
the destination during blending.

@param[in]  paint         The paint.
@param[in]  color_filter  The color filter.
*/
// UNSUPPORTED :: PaintSetColorFilter  param count = 2

/*
Set the color source of the paint.

Color sources are functions that generate colors for each
texture element covered by a draw call.

@param[in]  paint         The paint.
@param[in]  color_source  The color source.
*/
// UNSUPPORTED :: PaintSetColorSource  param count = 2

/*
Set the image filter of a paint.

Image filters are functions that are applied to regions of a
texture to produce a single color.

@param[in]  paint         The paint.
@param[in]  image_filter  The image filter.
*/
// UNSUPPORTED :: PaintSetImageFilter  param count = 2

/*
Set the mask filter of a paint.

@param[in]  paint        The paint.
@param[in]  mask_filter  The mask filter.
*/
// UNSUPPORTED :: PaintSetMaskFilter  param count = 2

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
// UNSUPPORTED :: TextureCreateWithContentsNew  param count = 4

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
// UNSUPPORTED :: TextureCreateWithOpenGLTextureHandleNew  param count = 3

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  texture  The texture.
*/
// UNSUPPORTED :: TextureRetain  param count = 1

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  texture  The texture.
*/
// UNSUPPORTED :: TextureRelease  param count = 1

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
// UNSUPPORTED :: TextureGetOpenGLHandle  param count = 1

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
// UNSUPPORTED :: FragmentProgramNew  param count = 2

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  fragment_program  The fragment program.
*/
// UNSUPPORTED :: FragmentProgramRetain  param count = 1

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  fragment_program  The fragment program.
*/
// UNSUPPORTED :: FragmentProgramRelease  param count = 1

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  color_source  The color source.
*/
// UNSUPPORTED :: ColorSourceRetain  param count = 1

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  color_source  The color source.
*/
// UNSUPPORTED :: ColorSourceRelease  param count = 1

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
// UNSUPPORTED :: ColorSourceCreateLinearGradientNew  param count = 7

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
// UNSUPPORTED :: ColorSourceCreateRadialGradientNew  param count = 7

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
// UNSUPPORTED :: ColorSourceCreateConicalGradientNew  param count = 9

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
// UNSUPPORTED :: ColorSourceCreateSweepGradientNew  param count = 8

/*
Create a color source that samples from an image.

@param[in]  image                 The image.
@param[in]  horizontal_tile_mode  The horizontal tile mode.
@param[in]  vertical_tile_mode    The vertical tile mode.
@param[in]  sampling              The sampling.
@param[in]  transformation        The transformation.

@return     The color source.
*/
// UNSUPPORTED :: ColorSourceCreateImageNew  param count = 5

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
// UNSUPPORTED :: ColorSourceCreateFragmentProgramNew  param count = 6

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  color_filter  The color filter.
*/
// UNSUPPORTED :: ColorFilterRetain  param count = 1

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  color_filter  The color filter.
*/
// UNSUPPORTED :: ColorFilterRelease  param count = 1

/*
Create a color filter that performs blending of pixel values
independently.

@param[in]  color       The color.
@param[in]  blend_mode  The blend mode.

@return     The color filter.
*/
// UNSUPPORTED :: ColorFilterCreateBlendNew  param count = 2

/*
Create a color filter that transforms pixel color values
independently.

@param[in]  color_matrix  The color matrix.

@return     The color filter.
*/
// UNSUPPORTED :: ColorFilterCreateColorMatrixNew  param count = 1

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  mask_filter  The mask filter.
*/
// UNSUPPORTED :: MaskFilterRetain  param count = 1

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  mask_filter  The mask filter.
*/
// UNSUPPORTED :: MaskFilterRelease  param count = 1

/*
Create a mask filter that blurs contents in the masked shape.

@param[in]  style  The style.
@param[in]  sigma  The sigma.

@return     The mask filter.
*/
// UNSUPPORTED :: MaskFilterCreateBlurNew  param count = 2

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  image_filter  The image filter.
*/
// UNSUPPORTED :: ImageFilterRetain  param count = 1

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  image_filter  The image filter.
*/
// UNSUPPORTED :: ImageFilterRelease  param count = 1

/*
Creates an image filter that applies a Gaussian blur.

The Gaussian blur applied may be an approximation for
performance.


@param[in]  x_sigma    The x sigma.
@param[in]  y_sigma    The y sigma.
@param[in]  tile_mode  The tile mode.

@return     The image filter.
*/
// UNSUPPORTED :: ImageFilterCreateBlurNew  param count = 3

/*
Creates an image filter that enhances the per-channel pixel
values to the maximum value in a circle around the pixel.

@param[in]  x_radius  The x radius.
@param[in]  y_radius  The y radius.

@return     The image filter.
*/
// UNSUPPORTED :: ImageFilterCreateDilateNew  param count = 2

/*
Creates an image filter that dampens the per-channel pixel
values to the minimum value in a circle around the pixel.

@param[in]  x_radius  The x radius.
@param[in]  y_radius  The y radius.

@return     The image filter.
*/
// UNSUPPORTED :: ImageFilterCreateErodeNew  param count = 2

/*
Creates an image filter that applies a transformation matrix to
the underlying image.

@param[in]  matrix    The transformation matrix.
@param[in]  sampling  The image sampling mode.

@return     The image filter.
*/
// UNSUPPORTED :: ImageFilterCreateMatrixNew  param count = 2

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
// UNSUPPORTED :: ImageFilterCreateFragmentProgramNew  param count = 6

/*
Creates a composed filter that when applied is identical to
subsequently applying the inner and then the outer filters.

  destination = outer_filter(inner_filter(source))

@param[in]  outer  The outer image filter.
@param[in]  inner  The inner image filter.

@return     The combined image filter.
*/
// UNSUPPORTED :: ImageFilterCreateComposeNew  param count = 2

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  display_list  The display list.
*/
// UNSUPPORTED :: DisplayListRetain  param count = 1

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  display_list  The display list.
*/
// UNSUPPORTED :: DisplayListRelease  param count = 1

/*
Create a new display list builder.

An optional cull rectangle may be specified. Impeller is allowed
to treat the contents outside this rectangle as being undefined.
This may aid performance optimizations.

@param[in]  cull_rect  The cull rectangle or NULL.

@return     The display list builder.
*/
// UNSUPPORTED :: DisplayListBuilderNew  param count = 1

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  builder  The display list builder.
*/
// UNSUPPORTED :: DisplayListBuilderRetain  param count = 1

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  builder  The display list builder.
*/
// UNSUPPORTED :: DisplayListBuilderRelease  param count = 1

/*
Create a new display list using the rendering intent already
encoded in the builder. The builder is reset after this call.

@param[in]  builder  The builder.

@return     The display list.
*/
// UNSUPPORTED :: DisplayListBuilderCreateDisplayListNew  param count = 1

/*
Stashes the current transformation and clip state onto a save
stack.

@param[in]  builder  The builder.
*/
// UNSUPPORTED :: DisplayListBuilderSave  param count = 1

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
// UNSUPPORTED :: DisplayListBuilderSaveLayer  param count = 4

/*
Pops the last entry pushed onto the save stack using a call to
`ImpellerDisplayListBuilderSave` or
`ImpellerDisplayListBuilderSaveLayer`.

@param[in]  builder  The builder.
*/
// UNSUPPORTED :: DisplayListBuilderRestore  param count = 1

/*
Apply a scale to the transformation matrix currently on top of
the save stack.

@param[in]  builder  The builder.
@param[in]  x_scale  The x scale.
@param[in]  y_scale  The y scale.
*/
// UNSUPPORTED :: DisplayListBuilderScale  param count = 3

/*
Apply a clockwise rotation to the transformation matrix
currently on top of the save stack.

@param[in]  builder        The builder.
@param[in]  angle_degrees  The angle in degrees.
*/
// UNSUPPORTED :: DisplayListBuilderRotate  param count = 2

/*
Apply a translation to the transformation matrix currently on
top of the save stack.

@param[in]  builder        The builder.
@param[in]  x_translation  The x translation.
@param[in]  y_translation  The y translation.
*/
// UNSUPPORTED :: DisplayListBuilderTranslate  param count = 3

/*
Appends the the provided transformation to the transformation
already on the save stack.

@param[in]  builder    The builder.
@param[in]  transform  The transform to append.
*/
// UNSUPPORTED :: DisplayListBuilderTransform  param count = 2

/*
Clear the transformation on top of the save stack and replace it
with a new value.

@param[in]  builder    The builder.
@param[in]  transform  The new transform.
*/
// UNSUPPORTED :: DisplayListBuilderSetTransform  param count = 2

/*
Get the transformation currently built up on the top of the
transformation stack.

@param[in]  builder        The builder.
@param[out] out_transform  The transform.
*/
// UNSUPPORTED :: DisplayListBuilderGetTransform  param count = 2

/*
Reset the transformation on top of the transformation stack to
identity.

@param[in]  builder  The builder.
*/
// UNSUPPORTED :: DisplayListBuilderResetTransform  param count = 1

/*
Get the current size of the save stack.

@param[in]  builder  The builder.

@return     The save stack size.
*/
// UNSUPPORTED :: DisplayListBuilderGetSaveCount  param count = 1

/*
Effectively calls ImpellerDisplayListBuilderRestore till the
size of the save stack becomes a specified count.

@param[in]  builder  The builder.
@param[in]  count    The count.
*/
// UNSUPPORTED :: DisplayListBuilderRestoreToCount  param count = 2

/*
Reduces the clip region to the intersection of the current clip
and the given rectangle taking into account the clip operation.

@param[in]  builder  The builder.
@param[in]  rect     The rectangle.
@param[in]  op       The operation.
*/
// UNSUPPORTED :: DisplayListBuilderClipRect  param count = 3

/*
Reduces the clip region to the intersection of the current clip
and the given oval taking into account the clip operation.

@param[in]  builder      The builder.
@param[in]  oval_bounds  The oval bounds.
@param[in]  op           The operation.
*/
// UNSUPPORTED :: DisplayListBuilderClipOval  param count = 3

/*
Reduces the clip region to the intersection of the current clip
and the given rounded rectangle taking into account the clip
operation.

@param[in]  builder  The builder.
@param[in]  rect     The rectangle.
@param[in]  radii    The radii.
@param[in]  op       The operation.
*/
// UNSUPPORTED :: DisplayListBuilderClipRoundedRect  param count = 4

/*
Reduces the clip region to the intersection of the current clip
and the given path taking into account the clip operation.

@param[in]  builder  The builder.
@param[in]  path     The path.
@param[in]  op       The operation.
*/
// UNSUPPORTED :: DisplayListBuilderClipPath  param count = 3

/*
Fills the current clip with the specified paint.

@param[in]  builder  The builder.
@param[in]  paint    The paint.
*/
// UNSUPPORTED :: DisplayListBuilderDrawPaint  param count = 2

/*
Draws a line segment.

@param[in]  builder  The builder.
@param[in]  from     The starting point of the line.
@param[in]  to       The end point of the line.
@param[in]  paint    The paint.
*/
// UNSUPPORTED :: DisplayListBuilderDrawLine  param count = 4

/*
Draws a dash line segment.

@param[in]  builder     The builder.
@param[in]  from        The starting point of the line.
@param[in]  to          The end point of the line.
@param[in]  on_length   On length.
@param[in]  off_length  Off length.
@param[in]  paint       The paint.
*/
// UNSUPPORTED :: DisplayListBuilderDrawDashedLine  param count = 6

/*
Draws a rectangle.

@param[in]  builder  The builder.
@param[in]  rect     The rectangle.
@param[in]  paint    The paint.
*/
// UNSUPPORTED :: DisplayListBuilderDrawRect  param count = 3

/*
Draws an oval.

@param[in]  builder      The builder.
@param[in]  oval_bounds  The oval bounds.
@param[in]  paint        The paint.
*/
// UNSUPPORTED :: DisplayListBuilderDrawOval  param count = 3

/*
Draws a rounded rect.

@param[in]  builder  The builder.
@param[in]  rect     The rectangle.
@param[in]  radii    The radii.
@param[in]  paint    The paint.
*/
// UNSUPPORTED :: DisplayListBuilderDrawRoundedRect  param count = 4

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
// UNSUPPORTED :: DisplayListBuilderDrawRoundedRectDifference  param count = 6

/*
Draws the specified shape.

@param[in]  builder  The builder.
@param[in]  path     The path.
@param[in]  paint    The paint.
*/
// UNSUPPORTED :: DisplayListBuilderDrawPath  param count = 3

/*
Flattens the contents of another display list into the one
currently being built.

@param[in]  builder       The builder.
@param[in]  display_list  The display list.
@param[in]  opacity       The opacity.
*/
// UNSUPPORTED :: DisplayListBuilderDrawDisplayList  param count = 3

/*
Draw a paragraph at the specified point.

@param[in]  builder    The builder.
@param[in]  paragraph  The paragraph.
@param[in]  point      The point.
*/
// UNSUPPORTED :: DisplayListBuilderDrawParagraph  param count = 3

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
// UNSUPPORTED :: DisplayListBuilderDrawShadow  param count = 6

/*
Draw a texture at the specified point.

@param[in]  builder   The builder.
@param[in]  texture   The texture.
@param[in]  point     The point.
@param[in]  sampling  The sampling.
@param[in]  paint     The paint.
*/
// UNSUPPORTED :: DisplayListBuilderDrawTexture  param count = 5

/*
Draw a portion of texture at the specified location.

@param[in]  builder   The builder.
@param[in]  texture   The texture.
@param[in]  src_rect  The source rectangle.
@param[in]  dst_rect  The destination rectangle.
@param[in]  sampling  The sampling.
@param[in]  paint     The paint.
*/
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

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  context  The typography context.
*/
// UNSUPPORTED :: TypographyContextRetain  param count = 1

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  context  The typography context.
*/
// UNSUPPORTED :: TypographyContextRelease  param count = 1

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

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  paragraph_style  The paragraph style.
*/
// UNSUPPORTED :: ParagraphStyleRetain  param count = 1

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  paragraph_style  The paragraph style.
*/
// UNSUPPORTED :: ParagraphStyleRelease  param count = 1

/*
Set the paint used to render the text glyph contents.

@param[in]  paragraph_style  The paragraph style.
@param[in]  paint            The paint.
*/
// UNSUPPORTED :: ParagraphStyleSetForeground  param count = 2

/*
Set the paint used to render the background of the text glyphs.

@param[in]  paragraph_style  The paragraph style.
@param[in]  paint            The paint.
*/
// UNSUPPORTED :: ParagraphStyleSetBackground  param count = 2

/*
Set the weight of the font to select when rendering glyphs.

@param[in]  paragraph_style  The paragraph style.
@param[in]  weight           The weight.
*/
// UNSUPPORTED :: ParagraphStyleSetFontWeight  param count = 2

/*
Set whether the glyphs should be bolded or italicized.

@param[in]  paragraph_style  The paragraph style.
@param[in]  style            The style.
*/
// UNSUPPORTED :: ParagraphStyleSetFontStyle  param count = 2

/*
Set the font family.

@param[in]  paragraph_style  The paragraph style.
@param[in]  family_name      The family name.
*/
// UNSUPPORTED :: ParagraphStyleSetFontFamily  param count = 2

/*
Set the font size.

@param[in]  paragraph_style  The paragraph style.
@param[in]  size             The size.
*/
// UNSUPPORTED :: ParagraphStyleSetFontSize  param count = 2

/*
The height of the text as a multiple of text size.

When height is 0.0, the line height will be determined by the
font's metrics directly, which may differ from the font size.
Otherwise the line height of the text will be a multiple of font
size, and be exactly fontSize * height logical pixels tall.

@param[in]  paragraph_style  The paragraph style.
@param[in]  height           The height.
*/
// UNSUPPORTED :: ParagraphStyleSetHeight  param count = 2

/*
Set the alignment of text within the paragraph.

@param[in]  paragraph_style  The paragraph style.
@param[in]  align            The align.
*/
// UNSUPPORTED :: ParagraphStyleSetTextAlignment  param count = 2

/*
Set the directionality of the text within the paragraph.

@param[in]  paragraph_style  The paragraph style.
@param[in]  direction        The direction.
*/
// UNSUPPORTED :: ParagraphStyleSetTextDirection  param count = 2

/*
Set one of more text decorations on the paragraph. Decorations
can be underlines, overlines, strikethroughs, etc.. The style of
decorations can be set as well (dashed, dotted, wavy, etc..)

@param[in]  ImpellerParagraphStyle  The paragraph style.
@param[in]  decoration              The text decoration.
*/
// UNSUPPORTED :: ParagraphStyleSetTextDecoration  param count = 2

/*
Set the maximum line count within the paragraph.

@param[in]  paragraph_style  The paragraph style.
@param[in]  max_lines        The maximum lines.
*/
// UNSUPPORTED :: ParagraphStyleSetMaxLines  param count = 2

/*
Set the paragraph locale.

@param[in]  paragraph_style  The paragraph style.
@param[in]  locale           The locale.
*/
// UNSUPPORTED :: ParagraphStyleSetLocale  param count = 2

/*
Set the UTF-8 string to use as the ellipsis. Pass `nullptr` to
clear the setting to default.

@param[in]  paragraph_style  The paragraph style.
@param[in]  data             The ellipsis string UTF-8 data, or null.
*/
// UNSUPPORTED :: ParagraphStyleSetEllipsis  param count = 2

/*
Create a new paragraph builder.

@param[in]  context  The context.

@return     The paragraph builder.
*/
// UNSUPPORTED :: ParagraphBuilderNew  param count = 1

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  paragraph_builder  The paragraph builder.
*/
// UNSUPPORTED :: ParagraphBuilderRetain  param count = 1

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  paragraph_builder  The paragraph_builder.
*/
// UNSUPPORTED :: ParagraphBuilderRelease  param count = 1

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
// UNSUPPORTED :: ParagraphBuilderPushStyle  param count = 2

/*
Pop a previously pushed paragraph style from the paragraph style
stack.

@param[in]  paragraph_builder  The paragraph builder.
*/
// UNSUPPORTED :: ParagraphBuilderPopStyle  param count = 1

/*
Add UTF-8 encoded text to the paragraph. The text will be styled
according to the paragraph style already on top of the paragraph
style stack.

@param[in]  paragraph_builder  The paragraph builder.
@param[in]  data               The data.
@param[in]  length             The length.
*/
// UNSUPPORTED :: ParagraphBuilderAddText  param count = 3

/*
Layout and build a new paragraph using the specified width. The
resulting paragraph is immutable. The paragraph builder must be
discarded and a new one created to build more paragraphs.

@param[in]  paragraph_builder  The paragraph builder.
@param[in]  width              The paragraph width.

@return     The paragraph if one can be created, NULL otherwise.
*/
// UNSUPPORTED :: ParagraphBuilderBuildParagraphNew  param count = 2

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  paragraph  The paragraph.
*/
// UNSUPPORTED :: ParagraphRetain  param count = 1

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  paragraph  The paragraph.
*/
// UNSUPPORTED :: ParagraphRelease  param count = 1

/*
@see        `ImpellerParagraphGetMinIntrinsicWidth`

@param[in]  paragraph  The paragraph.


@return     The width provided to the paragraph builder during the call to
layout. This is the maximum width any line in the laid out
paragraph can occupy. But, it is not necessarily the actual
width of the paragraph after layout.
*/
// UNSUPPORTED :: ParagraphGetMaxWidth  param count = 1

/*
@param[in]  paragraph  The paragraph.

@return     The height of the laid out paragraph. This is **not** a tight
bounding box and some glyphs may not reach the minimum location
they are allowed to reach.
*/
// UNSUPPORTED :: ParagraphGetHeight  param count = 1

/*
@param[in]  paragraph  The paragraph.

@return     The length of the longest line in the paragraph. This is the
horizontal distance between the left edge of the leftmost glyph
and the right edge of the rightmost glyph, in the longest line
in the paragraph.
*/
// UNSUPPORTED :: ParagraphGetLongestLineWidth  param count = 1

/*
@see        `ImpellerParagraphGetMaxWidth`

@param[in]  paragraph  The paragraph.

@return     The actual width of the longest line in the paragraph after
layout. This is expected to be less than or equal to
`ImpellerParagraphGetMaxWidth`.
*/
// UNSUPPORTED :: ParagraphGetMinIntrinsicWidth  param count = 1

/*
@param[in]  paragraph  The paragraph.

@return     The width of the paragraph without line breaking.
*/
// UNSUPPORTED :: ParagraphGetMaxIntrinsicWidth  param count = 1

/*
@param[in]  paragraph  The paragraph.

@return     The distance from the top of the paragraph to the ideographic
baseline of the first line when using ideographic fonts
(Japanese, Korean, etc...).
*/
// UNSUPPORTED :: ParagraphGetIdeographicBaseline  param count = 1

/*
@param[in]  paragraph  The paragraph.

@return     The distance from the top of the paragraph to the alphabetic
baseline of the first line when using alphabetic fonts (A-Z,
a-z, Greek, etc...).
*/
// UNSUPPORTED :: ParagraphGetAlphabeticBaseline  param count = 1

/*
@param[in]  paragraph  The paragraph.

@return     The number of lines visible in the paragraph after line
breaking.
*/
// UNSUPPORTED :: ParagraphGetLineCount  param count = 1

/*
Get the range into the UTF-16 code unit buffer that represents
the word at the specified caret location in the same buffer.

Word boundaries are defined more precisely in [Unicode Standard
Annex #29](http://www.unicode.org/reports/tr29/#Word_Boundaries)

@param[in]  paragraph        The paragraph
@param[in]  code_unit_index  The code unit index
@param[out]  code_unit_index The range.
*/
// UNSUPPORTED :: ParagraphGetWordBoundary  param count = 3

/*
Get the line metrics of this laid out paragraph. Calculating the
line metrics is expensive. The first time line metrics are
requested, they will be cached along with the paragraph (which
is immutable).

@param[in]  paragraph  The paragraph.

@return     The line metrics.
*/
// UNSUPPORTED :: ParagraphGetLineMetrics  param count = 1

/*
Create a new instance of glyph info that can be queried for
information about the glyph at the given UTF-16 code unit index.
The instance must be freed using `ImpellerGlyphInfoRelease`.

@param[in]  paragraph        The paragraph.
@param[in]  code_unit_index  The UTF-16 code unit index.

@return     The glyph information.
*/
// UNSUPPORTED :: ParagraphCreateGlyphInfoAtCodeUnitIndexNew  param count = 2

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
// UNSUPPORTED :: ParagraphCreateGlyphInfoAtParagraphCoordinatesNew  param count = 3

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  line_metrics  The line metrics.
*/
// UNSUPPORTED :: LineMetricsRetain  param count = 1

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  line_metrics  The line metrics.
*/
// UNSUPPORTED :: LineMetricsRelease  param count = 1

/*
The rise from the baseline as calculated from the font and style
for this line ignoring the height from the text style.

@param[in]  metrics  The metrics.
@param[in]  line     The line index (zero based).

@return     The unscaled ascent.
*/
// UNSUPPORTED :: LineMetricsGetUnscaledAscent  param count = 2

/*
The rise from the baseline as calculated from the font and style
for this line.

@param[in]  metrics  The metrics.
@param[in]  line     The line index (zero based).

@return     The ascent.
*/
// UNSUPPORTED :: LineMetricsGetAscent  param count = 2

/*
The drop from the baseline as calculated from the font and style
for this line.

@param[in]  metrics  The metrics.
@param[in]  line     The line index (zero based).

@return     The descent.
*/
// UNSUPPORTED :: LineMetricsGetDescent  param count = 2

/*
The y coordinate of the baseline for this line from the top of
the paragraph.

@param[in]  metrics  The metrics.
@param[in]  line     The line index (zero based).

@return     The baseline.
*/
// UNSUPPORTED :: LineMetricsGetBaseline  param count = 2

/*
Used to determine if this line ends with an explicit line break
(e.g. '\n') or is the end of the paragraph.

@param[in]  metrics  The metrics.
@param[in]  line     The line index (zero based).

@return     True if the line is a hard break.
*/
// UNSUPPORTED :: LineMetricsIsHardbreak  param count = 2

/*
Width of the line from the left edge of the leftmost glyph to
the right edge of the rightmost glyph.

@param[in]  metrics  The metrics.
@param[in]  line     The line index (zero based).

@return     The width.
*/
// UNSUPPORTED :: LineMetricsGetWidth  param count = 2

/*
Total height of the line from the top edge to the bottom edge.

@param[in]  metrics  The metrics.
@param[in]  line     The line index (zero based).

@return     The height.
*/
// UNSUPPORTED :: LineMetricsGetHeight  param count = 2

/*
The x coordinate of left edge of the line.

@param[in]  metrics  The metrics.
@param[in]  line     The line index (zero based).

@return     The left edge coordinate.
*/
// UNSUPPORTED :: LineMetricsGetLeft  param count = 2

/*
Fetch the start index in the buffer of UTF-16 code units used to
represent the paragraph line.

@param[in]  metrics  The metrics.
@param[in]  line     The line index (zero based).

@return     The UTF-16 code units start index.
*/
// UNSUPPORTED :: LineMetricsGetCodeUnitStartIndex  param count = 2

/*
Fetch the end index in the buffer of UTF-16 code units used to
represent the paragraph line.

@param[in]  metrics  The metrics.
@param[in]  line     The line index (zero based).

@return     The UTF-16 code units end index.
*/
// UNSUPPORTED :: LineMetricsGetCodeUnitEndIndex  param count = 2

/*
Fetch the end index (excluding whitespace) in the buffer of
UTF-16 code units used to represent the paragraph line.

@param[in]  metrics  The metrics.
@param[in]  line     The line index (zero based).

@return     The UTF-16 code units end index excluding whitespace.
*/
// UNSUPPORTED :: LineMetricsGetCodeUnitEndIndexExcludingWhitespace  param count = 2

/*
Fetch the end index (including newlines) in the buffer of UTF-16
code units used to represent the paragraph line.

@param[in]  metrics  The metrics.
@param[in]  line     The line index (zero based).

@return     The UTF-16 code units end index including newlines.
*/
// UNSUPPORTED :: LineMetricsGetCodeUnitEndIndexIncludingNewline  param count = 2

/*
Retain a strong reference to the object. The object can be NULL
in which case this method is a no-op.

@param[in]  glyph_info  The glyph information.
*/
// UNSUPPORTED :: GlyphInfoRetain  param count = 1

/*
Release a previously retained reference to the object. The
object can be NULL in which case this method is a no-op.

@param[in]  glyph_info  The glyph information.
*/
// UNSUPPORTED :: GlyphInfoRelease  param count = 1

/*
Fetch the start index in the buffer of UTF-16 code units used to
represent the grapheme cluster for a glyph.

@param[in]  glyph_info  The glyph information.

@return     The UTF-16 code units start index.
*/
// UNSUPPORTED :: GlyphInfoGetGraphemeClusterCodeUnitRangeBegin  param count = 1

/*
Fetch the end index in the buffer of UTF-16 code units used to
represent the grapheme cluster for a glyph.

@param[in]  glyph_info  The glyph information.

@return     The UTF-16 code units end index.
*/
// UNSUPPORTED :: GlyphInfoGetGraphemeClusterCodeUnitRangeEnd  param count = 1

/*
Fetch the bounds of the grapheme cluster for the glyph in the
coordinate space of the paragraph.

@param[in]  glyph_info  The glyph information.
@param[out] out_bounds  The grapheme cluster bounds.
*/
// UNSUPPORTED :: GlyphInfoGetGraphemeClusterBounds  param count = 2

/*
@param[in]  glyph_info  The glyph information.

@return     True if the glyph represents an ellipsis. False otherwise.
*/
// UNSUPPORTED :: GlyphInfoIsEllipsis  param count = 1

/*
@param[in]  glyph_info  The glyph information.

@return     The direction of the run that contains the glyph.
*/
// UNSUPPORTED :: GlyphInfoGetTextDirection  param count = 1
