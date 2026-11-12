// This is a generated file. DO NOT EDIT.

package impeller

import (
	"structs"
	"unsafe"
)

type Context struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

type DisplayList struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

type DisplayListBuilder struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

type Paint struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

type ColorFilter struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

type ColorSource struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

type ImageFilter struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

type MaskFilter struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

type TypographyContext struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

type Paragraph struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

type ParagraphBuilder struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

type ParagraphStyle struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

type LineMetrics struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

type GlyphInfo struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

type Path struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

type PathBuilder struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

type Surface struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

type Texture struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

type VulkanSwapchain struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

type FragmentProgram struct {
	_      structs.HostLayout
	handle unsafe.Pointer
}

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

	| m[0] m[4] m[8]  m[12] |
	| m[1] m[5] m[9]  m[13] |
	| m[2] m[6] m[10] m[14] |
	| m[3] m[7] m[11] m[15] |
*/
type Matrix struct {
	_ structs.HostLayout
	M [16]float32
}

/*
A 4x5 matrix using row-major storage used for transforming color values.

To transform color values, a 5x5 matrix is constructed with the 5th row being identity. Then the following transformation is performed:

	| R' |   | m[0]  m[1]  m[2]  m[3]  m[4]  |   | R |
	| G' |   | m[5]  m[6]  m[7]  m[8]  m[9]  |   | G |
	| B' | = | m[10] m[11] m[12] m[13] m[14] | * | B |
	| A' |   | m[15] m[16] m[17] m[18] m[19] |   | A |
	| 1  |   | 0     0     0     0     1     |   | 1 |

The translation column (m[4], m[9], m[14], m[19]) must be specified in non-normalized 8-bit unsigned integer space (0 to 255). Values outside this range will produce undefined results.

The identity transformation is thus:

	1, 0, 0, 0, 0,
	0, 1, 0, 0, 0,
	0, 0, 1, 0, 0,
	0, 0, 0, 1, 0,

Some examples:

To invert all colors:

	-1,  0,  0, 0, 255,
	0, -1,  0, 0, 255,
	0,  0, -1, 0, 255,
	0,  0,  0, 1,   0,

To apply a sepia filter:

	0.393, 0.769, 0.189, 0, 0,
	0.349, 0.686, 0.168, 0, 0,
	0.272, 0.534, 0.131, 0, 0,
	0,     0,     0,     1, 0,

To apply a grayscale conversion filter:

	0.2126, 0.7152, 0.0722, 0, 0,
	0.2126, 0.7152, 0.0722, 0, 0,
	0.2126, 0.7152, 0.0722, 0, 0,
	0,      0,      0,      1, 0,

see [ColorFilter]
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
	   A mask of [TextDecorationType]s to enable.
	*/
	Types int32
	/*
	   The decoration color.
	*/
	Color Color
	/*
	   The decoration style.
	*/
	Style               TextDecorationStyle
	ThicknessMultiplier float32
}
