// This is a generated file. DO NOT EDIT.

package structtest

import impeller "github.com/pekim/impeller"

// #cgo CFLAGS: --std=c11
// #include "../interop/include/impeller.h"
import "C"

var structs = []struct {
	name string
	c    any
	go_  any
}{{
	c:    C.ImpellerRect{},
	go_:  impeller.Rect{},
	name: "ImpellerRect",
}, {
	c:    C.ImpellerPoint{},
	go_:  impeller.Point{},
	name: "ImpellerPoint",
}, {
	c:    C.ImpellerSize{},
	go_:  impeller.Size{},
	name: "ImpellerSize",
}, {
	c:    C.ImpellerISize{},
	go_:  impeller.ISize{},
	name: "ImpellerISize",
}, {
	c:    C.ImpellerRange{},
	go_:  impeller.Range{},
	name: "ImpellerRange",
}, {
	c:    C.ImpellerMatrix{},
	go_:  impeller.Matrix{},
	name: "ImpellerMatrix",
}, {
	c:    C.ImpellerColorMatrix{},
	go_:  impeller.ColorMatrix{},
	name: "ImpellerColorMatrix",
}, {
	c:    C.ImpellerRoundingRadii{},
	go_:  impeller.RoundingRadii{},
	name: "ImpellerRoundingRadii",
}, {
	c:    C.ImpellerColor{},
	go_:  impeller.Color{},
	name: "ImpellerColor",
}, {
	c:    C.ImpellerTextureDescriptor{},
	go_:  impeller.TextureDescriptor{},
	name: "ImpellerTextureDescriptor",
}, {
	c:    C.ImpellerMapping{},
	go_:  impeller.Mapping{},
	name: "ImpellerMapping",
}, {
	c:    C.ImpellerContextVulkanSettings{},
	go_:  impeller.ContextVulkanSettings{},
	name: "ImpellerContextVulkanSettings",
}, {
	c:    C.ImpellerContextVulkanInfo{},
	go_:  impeller.ContextVulkanInfo{},
	name: "ImpellerContextVulkanInfo",
}, {
	c:    C.ImpellerTextDecoration{},
	go_:  impeller.TextDecoration{},
	name: "ImpellerTextDecoration",
},
}
