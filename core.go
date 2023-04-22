package projectm

/*
#include <projectM-4/projectM.h>
*/
import "C"
import "unsafe"

type Handle struct {
	handle unsafe.Pointer
}

/*
Create

	Creates a new projectM instance.
	If this function returns nil, in most cases the OpenGL context is not initialized, not made
	current or insufficient to render projectM visuals.

	@return A projectM handle for the newly created instance.
	nil if the instance could not be created successfully.
*/
func Create() *Handle {
	// TODO call projectm_create
	return &Handle{}
}

// Destroy destroys the given instance and frees the resources.
func (h *Handle) Destroy() {}

// LoadPresetFile loads a preset from the given filename/URL.
func (h *Handle) LoadPresetFile(fileName string, smoothTransition bool) {}

// LoadPresetData loads a preset from binary.
func (h *Handle) LoadPresetData(data []byte, smoothTransition bool) {}

// ResetTextures reloads all textures.
func (h *Handle) ResetTextures() {}

// GetVersionComponents returns the runtime library version components as individual integers.
func (h *Handle) GetVersionComponents() (int, int, int) {
	return 4, 0, 0
}

// GetVersionString returns the runtime library version as a string.
func (h *Handle) GetVersionString() string {
	return "4.0.0"
}

// GetVCSVersionString returns the VCS revision from which the projectM library was built.
func (h *Handle) GetVCSVersionString() string {
	return "422af469731559c32c93e9513cebf2fe6c2fec78"
}
