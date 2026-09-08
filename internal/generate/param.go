package generate

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/go-clang/clang-v15/clang"
)

type param struct {
	typ
	name          string
	lengthParam   *param
	isLengthParam bool
}

func newParam(gen *gen, cursor clang.Cursor) param {
	param := param{
		name: cursor.Spelling(),
		typ:  newTyp(gen, cursor.Type()),
	}

	return param
}

func (param param) supported() (bool, string) {
	if param.isCallback {
		return true, ""
	}
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

func (param param) goDecl() jen.Code {
	return jen.Id(param.name).Add(param.typ.goDecl())
}

func (param param) cArgVar(g *jen.Group) {
	if param.isString {
		g.Id("c_" + param.name).Op(":=").Id("cString").Call(jen.Id(param.name))
	}
}

func (param param) cArgName() jen.Code {
	if param.isCallback {
		return jen.Id(param.name)
	}
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

	for i := range len(params) - 1 {
		param1 := params[i]
		param2 := params[i+1]

		if param1.isString && param2.isScalar && strings.HasSuffix(param2.name, "length") {
			params[i].lengthParam = &params[i+1]
			params[i+1].isLengthParam = true
		}
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
		if param.isLengthParam {
			continue
		}
		g.Add(param.goDecl())
	}
}

func (params params) cArgVars(g *jen.Group) {
	for _, param := range params {
		param.cArgVar(g)
	}
}

func (params params) lengthVar(g *jen.Group) {
	for _, param := range params {
		if param.lengthParam == nil {
			continue
		}
		g.Id(param.lengthParam.name).Op(":=").Len(jen.Id(param.name))
	}
}
