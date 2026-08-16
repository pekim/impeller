package generate

import (
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/go-clang/clang-v15/clang"
)

type enum struct {
	gen    *gen
	name   string
	cursor clang.Cursor
}

func (enum enum) generate(file file) {
	file.Type().Id(enum.name).Id("enum")

	enum.cursor.Visit(func(cursor, _parent clang.Cursor) (status clang.ChildVisitResult) {
		name := cursor.Spelling()
		name = strings.TrimPrefix(name, "kImpeller")
		value := int(cursor.EnumConstantDeclValue())

		file.
			Const().
			Id(name).Id(enum.name).
			Op("=").
			Lit(value).
			Do(func(s *jen.Statement) {
				comment := enum.gen.commentText(cursor)
				if comment != "" {
					s.Comment(comment)
				}
			})

		return clang.ChildVisit_Continue
	})
}

type enums []enum

func (enums enums) generate() {
	file := newFile("enum.go", "impeller")
	defer file.save()

	for _, enum := range enums {
		enum.generate(file)
		file.Line()
	}
}
