package generate

import (
	"github.com/dave/jennifer/jen"
	"github.com/go-clang/clang-v15/clang"
)

var scalars = map[clang.CursorKind]struct {
	goType jen.Code
}{
	clang.Type_Bool: {
		goType: jen.Id("Bool"),
	},

	clang.Type_Float: {
		goType: jen.Float32(),
	},

	clang.Type_Int: {
		goType: jen.Int32(),
	},

	clang.Type_Long: {
		goType: jen.Int64(),
	},

	clang.Type_UInt: {
		goType: jen.Uint32(),
	},

	clang.Type_ULong: {
		goType: jen.Uint64(),
	},
}
