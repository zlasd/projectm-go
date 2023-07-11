package projectm

/*
#cgo CFLAGS: -I${SRCDIR}/include
#cgo LDFLAGS: -lprojectM-4
#include <projectM-4/projectM.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

type Handle struct {
	handle             unsafe.Pointer
	TextureSearchPaths []string // the texture search paths
	BeatSensitivity    float32  // the beat sensitivity
	HardCutDuration    float64  // the minimum display time before a hard cut can happen
	HardCutEnabled     bool     // enables or disables hard cuts
	HardCutSensitivity float32  // the hard cut volume sensitivity
	SoftCutDuration    float64  // the time in seconds for a soft transition between two presets
	PresetDuration     float64  // the preset display duration before switching to the next using a soft cut
	MeshSize           []uint32 // the per-pixel equation mesh size in units
	FPS                int32    // the current/average frames per second
	AspectCorrection   bool     // enabled or disables aspect ratio correction in presets that support it
	EasterEgg          float32  // the "easter egg" value
	PresetLocked       bool     // locks or unlocks the current preset
	WindowSize         []uint32 // the current viewport size in pixels

	// temp C pointer
	tempPointer []unsafe.Pointer
}

/*
Create

	Creates a new projectM instance.
	If this function returns nil, in most cases the OpenGL context is not initialized, not made
	current or insufficient to render projectM visuals.

	@return A projectM playListHandle for the newly created instance.
	nil if the instance could not be created successfully.
*/
func Create() *Handle {
	return &Handle{handle: unsafe.Pointer(C.projectm_create())}
}

// Destroy destroys the given instance and frees the resources.
func (h *Handle) Destroy() {
	C.projectm_destroy(C.projectm_handle(h.handle))
	for _, p := range h.tempPointer {
		C.free(p)
	}
}

// LoadPresetFile loads a preset from the given filename/URL.
func (h *Handle) LoadPresetFile(fileName string, smoothTransition bool) {
	cFileName := C.CString(fileName)
	C.projectm_load_preset_file(C.projectm_handle(h.handle), cFileName, C.bool(smoothTransition))
	h.tempPointer = append(h.tempPointer, unsafe.Pointer(cFileName))
}

// LoadPresetData loads a preset from binary.
func (h *Handle) LoadPresetData(data []byte, smoothTransition bool) {
	cData := C.CBytes(data)
	C.projectm_load_preset_data(C.projectm_handle(h.handle), (*C.char)(cData), C.bool(smoothTransition))
	h.tempPointer = append(h.tempPointer, cData)
}

// ResetTextures reloads all textures.
func (h *Handle) ResetTextures() {
	C.projectm_reset_textures(C.projectm_handle(h.handle))
}

// RenderFrame renders a single frame.
func (h *Handle) RenderFrame() {}

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
