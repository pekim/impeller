package impeller

import (
	"unsafe"

	"github.com/go-webgpu/goffi/ffi"
)

// Callback is
// a callback invoked by Impeller that passes a user supplied baton back to the
// user. Impeller does not interpret the baton in any way. The way the baton is
// specified and the thread on which the callback is invoked depends on how the
// user supplies the callback to Impeller.
type Callback uintptr

// ProcAddressCallback is
// a callback used by Impeller to allow the user to resolve function pointers.
// A user supplied baton that is uninterpreted by Impeller is passed back to
// the user in the callback. How the baton is specified to Impeller and the
// thread on which the callback is invoked depends on how the callback is
// specified to Impeller.
type ProcAddressCallback uintptr

// VulkanProcAddressCallback is
// a callback used by Impeller to allow the user to resolve Vulkan function
// pointers. A user supplied baton that is uninterpreted by Impeller is passed
// back to the user in the callback.
type VulkanProcAddressCallback uintptr

func NewCallback(cb func(user_data unsafe.Pointer)) Callback {
	return Callback(ffi.NewCallback(cb))
}

func NewProcAddressCallback(cb func(
	procName string,
	user_data unsafe.Pointer,
)) ProcAddressCallback {
	return ProcAddressCallback(ffi.NewCallback(func(
		proc_name *byte,
		user_data unsafe.Pointer,
	) {
		cb(
			goString(proc_name),
			user_data,
		)
	}))
}

func NewVulkanProcAddressCallback(cb func(
	vulkan_instance unsafe.Pointer,
	vulkan_proc_name string,
	user_data unsafe.Pointer,
)) VulkanProcAddressCallback {
	return VulkanProcAddressCallback(ffi.NewCallback(func(
		vulkan_instance unsafe.Pointer,
		vulkan_proc_name *byte,
		user_data unsafe.Pointer,
	) {
		cb(
			vulkan_instance,
			goString(vulkan_proc_name),
			user_data,
		)
	}))
}
