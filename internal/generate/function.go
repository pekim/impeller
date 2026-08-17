package generate

import (
	"github.com/dave/jennifer/jen"
	"github.com/go-clang/clang-v15/clang"
)

type function struct {
	gen    *gen
	cursor clang.Cursor
	cName  string
	name   string
	result result
	params params
}

func newFunction(gen *gen, cursor clang.Cursor) function {
	cName := cursor.Spelling()
	return function{
		gen:    gen,
		cursor: cursor,
		cName:  cName,
		name:   goName(cName),
		result: result{typ: cursor.ResultType()},
		params: newParams(cursor),
	}
}

func (fn function) generate(file file) {
	file.Comment(fn.gen.commentText(fn.cursor))

	if !fn.supported() {
		file.Commentf("UNSUPPORTED :: %s  param count = %d", fn.name, fn.cursor.NumArguments())
		return
	}

	file.
		Func().
		Id(fn.name).
		Params().
		Add(fn.result.goDecl()).
		BlockFunc(func(g *jen.Group) {
			g.Qual("fmt", "Println").Call(jen.Lit("TODO"))
			g.Return(jen.Lit(42))

			/*
				var result uint32
					_, err := ffi.CallFunction(ImpellerGetVersion_cif, ImpellerGetVersion,
					unsafe.Pointer(&result),
					[]unsafe.Pointer{},
				)
				if err != nil {
					panic(err)
				}
				return result
			*/
		})
}

func (fn function) supported() bool {
	if !fn.result.supported() {
		return false
	}
	return fn.params.supported()
}

func (functions functions) generateGo() {
	file := newFile("function.go", "impeller")
	defer file.save()

	for _, function := range functions {
		function.generate(file)
		file.Line()
	}
}
