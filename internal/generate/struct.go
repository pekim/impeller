package generate

import (
	"github.com/dave/jennifer/jen"
	"github.com/go-clang/clang-v15/clang"
)

type struct_ struct {
	gen    *gen
	name   string
	cursor clang.Cursor
}

func (struct_ struct_) generate(file file) {
	file.Comment(struct_.gen.commentText(struct_.cursor))
	file.Type().Id(struct_.name).StructFunc(func(_g *jen.Group) {

	})
}

type structs []struct_

func (structs structs) generate() {
	file := newFile("struct.go", "impeller")
	defer file.save()

	for _, struct_ := range structs {
		struct_.generate(file)
		file.Line()
	}
}
