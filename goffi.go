// This is a generated file. DO NOT EDIT.

package impeller

import (
	"unsafe"

	ffi "github.com/go-webgpu/goffi/ffi"
	types "github.com/go-webgpu/goffi/types"
)

var _ImpellerGetVersion unsafe.Pointer
var _ImpellerGetVersion_cif = &types.CallInterface{}

var _ImpellerPathBuilderNew unsafe.Pointer
var _ImpellerPathBuilderNew_cif = &types.CallInterface{}

var _ImpellerPaintNew unsafe.Pointer
var _ImpellerPaintNew_cif = &types.CallInterface{}

var _ImpellerTypographyContextNew unsafe.Pointer
var _ImpellerTypographyContextNew_cif = &types.CallInterface{}

var _ImpellerParagraphStyleNew unsafe.Pointer
var _ImpellerParagraphStyleNew_cif = &types.CallInterface{}

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

	_ImpellerPathBuilderNew, err = ffi.GetSymbol(handle, "ImpellerPathBuilderNew")
	if err != nil {
		return err
	}
	err = ffi.PrepareCallInterface(
		_ImpellerPathBuilderNew_cif,
		types.DefaultCall,
		types.PointerTypeDescriptor,
		[]*types.TypeDescriptor{})
	if err != nil {
		return err
	}

	_ImpellerPaintNew, err = ffi.GetSymbol(handle, "ImpellerPaintNew")
	if err != nil {
		return err
	}
	err = ffi.PrepareCallInterface(
		_ImpellerPaintNew_cif,
		types.DefaultCall,
		types.PointerTypeDescriptor,
		[]*types.TypeDescriptor{})
	if err != nil {
		return err
	}

	_ImpellerTypographyContextNew, err = ffi.GetSymbol(handle, "ImpellerTypographyContextNew")
	if err != nil {
		return err
	}
	err = ffi.PrepareCallInterface(
		_ImpellerTypographyContextNew_cif,
		types.DefaultCall,
		types.PointerTypeDescriptor,
		[]*types.TypeDescriptor{})
	if err != nil {
		return err
	}

	_ImpellerParagraphStyleNew, err = ffi.GetSymbol(handle, "ImpellerParagraphStyleNew")
	if err != nil {
		return err
	}
	err = ffi.PrepareCallInterface(
		_ImpellerParagraphStyleNew_cif,
		types.DefaultCall,
		types.PointerTypeDescriptor,
		[]*types.TypeDescriptor{})
	if err != nil {
		return err
	}

	return nil
}
