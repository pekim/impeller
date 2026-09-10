package impeller

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

//go:embed internal/interop/lib/libimpeller.so
var sharedObject []byte

func libraryFilepath() (string, error) {
	libraryHash := "TODO"
	filepath := filepath.Join(os.TempDir(), fmt.Sprintf("impeller-%s.so", libraryHash))

	// Check if the file exists. If it does, don't create it.
	//
	// The code would be simpler, and very nearly as quick, if the check were omitted and the
	// file written every time. However that causes problems if multiple applications are run
	// concurrently. The replaced file becomes unavailable, and a segment violation results.
	_, err := os.Stat(filepath)
	if err == nil {
		return filepath, nil
	}
	if !errors.Is(err, os.ErrNotExist) {
		return "", fmt.Errorf("failed to check existence of shared object file %q : %w", filepath, err)
	}

	err = os.WriteFile(filepath, sharedObject, 0755)
	if err != nil {
		return "", fmt.Errorf("failed to write shared object file %q : %w", filepath, err)
	}

	return filepath, nil
}
