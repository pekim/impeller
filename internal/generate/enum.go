package generate

import (
	"github.com/dave/jennifer/jen"
	"github.com/go-clang/clang-v15/clang"
)

type enum struct {
	gen     *gen
	cursor  clang.Cursor
	cName   string
	name    string
	comment comment
}

func (enum enum) generate(file file) {
	enum.comment.text()
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
				comment := enum.gen.newComment(cursor).text()
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

func (enums enums) find(cName string) (*enum, bool) {
	for i, enum := range enums {
		if enum.cName == cName {
			return &(enums[i]), true
		}
	}

	return nil, false
}

func (enums enums) findByGoName(name string) (*enum, bool) {
	for i, enum_ := range enums {
		if enum_.name == name {
			return &(enums[i]), true
		}
	}

	return nil, false
}

func (enums enums) resolve() {
	for i := range enums {
		(&enums[i]).comment.resolve()
	}
}
