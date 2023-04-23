package projectm

const (
	PROJECTM_MONO   = 1
	PROJECTM_STEREO = 2
)

const (
	PROJECTM_CHANNEL_L = 0
	PROJECTM_CHANNEL_0 = 0
	PROJECTM_CHANNEL_R = 1
	PROJECTM_CHANNEL_1 = 1
)

// PCMGetMaxChannels returns the maximum number of audio samples that can be stored.
func PCMGetMaxChannels() uint32 {
	return 0
}

// PCMAddFloat adds 32-bit floating-point audio samples.
func (h *Handle) PCMAddFloat(samples []float32, channels int) {}

// PCMAddInt16 adds 16-bit integer audio samples.
func (h *Handle) PCMAddInt16(samples []int16, channels int) {}

// PCMAddUint8 adds 8-bit unsigned integer audio samples.
func (h *Handle) PCMAddUint8(samples []uint8, channels int) {}
