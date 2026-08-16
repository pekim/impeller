// This is a generated file. DO NOT EDIT.

package impeller

import "structs"

type Context *struct{}

type DisplayList *struct{}

type DisplayListBuilder *struct{}

type Paint *struct{}

type ColorFilter *struct{}

type ColorSource *struct{}

type ImageFilter *struct{}

type MaskFilter *struct{}

type TypographyContext *struct{}

type Paragraph *struct{}

type ParagraphBuilder *struct{}

type ParagraphStyle *struct{}

type LineMetrics *struct{}

type GlyphInfo *struct{}

type Path *struct{}

type PathBuilder *struct{}

type Surface *struct{}

type Texture *struct{}

type VulkanSwapchain *struct{}

type FragmentProgram *struct{}

type Rect struct {
	_ structs.HostLayout

	// X
	// Y
	// Width
	// Height
}

type Point struct {
	_ structs.HostLayout

	// X
	// Y
}

type Size struct {
	_ structs.HostLayout

	// Width
	// Height
}

type ISize struct {
	_ structs.HostLayout

	// Width
	// Height
}

type Range struct {
	_ structs.HostLayout

	// Start
	// End
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

	// M
}

/*
A 4x5 matrix using row-major storage used for transforming color values.

To transform color values, a 5x5 matrix is constructed with the 5th row
being identity. Then the following transformation is performed:

	| R' |   | m[0]  m[1]  m[2]  m[3]  m[4]  |   | R |
	| G' |   | m[5]  m[6]  m[7]  m[8]  m[9]  |   | G |
	| B' | = | m[10] m[11] m[12] m[13] m[14] | * | B |
	| A' |   | m[15] m[16] m[17] m[18] m[19] |   | A |
	| 1  |   | 0     0     0     0     1     |   | 1 |

The translation column (m[4], m[9], m[14], m[19]) must be specified in
non-normalized 8-bit unsigned integer space (0 to 255). Values outside this
range will produce undefined results.

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

@see      ImpellerColorFilter
*/
type ColorMatrix struct {
	_ structs.HostLayout

	// M
}

type RoundingRadii struct {
	_ structs.HostLayout

	// TopLeft
	// BottomLeft
	// TopRight
	// BottomRight
}

type Color struct {
	_ structs.HostLayout

	// Red
	// Green
	// Blue
	// Alpha
	// ColorSpace
}

type TextureDescriptor struct {
	_ structs.HostLayout

	// PixelFormat
	// Size
	// MipCount
}

type Mapping struct {
	_ structs.HostLayout

	// Data
	// Length
	// OnRelease
}

type ContextVulkanSettings struct {
	_ structs.HostLayout

	// UserData
	// ProcAddressCallback
	// EnableVulkanValidation
}

type ContextVulkanInfo struct {
	_ structs.HostLayout

	// VkInstance
	// VkPhysicalDevice
	// VkLogicalDevice
	// GraphicsQueueFamilyIndex
	// GraphicsQueueIndex
}

type TextDecoration struct {
	_ structs.HostLayout

	// Types
	// Color
	// Style
	// ThicknessMultiplier
}
