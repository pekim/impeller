package generate

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/dave/jennifer/jen"
)

type file struct {
	*jen.File
	filename string
}

func newFile(filename string, packageName string) file {
	f := file{
		File:     jen.NewFile(packageName),
		filename: filename,
	}

	f.HeaderComment("This is a generated file. DO NOT EDIT.")
	f.Line()

	return f
}

func (f file) save() {
	filename := filepath.Join(".", f.filename)
	f.NoFormat = true
	err := f.Save(filename)
	if err != nil {
		fatalOnError(err)
	}

	cmd := exec.Command("goimports", "-w", filename)
	fmtOutput, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("failed to format source, writing unformatted source to %s\n", filename)
		f.NoFormat = true
		err := f.Save(filename)
		fatalOnError(err)
	}
	if len(fmtOutput) > 0 {
		fmt.Println(string(fmtOutput))
	}
}
