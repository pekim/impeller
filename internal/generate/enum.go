package generate

import (
	"github.com/dave/jennifer/jen"
	"github.com/go-clang/clang-v15/clang"
)

type enum struct {
	gen    *gen
	cursor clang.Cursor
	name   string
}

func (enum enum) generate(file file) {
	file.Type().Id(enum.name).Id("enum")

	enum.cursor.Visit(func(cursor, _parent clang.Cursor) (status clang.ChildVisitResult) {
		name := goName(cursor.Spelling())
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
