package generate

import (
	"os"
	"os/exec"
	"path"
	"strings"
	"sync"

	"github.com/go-clang/clang-v15/clang"
)

var clangResourceDir = sync.OnceValue(func() string {
	out, err := exec.Command("clang", "-print-resource-dir").Output()
	fatalOnError(err)

	resDir := strings.TrimSpace(string(out))
	parts := strings.Split(resDir, "\n")
	resDir = parts[0]

	if resDir == "" {
		fatal("no output when getting clang resource dir")
	}
	if !strings.HasPrefix(resDir, "/") {
		fatalf("expected clang resource dir to start with '/', but it %s", resDir)
	}

	return resDir
})

func (gen *gen) parseHeaderFile() {
	resourcesDir := clangResourceDir()
	parseArgs := []string{
		"-I", path.Join(resourcesDir, "include"),
		"-x", "c-header",
	}

	// Read header file, and turn comment blocks ("/* ... */") in to documentation
	// comment blocks ("/** ... */").
	// This enables clang to make the comment blocks available in the AST.
	headerData, err := os.ReadFile(gen.headerFilename)
	fatalOnError(err)
	headerContents := string(headerData)

	// fix incorrect doc comment
	headerContents = strings.Replace(headerContents,
		"@param[out]  code_unit_index The range.",
		"@param[out]  out_range The range.",
		1)

	headerLines := strings.Split(string(headerContents), "\n")
	for i, line := range headerLines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "/********") {
			if strings.HasSuffix(line, "*/") {
				headerLines[i] = ""
			} else {
				headerLines[i] = "/**"
			}
		} else if strings.HasPrefix(line, "/*") {
			headerLines[i] = strings.Replace(line, "/*", "/**", 1)
		}
	}
	headerContents = strings.Join(headerLines, "\n")

	// write a copy of the modified header file, just for reference
	modifiedHeaderFilename := strings.Replace(gen.headerFilename, "impeller.h", "impeller_modified.h", 1)
	err = os.WriteFile(modifiedHeaderFilename, []byte(headerContents), 0600)
	fatalOnError(err)

	srcFile := clang.NewUnsavedFile(gen.srcFilename, string(headerContents))

	index := clang.NewIndex(0, 1)
	errCode := index.ParseTranslationUnit2(
		gen.srcFilename,
		parseArgs,
		[]clang.UnsavedFile{srcFile},
		uint32(clang.TranslationUnit_SkipFunctionBodies|clang.TranslationUnit_DetailedPreprocessingRecord),
		&gen.tu,
	)
	if errCode != clang.Error_Success {
		fatal(errCode)
	}

	gen.lines = headerLines
}
