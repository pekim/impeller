package generate

import (
	"github.com/dave/jennifer/jen"
	"github.com/go-clang/clang-v15/clang"
)

type param struct {
	cursor clang.Cursor
	name   string
}

func newParam(cursor clang.Cursor) param {
	return param{
		cursor: cursor,
		name:   cursor.Spelling(),
	}
}

func (param param) supported() bool {
	return false
}

func (param param) typeDescriptor() jen.Code {
	panic("TODO")
}

type params []param

func newParams(cursor clang.Cursor) params {
	params := make(params, cursor.NumArguments())
	for i := range len(params) {
		params[i] = newParam(cursor.Argument(uint32(i)))
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
