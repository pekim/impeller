package generate

import (
	"os"
	"os/exec"
	"path"
	"strings"
	"sync"

	"github.com/pekim/clang"
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

	headerData, err := os.ReadFile(gen.headerFilename)
	fatalOnError(err)
	headerContents := string(headerData)

	// fix incorrect doc comment
	headerContents = strings.Replace(headerContents,
		"@param[out]  code_unit_index The range.",
		"@param[out]  out_range The range.",
		1)

	// write a copy of the modified header file, just for reference
	modifiedHeaderFilename := strings.Replace(gen.headerFilename, "impeller.h", "impeller_modified.h", 1)
	err = os.WriteFile(modifiedHeaderFilename, []byte(headerContents), 0600)
	fatalOnError(err)

	var srcFile clang.UnsavedFile
	srcFile.SetFilename(gen.srcFilename)
	srcFile.SetContents(headerContents)
	srcFile.Length = uint64(len(headerContents))

	index := clang.CreateIndex(0, 1)
	tu, errCode := index.ParseTranslationUnit2(
		gen.srcFilename,
		parseArgs,
		[]clang.UnsavedFile{srcFile},
		uint32(clang.TranslationUnit_SkipFunctionBodies|clang.TranslationUnit_DetailedPreprocessingRecord),
	)
	if errCode != clang.Error_Success {
		fatal(errCode)
	}
	gen.tu = tu

	gen.lines = strings.Split(headerContents, "\n")
}
