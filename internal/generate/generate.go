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
	// constantsDoclinks dochtml.Links
	// structs           structs
	// typedefs          typedefs
	// functions         functions
	// functionsDoclinks dochtml.Links
	// unusedFunctions   []string
}

func Generate() {
	gen := gen{
		headerFilename: "internal/interop/include/impeller.h",
		srcFilename:    "impeller.h",
		// constantsDoclinks: dochtml.ConstantsLinks(),
		// functionsDoclinks: dochtml.FunctionsLinks(),
	}

	timeFunction("TOTAL", func() {
		timeFunction("code generation", gen.generate)
		timeFunction("statistics", gen.statistics)

	})

	// fmt.Println()
	// usedFunctionCount := len(gen.functions) - len(gen.unusedFunctions)
	// fmt.Printf("sqlite3 API functions supported : %d/%d  %1.1f%%\n",
	// 	usedFunctionCount,
	// 	len(gen.functions),
	// 	(float64(usedFunctionCount)/float64(len(gen.functions)))*100,
	// )
	// fmt.Println()
}

func (gen *gen) generate() {
	gen.parseHeaderFile()
	gen.findEntities()
	// gen.enrich()
	gen.generateFiles()
	// gen.commentFunctions()
}

func (gen *gen) statistics() {
	// gen.findUnusedFunctions()
	// gen.writeUnusedFunctionsFile()
}

func (gen *gen) findEntities() {
	gen.tu.TranslationUnitCursor().Visit(func(cursor, _parent clang.Cursor) clang.ChildVisitResult {
		// ignore cursors that are not from the header file
		file, _, _, _ := cursor.Location().FileLocation()
		if file.Name() != gen.srcFilename {
			return clang.ChildVisit_Continue
		}

		switch cursor.Kind() {

		// 	case clang.Cursor_FunctionDecl:
		// 		gen.addFunction(cursor)

		case clang.Cursor_StructDecl:
			// fmt.Println("STRUCT =>", cursor.Spelling())
			// 		updated := false
			// 		for i, struct_ := range gen.structs {
			// 			if struct_.cName == cursor.Spelling() {
			// 				gen.updateStruct(i, cursor)
			// 				updated = true
			// 			}
			// 		}
			// 		if !updated {
			// 			gen.addStruct(cursor)
			// 		}

		case clang.Cursor_TypedefDecl:
			underlyingType := cursor.TypedefDeclUnderlyingType().Spelling()
			// fmt.Println("TYPEDEF =>", cursor.Spelling(), underlyingType)
			// fmt.Println("  ", cursor.TypedefDeclUnderlyingType().Kind().Spelling())

			// if cursor.TypedefDeclUnderlyingType().Kind() == clang.Type_Pointer {
			// 	fmt.Println("    ",
			// 		cursor.TypedefDeclUnderlyingType().PointeeType().Kind().Spelling(),
			// 		cursor.TypedefDeclUnderlyingType().PointeeType().Spelling(),
			// 	)
			// }

			if strings.Contains(underlyingType, "enum ") {
				name := strings.Split(underlyingType, " ")[1]
				name = strings.TrimPrefix(name, "Impeller")

				cursor.Visit(func(cursor, _parent clang.Cursor) (status clang.ChildVisitResult) {
					if cursor.Kind() == clang.Cursor_EnumDecl {
						gen.enums = append(gen.enums, enum{
							name:   name,
							cursor: cursor,
						})
					}
					return clang.ChildVisit_Continue
				})

				// fmt.Println(enumName)
				// fmt.Println("TYPEDEF =>", cursor.Type().Spelling(), cursor.TypedefDeclUnderlyingType().Spelling())
			}

			// 		if strings.Contains(underlyingType, "struct ") {
			// 			structName := strings.Split(underlyingType, " ")[1]
			// 			updated := false
			// 			for i, struct_ := range gen.structs {
			// 				if struct_.cName == structName {
			// 					if gen.structs[i].comment == "" {
			// 						gen.structs[i].comment = gen.commentText(cursor)
			// 					}
			// 					updated = true
			// 				}
			// 			}
			// 			if !updated {
			// 				gen.addStruct(cursor)
			// 			}

			// 		} else {
			// 			gen.addTypedef(cursor)
			// 		}

		}

		return clang.ChildVisit_Continue
	})
}

// func (gen *gen) enrich() {
// 	gen.functions.enrich()
// 	gen.structs.enrich()
// }

func (gen *gen) generateFiles() {
	gen.enums.generate()
	// gen.functions.generate()
	// gen.structs.generate()
	// gen.typedefs.generate()
	// gen.generateStructTest()
}

func timeFunction(title string, fn func()) {
	start := time.Now()
	fn()
	fmt.Printf("%-20s : %4.0fms\n",
		title,
		time.Since(start).Seconds()*1_000,
	)
}
