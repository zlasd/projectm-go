package projectm

/*
#cgo CFLAGS: -I${SRCDIR}/include
#cgo LDFLAGS: -lprojectM-4
#include <projectM-4/projectM.h>
*/
import "C"

// PCMGetMaxChannels returns the maximum number of audio samples that can be stored.
func PCMGetMaxChannels() uint32 {
	return uint32(C.projectm_pcm_get_max_samples())
}

// PCMAddFloat adds 32-bit floating-point audio samples.
func (h *Handle) PCMAddFloat(samples []float32, channels int) {
	if len(samples) == 0 {
		return
	}
	C.projectm_pcm_add_float(C.projectm_handle(h.handle), (*C.float)(&samples[0]), C.uint(len(samples)), C.projectm_channels(channels))
}

// PCMAddInt16 adds 16-bit integer audio samples.
func (h *Handle) PCMAddInt16(samples []int16, channels int) {
	if len(samples) == 0 {
		return
	}
	C.projectm_pcm_add_int16(C.projectm_handle(h.handle), (*C.int16_t)(&samples[0]), C.uint(len(samples)), C.projectm_channels(channels))
}

// PCMAddUint8 adds 8-bit unsigned integer audio samples.
func (h *Handle) PCMAddUint8(samples []uint8, channels int) {
	if len(samples) == 0 {
		return
	}
	C.projectm_pcm_add_uint8(C.projectm_handle(h.handle), (*C.uint8_t)(&samples[0]), C.uint(len(samples)), C.projectm_channels(channels))
}
