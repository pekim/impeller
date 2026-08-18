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
		params:     newParams(gen, cursor),
		varName:    "func" + cName,
		cifVarName: "cif" + cName,
	}
}

func (fn function) generate(file file) {
	if supported, reason := fn.supported(); !supported {
		file.Commentf("UNSUPPORTED %s : %s", fn.name, reason)
		return
	}

	file.Comment(fn.gen.commentText(fn.cursor))

	file.
		Func().
		Id(fn.name).
		ParamsFunc(fn.params.goDecl).
		Add(fn.result.goDecl()).
		BlockFunc(func(g *jen.Group) {
			fn.result.resultVar(g)

			g.
				List(jen.Id("_"), jen.Id("err")).
				Op(":=").
				Qual(goffiImportPath, "CallFunction").
				CallFunc(func(g *jen.Group) {
					g.Line().Id(fn.cifVarName)
					g.Line().Id(fn.varName)
					g.Line().Qual("unsafe", "Pointer").Parens(fn.result.returnValuePointer())
					g.Line().Index().Qual("unsafe", "Pointer").ValuesFunc(func(g *jen.Group) {
						for _, param := range fn.params {
							g.Line().Qual("unsafe", "Pointer").Parens(jen.Op("&").Add(param.cArgName()))
						}
						g.Line()
					})
					g.Line()
				})

			g.If(jen.Id("err").Op("!=").Nil()).Block(
				jen.Panic(jen.Id("err")),
			)

			fn.result.returnVar(g)
		})
}

func (fn function) supported() (bool, string) {
	if supported, reason := fn.result.supported(); !supported {
		return supported, reason
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
