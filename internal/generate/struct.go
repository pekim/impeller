package generate

import (
	"github.com/dave/jennifer/jen"
	"github.com/go-clang/clang-v15/clang"
)

type struct_ struct {
	gen     *gen
	cursor  clang.Cursor
	cName   string
	name    string
	comment string
	handle  bool // defined using the IMPELLER_DEFINE_HANDLE macro
}

func (struct_ struct_) generate(file file) {
	pointer := jen.Null()
	if struct_.handle {
		pointer = jen.Op("*")
	}

	file.Comment(struct_.comment)
	file.Type().Id(struct_.name).Add(pointer).StructFunc(func(g *jen.Group) {
		if struct_.handle {
			return
		}

		g.Id("_").Qual("structs", "HostLayout")
		g.Line()

		struct_.cursor.Visit(func(cursor, _parent clang.Cursor) (status clang.ChildVisitResult) {
			if cursor.Kind() == clang.Cursor_FieldDecl {
				struct_.generateField(g, cursor)
			}

			return clang.ChildVisit_Continue
		})
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

func (structs structs) find(cName string) (*struct_, bool) {
	for i, struct_ := range structs {
		if struct_.cName == cName {
			return &(structs[i]), true
		}
	}

	return nil, false
}
