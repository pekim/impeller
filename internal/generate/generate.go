package generate

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-clang/clang-v15/clang"
)

type gen struct {
	headerFilename string
	srcFilename    string
	tu             clang.TranslationUnit
	lines          []string
	enums          enums
	structs        structs
	functions      functions
}

func Generate() {
	gen := gen{
		headerFilename: "internal/interop/include/impeller.h",
		srcFilename:    "impeller.h",
	}

	timeFunction("code generation", gen.generate)
	gen.printStatistics()
	fmt.Println()
}

func (gen *gen) generate() {
	gen.parseHeaderFile()
	gen.findEntities()
	gen.generateFiles()
}

func (gen *gen) printStatistics() {
	functionCount := len(gen.functions)
	supportedFunctionCount := 0
	for _, fn := range gen.functions {
		if supported, _ := fn.supported(); supported {
			supportedFunctionCount++
		}
	}
	fmt.Printf("functions supported  : %d/%d  %1.1f%%\n",
		supportedFunctionCount,
		functionCount,
		(float64(supportedFunctionCount)/float64(functionCount))*100,
	)
}

func (gen *gen) findEntities() {
	gen.tu.TranslationUnitCursor().Visit(func(cursor, _parent clang.Cursor) clang.ChildVisitResult {
		// ignore cursors that are not from the header file
		file, _, _, _ := cursor.Location().FileLocation()
		if file.Name() != gen.srcFilename {
			return clang.ChildVisit_Continue
		}

		switch cursor.Kind() {

		case clang.Cursor_FunctionDecl:
			gen.functions = append(gen.functions, newFunction(gen, cursor))

		case clang.Cursor_EnumDecl:
			gen.enums = append(gen.enums, enum{
				gen:    gen,
				cursor: cursor,
				name:   goName(cursor.Spelling()),
			})

		case clang.Cursor_StructDecl:
			cName := cursor.Spelling()
			gen.structs = append(gen.structs, struct_{
				gen:     gen,
				cursor:  cursor,
				cName:   cName,
				name:    goName(cName),
				comment: gen.commentText(cursor),
			})

		case clang.Cursor_TypedefDecl:
			underlyingType := cursor.TypedefDeclUnderlyingType().Spelling()
			parts := strings.Split(underlyingType, " ")

			switch cursor.TypedefDeclUnderlyingType().Kind() {
			case clang.Type_Pointer:
				if parts[0] == "struct" {
					structName := parts[1]
					struct_, found := gen.structs.find(structName)
					if found {
						struct_.handle = true
						if struct_.comment == "" {
							struct_.comment = gen.commentText(cursor)
						}
					}
				}
			}
		}

		return clang.ChildVisit_Continue
	})
}

func (gen *gen) generateFiles() {
	gen.enums.generate()
	gen.functions.generate()
	gen.structs.generate()
	gen.generateStructTest()
}

func timeFunction(title string, fn func()) {
	start := time.Now()
	fn()
	fmt.Printf("%-20s : %.0fms\n",
		title,
		time.Since(start).Seconds()*1_000,
	)
}
