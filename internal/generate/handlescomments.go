package generate

import (
	"path"
	"slices"
	"strings"

	"github.com/pekim/clang"
)

func (gen *gen) findHandledComments() {
	sourceRange := gen.tu.TranslationUnitCursor().Extent()
	tokens := gen.tu.Tokenize(sourceRange)

	for i, token := range tokens {
		if token.Kind() == clang.Token_Identifier &&
			gen.tu.TokenSpelling(token) == "IMPELLER_DEFINE_HANDLE" {
			openBracket := gen.tu.TokenSpelling(tokens[i+1])
			closeBracket := gen.tu.TokenSpelling(tokens[i+3])
			if openBracket != "(" || closeBracket != ")" {
				continue
			}
			handle := gen.tu.TokenSpelling(tokens[i+2])

			var commentLines []string
			for c := i - 1; c > 0; c-- {
				token := tokens[c]
				if token.Kind() == clang.Token_Comment {
					commentLines = append(commentLines, gen.tu.TokenSpelling(token))
				} else {
					break
				}
			}

			slices.Reverse(commentLines)
			commentLines = append(commentLines, "void* handle;")
			comment := strings.Join(commentLines, "\n")

			resourcesDir := clangResourceDir()
			parseArgs := []string{
				"-I", path.Join(resourcesDir, "include"),
				"-x", "c-header",
			}

			srcFilename := handle + "_comment.h"
			var srcFile clang.UnsavedFile
			srcFile.SetFilename(srcFilename)
			srcFile.SetContents(comment)
			index := clang.CreateIndex(0, 1)
			tu, errCode := index.ParseTranslationUnit2(
				srcFilename,
				parseArgs,
				[]clang.UnsavedFile{srcFile},
				uint32(clang.TranslationUnit_SkipFunctionBodies|clang.TranslationUnit_DetailedPreprocessingRecord),
			)
			if errCode != clang.Error_Success {
				fatal(errCode)
			}

			tu.TranslationUnitCursor().VisitChildren(clang.MakeChildVisitor(func(cursor, _parent clang.Cursor) clang.ChildVisitResult {
				file, _, _, _ := cursor.Location().FileLocation()
				if file.Name() != srcFilename {
					return clang.ChildVisit_Continue
				}

				gen.handlesComments[handle] = gen.newComment(cursor)
				return clang.ChildVisit_Continue
			}))
		}
	}
}
