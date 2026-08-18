package generate

import (
	"github.com/dave/jennifer/jen"
	"github.com/go-clang/clang-v15/clang"
)

type typ struct {
	typ      clang.Type
	isVoid   bool
	isScalar bool
	scalar   scalar
	isStruct bool
	struct_  *struct_
}

func newTyp(gen *gen, typ_ clang.Type) typ {
	typ := typ{
		typ: typ_,
	}

	typ.isVoid = typ.typ.CanonicalType().Kind() == clang.Type_Void
	typ.scalar, typ.isScalar = scalars[clang.CursorKind(typ.typ.CanonicalType().Kind())]
	typ.struct_, typ.isStruct = gen.structs.findByGoName(goName(typ.typ.Spelling()))

	return typ
}

func (typ typ) goDecl() jen.Code {
	if typ.isVoid {
		return jen.Null()
	}

	if typ.isScalar {
		return typ.scalar.goType
	}

	if typ.isStruct {
		return jen.Id(typ.struct_.name)
	}

	panic("unhandled type")
}

func (typ typ) goffiVar(g *jen.Group) {
	if typ.isVoid {
		return
	}

	if typ.isScalar {
		g.Var().Id("result").Add(typ.scalar.goType)

	} else if typ.isStruct {
		g.Var().Id("result").Id(typ.struct_.name)

	} else {
		panic("unhandled type")
	}
}

func (typ typ) typeDescriptor() jen.Code {
	if typ.isVoid {
		return voidTypeDescriptor
	}
	if typ.isScalar {
		return typ.scalar.typeDescriptor
	}
	if typ.isStruct && typ.struct_.handle {
		return pointerTypeDescriptor
	}

	panic("unhandled type")
}
