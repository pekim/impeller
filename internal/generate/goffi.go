package generate

import (
	"github.com/dave/jennifer/jen"
)

const goffiImportPath = "github.com/go-webgpu/goffi/ffi"
const typesImportPath = "github.com/go-webgpu/goffi/types"

var pointerTypeDescriptor = jen.Qual(typesImportPath, "PointerTypeDescriptor")

func (fn function) generateGoffiVar(file file) {
	file.Var().Id(fn.varName).Qual("unsafe", "Pointer")
	file.Var().Id(fn.cifVarName).Op("=").Op("&").Qual(typesImportPath, "CallInterface").Block()
	file.Line()
}

func (fn function) generateGoffiGetSymbol(g *jen.Group) {
	g.
		List(jen.Id(fn.varName), jen.Id("err")).
		Op("=").
		Qual(goffiImportPath, "GetSymbol").Call(
		jen.Id("handle"),
		jen.Lit(fn.cName),
	)

	g.If(
		jen.Id("err").Op("!=").Nil()).Block(
		jen.Return(jen.Id("err")),
	)
}

func (fn function) generateGoffiPrepareCallInterface(g *jen.Group) {
	if !fn.supported() {
		return
	}

	g.Id("err").Op("=").Qual(goffiImportPath, "PrepareCallInterface").Call(
		jen.Line().Id(fn.cifVarName),
		jen.Line().Qual(typesImportPath, "DefaultCall"),
		jen.Line().Add(fn.result.typeDescriptor()),
		jen.Line().Index().Op("*").Qual(typesImportPath, "TypeDescriptor").ValuesFunc(func(g *jen.Group) {
			for _, param := range fn.params {
				g.Line().Add(param.typeDescriptor())
			}
			g.Line()
		}),
	)

	g.If(
		jen.Id("err").Op("!=").Nil()).Block(
		jen.Return(jen.Id("err")),
	)

	g.Line()
}

type functions []function

func (functions functions) generate() {
	functions.generateGoffi()
	functions.generateGo()
}

func (functions functions) generateGoffi() {
	file := newFile("goffi.go", "impeller")
	defer file.save()

	for _, fn := range functions {
		if fn.supported() {
			fn.generateGoffiVar(file)
		}
	}
	functions.generateGoffiInit(file)
}

func (functions functions) generateGoffiInit(file file) {
	file.Var().Id("initialised").Op("=").False()
	file.Line()

	file.Func().Id("Init").Params().Error().BlockFunc(func(g *jen.Group) {
		g.If(jen.Id("initialised")).Block(
			jen.Return(jen.Nil()),
		)
		g.Defer().Func().Params().Block(
			jen.
				Id("initialised").
				Op("=").
				True(),
		).Call()
		g.Line()

		g.
			List(
				jen.Id("filepath"), jen.Id("err"),
			).
			Op(":=").
			Id("libraryFilepath").Call()
		g.If(
			jen.Id("err").Op("!=").Nil()).Block(
			jen.Return(jen.Id("err")),
		)
		g.Line()

		g.
			List(
				jen.Id("handle"), jen.Id("err"),
			).
			Op(":=").
			Qual(goffiImportPath, "LoadLibrary").Call(
			jen.Id("filepath"),
		)
		g.If(
			jen.Id("err").Op("!=").Nil()).Block(
			jen.Return(jen.Id("err")),
		)
		g.Line()

		for _, fn := range functions {
			if fn.supported() {
				fn.generateGoffiGetSymbol(g)
				fn.generateGoffiPrepareCallInterface(g)
				g.Line()
			}
		}

		g.Return(jen.Nil())
	})
}
