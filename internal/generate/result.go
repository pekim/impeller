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
	isStruct bool
	struct_  *struct_
}

func newResult(gen *gen, typ clang.Type) result {
	result := result{
		typ:  typ,
		void: typ.CanonicalType().Kind() == clang.Type_Void,
	}

	result.scalar, result.isScalar = scalars[clang.CursorKind(result.typ.CanonicalType().Kind())]
	result.struct_, result.isStruct = gen.structs.findByGoName(goName(typ.Spelling()))

	return result
}

func (result result) supported() bool {
	if result.void || result.isScalar || result.isStruct {
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

	if result.isStruct {
		return jen.Id(result.struct_.name)
	}

	panic("result type")
}

func (result result) goffiVar(g *jen.Group) {
	if result.void {
		return
	}

	if result.isScalar {
		g.Var().Id("result").Add(result.scalar.goType)

	} else if result.isStruct {
		g.Var().Id("result").Id(result.struct_.name)

	} else {
		panic("result type")
	}
}

func (result result) returnVar(g *jen.Group) {
	if result.void {
		return
	}

	g.Return().Id("result")
}

func (result result) returnValuePointer() jen.Code {
	if result.void {
		return jen.Nil()
	}

	return jen.Op("&").Id("result")
}

func (result result) typeDescriptor() jen.Code {
	if result.isScalar {
		return result.scalar.typeDescriptor
	}
	if result.isStruct && result.struct_.handle {
		return pointerTypeDescriptor
	}

	panic("result type")
}
