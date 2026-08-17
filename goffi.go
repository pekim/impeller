// This is a generated file. DO NOT EDIT.

package impeller

import (
	"unsafe"

	ffi "github.com/go-webgpu/goffi/ffi"
	types "github.com/go-webgpu/goffi/types"
)

var funcImpellerGetVersion unsafe.Pointer
var cifImpellerGetVersion = &types.CallInterface{}

var funcImpellerPathBuilderNew unsafe.Pointer
var cifImpellerPathBuilderNew = &types.CallInterface{}

var funcImpellerPaintNew unsafe.Pointer
var cifImpellerPaintNew = &types.CallInterface{}

var funcImpellerTypographyContextNew unsafe.Pointer
var cifImpellerTypographyContextNew = &types.CallInterface{}

var funcImpellerParagraphStyleNew unsafe.Pointer
var cifImpellerParagraphStyleNew = &types.CallInterface{}

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

	funcImpellerGetVersion, err = ffi.GetSymbol(handle, "ImpellerGetVersion")
	if err != nil {
		return err
	}
	err = ffi.PrepareCallInterface(
		cifImpellerGetVersion,
		types.DefaultCall,
		types.UInt32TypeDescriptor,
		[]*types.TypeDescriptor{})
	if err != nil {
		return err
	}

	funcImpellerPathBuilderNew, err = ffi.GetSymbol(handle, "ImpellerPathBuilderNew")
	if err != nil {
		return err
	}
	err = ffi.PrepareCallInterface(
		cifImpellerPathBuilderNew,
		types.DefaultCall,
		types.PointerTypeDescriptor,
		[]*types.TypeDescriptor{})
	if err != nil {
		return err
	}

	funcImpellerPaintNew, err = ffi.GetSymbol(handle, "ImpellerPaintNew")
	if err != nil {
		return err
	}
	err = ffi.PrepareCallInterface(
		cifImpellerPaintNew,
		types.DefaultCall,
		types.PointerTypeDescriptor,
		[]*types.TypeDescriptor{})
	if err != nil {
		return err
	}

	funcImpellerTypographyContextNew, err = ffi.GetSymbol(handle, "ImpellerTypographyContextNew")
	if err != nil {
		return err
	}
	err = ffi.PrepareCallInterface(
		cifImpellerTypographyContextNew,
		types.DefaultCall,
		types.PointerTypeDescriptor,
		[]*types.TypeDescriptor{})
	if err != nil {
		return err
	}

	funcImpellerParagraphStyleNew, err = ffi.GetSymbol(handle, "ImpellerParagraphStyleNew")
	if err != nil {
		return err
	}
	err = ffi.PrepareCallInterface(
		cifImpellerParagraphStyleNew,
		types.DefaultCall,
		types.PointerTypeDescriptor,
		[]*types.TypeDescriptor{})
	if err != nil {
		return err
	}

	return nil
}
