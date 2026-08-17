package generate

import (
	"fmt"

	"github.com/dave/jennifer/jen"
)

const goffiImportPath = "github.com/go-webgpu/goffi/ffi"
const typesImportPath = "github.com/go-webgpu/goffi/types"

func (fn function) generateGoffiVar(file file) {
	file.Var().Id(fn.cName).Qual("unsafe", "Pointer")
	file.Var().Id(fn.cName+"_cif").Op("=").Op("&").Qual(typesImportPath, "CallInterface").Block()
}

func (fn function) generateGoffiGetSymbol(g *jen.Group) {
	g.
		List(jen.Id(fn.cName), jen.Id("err")).
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

	fmt.Println(fn.cName)
	g.Id("err").Op("=").Qual(goffiImportPath, "PrepareCallInterface").Call(
		jen.Line().Id(fn.cName+"_cif"),
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
		fn.generateGoffiVar(file)
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
			fn.generateGoffiGetSymbol(g)
		}
		g.Line()

		for _, fn := range functions {
			fn.generateGoffiPrepareCallInterface(g)
		}
		g.Line()

		g.Return(jen.Nil())
	})
}
