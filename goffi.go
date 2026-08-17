// This is a generated file. DO NOT EDIT.

package impeller

import (
	"unsafe"

	ffi "github.com/go-webgpu/goffi/ffi"
	types "github.com/go-webgpu/goffi/types"
)

var _ImpellerGetVersion unsafe.Pointer
var _ImpellerGetVersion_cif = &types.CallInterface{}

var initialised = false

func Init() error {
	if initialised {
		return nil
	}
	defer func() {
		initialised = true
	}()

	filepath, err := libraryFilepath()
	if err != nil {
		return err
	}

	handle, err := ffi.LoadLibrary(filepath)
	if err != nil {
		return err
	}

	_ImpellerGetVersion, err = ffi.GetSymbol(handle, "ImpellerGetVersion")
	if err != nil {
		return err
	}
	err = ffi.PrepareCallInterface(
		_ImpellerGetVersion_cif,
		types.DefaultCall,
		types.UInt32TypeDescriptor,
		[]*types.TypeDescriptor{})
	if err != nil {
		return err
	}

	return nil
}
