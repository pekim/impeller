package generate

import (
	"unicode"

	"github.com/dave/jennifer/jen"
	"github.com/go-clang/clang-v15/clang"
)

func (struct_ struct_) generateField(g *jen.Group, cursor clang.Cursor) {
	g.Comment(struct_.gen.commentText(cursor))

	cName := cursor.Spelling()
	name := goName(cName)
	typ := cursor.Type().Spelling()

	if scalar, ok := scalars[clang.CursorKind(cursor.Type().CanonicalType().Kind())]; ok {
		g.Id(name).Add(scalar.goType)

	} else if unicode.IsUpper([]rune(typ)[0]) {
		g.Id(name).Id(goName(typ))

	} else if cursor.Type().Kind() == clang.Type_Pointer {
		g.Id(name).Qual("unsafe", "Pointer")

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
