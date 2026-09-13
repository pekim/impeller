// This is a generated file. DO NOT EDIT.

package impeller

import (
	"structs"
	"unsafe"
)

/*
An Impeller graphics context. Contexts are platform and client-rendering-API
specific.

Contexts are thread-safe objects that are expensive to create. Most
applications will only ever create a single context during their lifetimes.
Once setup, Impeller is ready to render frames as performantly as possible.

During setup, context create the underlying graphics pipelines, allocators,
worker threads, etc...

The general guidance is to create as few contexts as possible (typically
just one) and share them as much as possible.
*/
type Context struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

/*
Display lists represent encoded rendering intent. These objects are
immutable, reusable, thread-safe, and context-agnostic.

While it is perfectly fine to create new display lists per frame, there may
be opportunities for optimization when display lists are reused multiple
times.
*/
type DisplayList struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

/*
Display list builders allow for the incremental creation of display lists.

Display list builders are context-agnostic.
*/
type DisplayListBuilder struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

/*
Paints control the behavior of draw calls encoded in a display list.

Like display lists, paints are context-agnostic.
*/
type Paint struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

/*
Color filters are functions that take two colors and mix them to produce a
single color. This color is then merged with the destination during
blending.
*/
type ColorFilter struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

/*
Color sources are functions that generate colors for each texture element
covered by a draw call. The colors for each element can be generated using a
mathematical function (to produce gradients for example) or sampled from a
texture.
*/
type ColorSource struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

/*
Image filters are functions that are applied regions of a texture to produce
a single color. Contrast this with color filters that operate independently
on a per-pixel basis. The generated color is then merged with the
destination during blending.
*/
type ImageFilter struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

/*
Mask filters are functions that are applied over a shape after it has been
drawn but before it has been blended into the final image.
*/
type MaskFilter struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

/*
Typography contexts allow for the layout and rendering of text.

These are typically expensive to create and applications will only ever need
to create a single one of these during their lifetimes.

Unlike graphics context, typograhy contexts are not thread-safe. These must
be created, used, and collected on a single thread.
*/
type TypographyContext struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

/*
An immutable, fully laid out paragraph.
*/
type Paragraph struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

/*
Paragraph builders allow for the creation of fully laid out paragraphs
(which themselves are immutable).

To build a paragraph, users push/pop paragraph styles onto a stack then add
UTF-8 encoded text. The properties on the top of paragraph style stack when
the text is added are used to layout and shape that subset of the paragraph.

See [ParagraphStyle]
*/
type ParagraphBuilder struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

/*
Specified when building a paragraph, paragraph styles are managed in a stack
with specify text properties to apply to text that is added to the paragraph
builder.
*/
type ParagraphStyle struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

/*
Describes the metrics of lines in a fully laid out paragraph.

Regardless of how the string of text is specified to the paragraph builder,
offsets into buffers that are returned by line metrics are always assumed to
be into buffers of UTF-16 code units.
*/
type LineMetrics struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

/*
Describes the metrics of glyphs in a paragraph line.
*/
type GlyphInfo struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

/*
Represents a two-dimensional path that is immutable and graphics context
agnostic.

Paths in Impeller consist of linear, cubic Bézier curve, and quadratic
Bézier curve segments. All other shapes are approximations using these
building blocks.

Paths are created using path builder that allow for the configuration of the
path segments, how they are filled, and/or stroked.
*/
type Path struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

/*
Path builders allow for the incremental building up of paths.
*/
type PathBuilder struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

/*
A surface represents a render target for Impeller to direct the rendering
intent specified the form of display lists to.

Render targets are how Impeller API users perform Window System Integration
(WSI). Users wrap swapchain images as surfaces and draw display lists onto
these surfaces to present content.

Creating surfaces is typically platform and client-rendering-API specific.
*/
type Surface struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

/*
A reference to a texture whose data is resident on the GPU. These can be
referenced in draw calls and paints.

Creating textures is extremely expensive. Creating a single one can
typically comfortably blow the frame budget of an application. Textures
should be created on background threads.

While textures themselves are thread safe, some context types
(like OpenGL) may need extra configuration to be able to operate
from multiple threads.
*/
type Texture struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

/*
The primary form of WSI when using a Vulkan context, these swapchains use
the `VK_KHR_surface` Vulkan extension.

Creating a swapchain is extremely expensive. One must be created at
application startup and re-used throughout the application lifecycle.

Swapchains are resilient to the underlying surfaces being resized. The
swapchain images will be re-created as necessary on-demand.
*/
type VulkanSwapchain struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

/*
A fragment shader is a small program that is authored in GLSL and compiled
using `impellerc` that runs on each pixel covered by a polygon and allows
the user to configure how it is shaded.

See https://docs.flutter.dev/ui/design/graphics/fragment-shaders
*/
type FragmentProgram struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

// Non-opaque structs
// -----------------------------------------------------------------------------

type Rect struct {
	_      structs.HostLayout
	X      float32
	Y      float32
	Width  float32
	Height float32
}

type Point struct {
	_ structs.HostLayout
	X float32
	Y float32
}

type Size struct {
	_      structs.HostLayout
	Width  float32
	Height float32
}

type ISize struct {
	_      structs.HostLayout
	Width  int64
	Height int64
}

type Range struct {
	_     structs.HostLayout
	Start uint64
	End   uint64
}

/*
A 4x4 transformation matrix using column-major storage.

	```
	| m[0] m[4] m[8]  m[12] |
	| m[1] m[5] m[9]  m[13] |
	| m[2] m[6] m[10] m[14] |
	| m[3] m[7] m[11] m[15] |

```
*/
type Matrix struct {
	_ structs.HostLayout
	M [16]float32
}

/*
A 4x5 matrix using row-major storage used for transforming color values.

To transform color values, a 5x5 matrix is constructed with the 5th row
being identity. Then the following transformation is performed:

	```
	| R' |   | m[0]  m[1]  m[2]  m[3]  m[4]  |   | R |
	| G' |   | m[5]  m[6]  m[7]  m[8]  m[9]  |   | G |
	| B' | = | m[10] m[11] m[12] m[13] m[14] | * | B |
	| A' |   | m[15] m[16] m[17] m[18] m[19] |   | A |
	| 1  |   | 0     0     0     0     1     |   | 1 |

```

The translation column (m[4], m[9], m[14], m[19]) must be specified in
non-normalized 8-bit unsigned integer space (0 to 255). Values outside this
range will produce undefined results.

The identity transformation is thus:

	```
	1, 0, 0, 0, 0,
	0, 1, 0, 0, 0,
	0, 0, 1, 0, 0,
	0, 0, 0, 1, 0,

```

Some examples:

To invert all colors:

	```
	-1,  0,  0, 0, 255,
	0, -1,  0, 0, 255,
	0,  0, -1, 0, 255,
	0,  0,  0, 1,   0,

```

To apply a sepia filter:

	```
	0.393, 0.769, 0.189, 0, 0,
	0.349, 0.686, 0.168, 0, 0,
	0.272, 0.534, 0.131, 0, 0,
	0,     0,     0,     1, 0,

```

To apply a grayscale conversion filter:

	```
	0.2126, 0.7152, 0.0722, 0, 0,
	0.2126, 0.7152, 0.0722, 0, 0,
	0.2126, 0.7152, 0.0722, 0, 0,
	0,      0,      0,      1, 0,

```

See [ColorFilter]
*/
type ColorMatrix struct {
	_ structs.HostLayout
	M [20]float32
}

type RoundingRadii struct {
	_           structs.HostLayout
	TopLeft     Point
	BottomLeft  Point
	TopRight    Point
	BottomRight Point
}

type Color struct {
	_          structs.HostLayout
	Red        float32
	Green      float32
	Blue       float32
	Alpha      float32
	ColorSpace ColorSpace
}

type TextureDescriptor struct {
	_           structs.HostLayout
	PixelFormat PixelFormat
	Size        ISize
	MipCount    uint32
}

type Mapping struct {
	_         structs.HostLayout
	Data      *byte
	Length    uint64
	OnRelease Callback
}

type ContextVulkanSettings struct {
	_                      structs.HostLayout
	UserData               *byte
	ProcAddressCallback    VulkanProcAddressCallback
	EnableVulkanValidation Bool
}

type ContextVulkanInfo struct {
	_                        structs.HostLayout
	VkInstance               *byte
	VkPhysicalDevice         *byte
	VkLogicalDevice          *byte
	GraphicsQueueFamilyIndex uint32
	GraphicsQueueIndex       uint32
}

type TextDecoration struct {
	_ structs.HostLayout
	/*
	   A mask of `ImpellerTextDecorationType`s to enable.
	*/
	Types int32
	/*
	   The decoration color.
	*/
	Color Color
	/*
	   The decoration style.
	*/
	Style TextDecorationStyle
	// The multiplier applied to the default thickness of the font to use for the
	// decoration.

	ThicknessMultiplier float32
}
