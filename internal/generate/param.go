package generate

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/go-clang/clang-v15/clang"
)

type param struct {
	name string
	typ
}

func newParam(gen *gen, cursor clang.Cursor) param {
	param := param{
		name: cursor.Spelling(),
		typ:  newTyp(gen, cursor.Type()),
	}

	return param
}

func (param param) supported() (bool, string) {
	if param.isEnum {
		return true, ""
	}
	if param.isScalar {
		return true, ""
	}
	if param.isStruct && param.struct_.handle {
		return true, ""
	}
	if param.isStruct && param.isPointer {
		return true, ""
	}
	if param.isString {
		return true, ""
	}
	if param.isPointer && param.isVoid {
		return true, ""
	}

	return false, fmt.Sprintf("param %s is %q", param.name, param.typ.typ.Spelling())
}

func (param param) goDecl(g *jen.Group) {
	g.Id(param.name).Add(param.typ.goDecl())
}

func (param param) cArgVar(g *jen.Group) {
	if param.isString {
		g.Id("c_" + param.name).Op(":=").Id("cString").Call(jen.Id(param.name))
	}
}

func (param param) cArgName() jen.Code {
	if param.isEnum {
		return jen.Id(param.name)
	}
	if param.isScalar {
		return jen.Id(param.name)
	}
	if param.isStruct {
		return jen.Id(param.name)
	}
	if param.isString {
		return jen.Id("c_" + param.name)
	}
	if param.isPointer && param.isVoid {
		return jen.Id(param.name)
	}

	panic("param type")
}

type params []param

func newParams(gen *gen, cursor clang.Cursor) params {
	params := make(params, cursor.NumArguments())
	for i := range len(params) {
		params[i] = newParam(gen, cursor.Argument(uint32(i)))
	}
	return params
}

func (params params) supported() (bool, string) {
	for _, param := range params {
		if supported, reason := param.supported(); !supported {
			return supported, reason
		}
	}

	return true, ""
}

func (params params) goDecl(g *jen.Group) {
	for _, param := range params {
		param.goDecl(g)
	}
}

func (params params) cArgVars(g *jen.Group) {
	for _, param := range params {
		param.cArgVar(g)
	}
}
