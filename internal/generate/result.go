package generate

import (
	"github.com/dave/jennifer/jen"
	"github.com/go-clang/clang-v15/clang"
)

type result struct {
	typ clang.Type
}

func (result result) supported() bool {
	if result.typ.CanonicalType().Kind() == clang.Type_UInt {
		return true
	}

	return false
}

func (result result) goDecl() jen.Code {
	if scalar, ok := scalars[clang.CursorKind(result.typ.CanonicalType().Kind())]; ok {
		return scalar.goType
	}

	panic("result type")
}

func (result result) typeDescriptor() jen.Code {
	if scalar, ok := scalars[clang.CursorKind(result.typ.CanonicalType().Kind())]; ok {
		return scalar.typeDescriptor
	}

	panic("result type")
}
