package generate

import (
	"github.com/dave/jennifer/jen"
	"github.com/go-clang/clang-v15/clang"
)

type function struct {
	gen        *gen
	cursor     clang.Cursor
	cName      string
	name       string
	result     result
	params     params
	varName    string
	cifVarName string
}

func newFunction(gen *gen, cursor clang.Cursor) function {
	cName := cursor.Spelling()
	return function{
		gen:        gen,
		cursor:     cursor,
		cName:      cName,
		name:       goName(cName),
		result:     newResult(gen, cursor.ResultType()),
		params:     newParams(cursor),
		varName:    "_" + cName,          // don't export the var
		cifVarName: "_" + cName + "_cif", // don't export the var
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
			fn.result.goffiVar(g)

			g.
				List(jen.Id("_"), jen.Id("err")).
				Op(":=").
				Qual(goffiImportPath, "CallFunction").
				CallFunc(func(g *jen.Group) {
					g.Line().Id(fn.cifVarName)
					g.Line().Id(fn.varName)
					g.Line().Qual("unsafe", "Pointer").Parens(fn.result.returnValuePointer())
					g.Line().Index().Qual("unsafe", "Pointer").ValuesFunc(func(_g *jen.Group) {

					})
					g.Line()
				})

			g.If(jen.Id("err").Op("!=").Nil()).Block(
				jen.Panic(jen.Id("err")),
			)

			fn.result.returnVar(g)
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
