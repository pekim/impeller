package generate

import (
	"github.com/dave/jennifer/jen"
	"github.com/go-clang/clang-v15/clang"
)

type scalar struct {
	goType         jen.Code
	typeDescriptor jen.Code
}

var scalars = map[clang.CursorKind]scalar{
	clang.Type_Bool: {
		goType:         jen.Id("Bool"),
		typeDescriptor: jen.Qual(typesImportPath, "UInt8TypeDescriptor"),
	},

	clang.Type_Float: {
		goType:         jen.Float32(),
		typeDescriptor: jen.Qual(typesImportPath, "FloatTypeDescriptor"),
	},

	clang.Type_Double: {
		goType:         jen.Float64(),
		typeDescriptor: jen.Qual(typesImportPath, "DoubleTypeDescriptor"),
	},

	clang.Type_Int: {
		goType:         jen.Int32(),
		typeDescriptor: jen.Qual(typesImportPath, "SInt32TypeDescriptor"),
	},

	clang.Type_Long: {
		goType:         jen.Int64(),
		typeDescriptor: jen.Qual(typesImportPath, "SInt64TypeDescriptor"),
	},

	clang.Type_UInt: {
		goType:         jen.Uint32(),
		typeDescriptor: jen.Qual(typesImportPath, "UInt32TypeDescriptor"),
	},

	clang.Type_ULong: {
		goType:         jen.Uint64(),
		typeDescriptor: jen.Qual(typesImportPath, "UInt64TypeDescriptor"),
	},
}

type typ struct {
	typ      clang.Type
	isEnum   bool
	enum     *enum
	isScalar bool
	scalar   scalar
	isStruct bool
	struct_  *struct_
	isVoid   bool
}

func newTyp(gen *gen, typ_ clang.Type) typ {
	typ := typ{
		typ: typ_,
	}

	typ.enum, typ.isEnum = gen.enums.findByGoName(goName(typ.typ.Spelling()))
	typ.scalar, typ.isScalar = scalars[clang.CursorKind(typ.typ.CanonicalType().Kind())]
	typ.struct_, typ.isStruct = gen.structs.findByGoName(goName(typ.typ.Spelling()))
	typ.isVoid = typ.typ.CanonicalType().Kind() == clang.Type_Void

	return typ
}

func (typ typ) goDecl() jen.Code {
	if typ.isVoid {
		return jen.Null()
	}
	if typ.isEnum {
		return jen.Id(typ.enum.name)
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

	} else if typ.isEnum {
		g.Var().Id("result").Id(typ.enum.name)

	} else if typ.isScalar {
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
	if typ.isEnum {
		return enumTypeDescriptor
	}
	if typ.isScalar {
		return typ.scalar.typeDescriptor
	}
	if typ.isStruct && typ.struct_.handle {
		return pointerTypeDescriptor
	}

	panic("unhandled type")
}
