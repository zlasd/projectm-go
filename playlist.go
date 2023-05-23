package projectm

/*
#cgo CFLAGS: -I${SRCDIR}/include
#cgo LDFLAGS: -lprojectM-4-playlist
#include <stdbool.h>
#include <projectM-4/playlist.h>
*/
import "C"
import "unsafe"

const (
	SORT_PREDICATE_FULL_PATH = iota
	SORT_PREDICATE_FILENAME_ONLY
)

const (
	SORT_ORDER_ASCENDING = iota
	SORT_ORDER_DESCENDING
)

type Playlist struct {
	handle     unsafe.Pointer
	Shuffle    bool
	RetryCount uint32
	Position   uint32
}

// CreatePlaylist creates a playlist manager for the given projectM instance
func (h *Handle) CreatePlaylist() *Playlist {
	return &Playlist{handle: unsafe.Pointer(C.projectm_playlist_create(C.projectm_handle(h.handle)))}
}

// Destroy destroys a previously created playlist manager.
func (p *Playlist) Destroy() {
	C.projectm_playlist_destroy(C.projectm_playlist_handle(p.handle))
}

// Connect connects the playlist manager to a projectM instance.
func (p *Playlist) Connect(handle *Handle) {}

// Size returns the number of presets in the current playlist.
func (p *Playlist) Size() uint32 {
	return 0
}

// Clear clears the playlist.
func (p *Playlist) Clear() {}

// Items returns a list of preset files inside the given range of the current playlist, in order.
func (p *Playlist) Items(start, count int32) []string {
	return nil
}

// Item returns the name of a preset at the given index in the current playlist.
func (p *Playlist) Item(index uint32) string {
	return "1"
}

// AddPath appends presets from the given path to the end of the current playlist.
func (p *Playlist) AddPath(path string, recurseSubDirs, allowDuplicates bool) uint32 {
	return 0
}

// InsertPath inserts presets from the given path to the end of the current playlist.
func (p *Playlist) InsertPath(path string, index uint32, recurseSubDirs, allowDuplicates bool) uint32 {
	return 0
}

// AddPreset adds a single preset to the end of the playlist.
func (p *Playlist) AddPreset(fileName string, allowDuplicates bool) bool {
	return false
}

// InsertPreset adds a single preset to the playlist at the specified position.
func (p *Playlist) InsertPreset(fileName string, index uint32, allowDuplicates bool) bool {
	return false
}

// AddPresets adds a single preset to the end of the playlist.
func (p *Playlist) AddPresets(fileNames []string, allowDuplicates bool) uint32 {
	return 0
}

// InsertPresets adds a single preset to the playlist at the specified position.
func (p *Playlist) InsertPresets(fileNames []string, index uint32, allowDuplicates bool) uint32 {
	return 0
}

// RemovePreset removes a single preset from the playlist at the specified position.
func (p *Playlist) RemovePreset(index uint32) bool {
	return false
}

// RemovePresets removes a number of presets from the playlist from the specified position.
func (p *Playlist) RemovePresets(index uint32) uint32 {
	return 0
}

// Sort sorts part or the whole playlist according to the given predicate and order.
func (p *Playlist) Sort(startIndex, count uint32, predicate, order int) {}

// PlayNext plays the next playlist item and returns the index of the new preset.
func (p *Playlist) PlayNext(hardCut bool) uint32 {
	return 0
}

// PlayPrevious plays the previous playlist item and returns the index of the new preset.
func (p *Playlist) PlayPrevious(hardCut bool) uint32 {
	return 0
}

// PlayLast plays the last preset played in the history and returns the index of the preset.
func (p *Playlist) PlayLast(hardCut bool) uint32 {
	return 0
}
