package generate

import (
	"github.com/dave/jennifer/jen"
	"github.com/go-clang/clang-v15/clang"
)

type result struct {
	typ      clang.Type
	void     bool
	isScalar bool
	scalar   scalar
}

func newResult(typ clang.Type) result {
	result := result{
		typ:  typ,
		void: typ.CanonicalType().Kind() == clang.Type_Void,
	}
	result.scalar, result.isScalar = scalars[clang.CursorKind(result.typ.CanonicalType().Kind())]

	return result
}

func (result result) supported() bool {
	if result.void || result.isScalar {
		return true
	}

	return false
}

func (result result) goDecl() jen.Code {
	if result.void {
		return jen.Null()
	}

	if result.isScalar {
		return result.scalar.goType
	}

	panic("result type")
}

func (result result) goffiVar(g *jen.Group) {
	if result.void {
		return
	}

	if result.isScalar {
		g.Var().Id("result").Add(result.scalar.goType)

	} else {
		panic("result type")
	}
}

func (result result) returnVar(g *jen.Group) {
	if result.void {
		return
	}

	if result.isScalar {
		g.Return().Id("result")

	} else {
		panic("result type")
	}
}

func (result result) returnValuePointer() jen.Code {
	if result.void {
		return jen.Nil()
	}

	if result.isScalar {
		return jen.Op("&").Id("result")
	}

	panic("result type")
}

func (result result) typeDescriptor() jen.Code {
	if result.isScalar {
		return result.scalar.typeDescriptor
	}

	panic("result type")
}
