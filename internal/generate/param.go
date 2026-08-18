package generate

import (
	"github.com/dave/jennifer/jen"
	"github.com/go-clang/clang-v15/clang"
)

type param struct {
	cursor clang.Cursor
	name   string
	typ
}

func newParam(gen *gen, cursor clang.Cursor) param {
	param := param{
		cursor: cursor,
		name:   cursor.Spelling(),
		typ:    newTyp(gen, cursor.Type()),
	}

	return param
}

func (param param) supported() bool {
	if param.isScalar {
		return true
	}
	if param.isStruct && param.struct_.handle {
		return true
	}

	return false
}

func (param param) goDecl(g *jen.Group) {
	g.Id(param.name).Add(param.typ.goDecl())
}

func (param param) cArgName() jen.Code {
	if param.isScalar {
		return jen.Id(param.name)
	}

	if param.isStruct {
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

func (params params) supported() bool {
	for _, param := range params {
		if !param.supported() {
			return false
		}
	}

	return true
}

func (params params) goDecl(g *jen.Group) {
	for _, param := range params {
		param.goDecl(g)
	}
}
