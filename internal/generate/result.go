package generate

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/go-clang/clang-v15/clang"
)

type result struct {
	typ
}

func newResult(gen *gen, typ clang.Type) result {
	return result{
		typ: newTyp(gen, typ),
	}
}

func (result result) supported() (bool, string) {
	if result.isVoid {
		return true, ""
	}
	if result.isEnum {
		return true, ""
	}
	if result.isScalar {
		return true, ""
	}
	if result.isStruct && result.struct_.handle {
		return true, ""
	}

	return false, fmt.Sprintf("result type is %q", result.typ.typ.Spelling())
}

func (result result) resultVar(g *jen.Group) {
	if result.isVoid {
		return

	} else if result.isEnum {
		g.Var().Id("result").Id(result.enum.name)

	} else if result.isScalar {
		g.Var().Id("result").Add(result.scalar.goType)

	} else if result.isStruct {
		g.Var().Id("result").Id(result.struct_.name)

	} else {
		panic("unhandled type")
	}
}

func (result result) returnVar(g *jen.Group) {
	if result.isVoid {
		return
	}

	g.Return().Id("result")
}

func (result result) returnValuePointer() jen.Code {
	if result.isVoid {
		return jen.Nil()
	}

	return jen.Op("&").Id("result")
}
