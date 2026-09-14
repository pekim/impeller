package generate

import (
	"unicode"

	"github.com/dave/jennifer/jen"
	"github.com/go-clang/clang-v15/clang"
)

type struct_ struct {
	gen     *gen
	cursor  clang.Cursor
	cName   string
	name    string
	comment comment
	handle  bool // defined using the IMPELLER_DEFINE_HANDLE macro
}

func (struct_ struct_) generate(file file) {
	file.Comment(struct_.comment.text())
	file.Type().Id(struct_.name).StructFunc(func(g *jen.Group) {
		g.Id("_").Qual("structs", "HostLayout")

		if struct_.handle {
			g.Id("handle").Qual("unsafe", "Pointer")
		} else {
			struct_.cursor.Visit(func(cursor, _parent clang.Cursor) (status clang.ChildVisitResult) {
				if cursor.Kind() == clang.Cursor_FieldDecl {
					struct_.generateField(g, cursor)
				}

				return clang.ChildVisit_Continue
			})
		}
	})
}

func (struct_ struct_) generateField(g *jen.Group, cursor clang.Cursor) {
	cName := cursor.Spelling()
	name := goName(cName)
	typ := cursor.Type().Spelling()

	comment := struct_.gen.newComment(cursor)
	comment.resolve()
	g.Comment(comment.text())

	if scalar, ok := scalars[clang.CursorKind(cursor.Type().CanonicalType().Kind())]; ok {
		g.Id(name).Add(scalar.goType)

	} else if unicode.IsUpper([]rune(typ)[0]) {
		g.Id(name).Id(goName(typ))

	} else if cursor.Type().Kind() == clang.Type_Pointer {
		g.Id(name).Op("*").Byte()

	} else if cursor.Type().Kind() == clang.Type_ConstantArray {
		arraySize := int(cursor.Type().ArraySize())
		elementType := cursor.Type().ArrayElementType()
		if scalar, ok := scalars[clang.CursorKind(elementType.Kind())]; ok {
			g.Id(name).Index(jen.Lit(arraySize)).Add(scalar.goType)

		} else {
			fatalf("UNSUPPORTED : field %s with array type %q", goName(cursor.Spelling()), elementType.Spelling())
		}

	} else {
		fatalf("UNSUPPORTED : field %s of type %q", goName(cursor.Spelling()), typ)
	}
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

func (structs structs) findByGoName(name string) (*struct_, bool) {
	for i, struct_ := range structs {
		if struct_.name == name {
			return &(structs[i]), true
		}
	}

	return nil, false
}

func (structs structs) resolve() {
	for i := range structs {
		(&structs[i]).comment.resolve()

		// for i,field:=range &structs[i].f
	}
}
