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
	// constantsDoclinks dochtml.Links
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

		case clang.Cursor_EnumDecl:
			gen.enums = append(gen.enums, enum{
				gen:    gen,
				cursor: cursor,
				name:   goName(cursor.Spelling()),
			})

		case clang.Cursor_StructDecl:
			cName := cursor.Spelling()
			gen.structs = append(gen.structs, struct_{
				gen:    gen,
				cursor: cursor,
				cName:  cName,
				name:   goName(cName),
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
						struct_.pointer = true
					}

				}
			}
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
	gen.structs.generate()
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
