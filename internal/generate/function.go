package generate

import (
	"strings"

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
	comment    comment
}

func newFunction(gen *gen, cursor clang.Cursor) function {
	cName := cursor.Spelling()
	comment := gen.newComment(cursor)
	fn := function{
		gen:        gen,
		cursor:     cursor,
		cName:      cName,
		name:       goName(cName),
		result:     newResult(gen, cursor.ResultType()),
		params:     newParams(gen, cursor, comment.params),
		varName:    "func" + cName,
		cifVarName: "cif" + cName,
		comment:    comment,
	}

	if fn.method() {
		fn.name = strings.TrimPrefix(fn.name, fn.params[0].struct_.name)
	}

	return fn
}

func (fn function) generate(file file) {
	if supported, reason := fn.supported(); !supported {
		file.Commentf("UNSUPPORTED %s : %s", fn.name, reason)
		return
	}

	file.Comment(fn.comment.text())

	file.
		Func().
		Do(func(s *jen.Statement) { // receiver
			if fn.method() {
				s.Parens(jen.Add(fn.params[0].goDecl()))
			}
		}).
		Id(fn.name).
		ParamsFunc(func(g *jen.Group) { // params
			if fn.method() {
				fn.params[1:].goDecl(g)
			} else {
				fn.params.goDecl(g)
			}
		}).
		// result type(s)
		ParamsFunc(func(g *jen.Group) {
			fn.params.outDecls(g)

			if !fn.result.isVoid {
				g.Add(fn.result.goDecl())
			}
		}).
		BlockFunc(func(g *jen.Group) {
			fn.result.resultVar(g)
			fn.params.cArgVars(g)
			fn.params.outVars(g)
			fn.params.lengthVar(g)

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

			if !fn.result.isVoid || fn.params.haveOut() {
				g.Return().ListFunc(func(g *jen.Group) {
					for _, param := range fn.params {
						if param.direction == out {
							g.Op("*").Id(param.name)
						}
					}

					if !fn.result.isVoid {
						g.Id("result")
					}
				})
			}
		})
}

func (fn function) method() bool {
	return len(fn.params) > 0 && fn.params[0].isStruct && fn.params[0].struct_.handle
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

func (functions functions) find(cName string) (*function, bool) {
	for i, fn := range functions {
		if fn.cName == cName {
			return &(functions[i]), true
		}
	}

	return nil, false
}

func (functions functions) resolve() {
	for i := range functions {
		(&functions[i]).comment.resolve()
	}
}
