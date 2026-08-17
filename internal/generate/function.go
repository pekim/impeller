package generate

import (
	"github.com/go-clang/clang-v15/clang"
)

type function struct {
	gen    *gen
	cursor clang.Cursor
	name   string
}

func (fn function) generate(file file) {
	file.Comment(fn.name)
}

type functions []function

func (functions functions) generate() {
	file := newFile("function.go", "impeller")
	defer file.save()

	for _, function := range functions {
		function.generate(file)
		file.Line()
	}
}
