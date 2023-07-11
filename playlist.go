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
	playListHandle unsafe.Pointer
	handle         unsafe.Pointer
	Shuffle        bool
	RetryCount     uint32
	Position       uint32
}

// CreatePlaylist creates a playlist manager for the given projectM instance
func (h *Handle) CreatePlaylist() *Playlist {
	return &Playlist{
		playListHandle: unsafe.Pointer(C.projectm_playlist_create(C.projectm_handle(h.handle))),
		handle:         unsafe.Pointer(h),
	}
}

// Destroy destroys a previously created playlist manager.
func (p *Playlist) Destroy() {
	C.projectm_playlist_destroy(C.projectm_playlist_handle(p.playListHandle))
}

// Connect connects the playlist manager to a projectM instance.
func (p *Playlist) Connect(handle *Handle) {
	C.projectm_playlist_connect(C.projectm_playlist_handle(p.playListHandle), C.projectm_handle(p.handle))
}

// Size returns the number of presets in the current playlist.
func (p *Playlist) Size() uint32 {
	return uint32(C.projectm_playlist_size(C.projectm_playlist_handle(p.playListHandle)))
}

// Clear clears the playlist.
func (p *Playlist) Clear() {
	C.projectm_playlist_clear(C.projectm_playlist_handle(p.playListHandle))
}

// Items returns a list of preset files inside the given range of the current playlist, in order.
func (p *Playlist) Items(start, count uint32) []string {
	var cStrList **C.char = (**C.char)(C.projectm_playlist_items(
		C.projectm_playlist_handle(p.playListHandle),
		C.uint32_t(start), C.uint32_t(count),
	))
	slice := unsafe.Slice(cStrList, count)
	items := make([]string, 0)
	for i := 0; i < int(count); i++ {
		cStr := (*C.char)(slice[i])
		if unsafe.Pointer(cStr) == nil {
			break
		}
		items = append(items, C.GoString(cStr))
	}
	return items
}

// Item returns the name of a preset at the given index in the current playlist.
func (p *Playlist) Item(index uint32) string {
	return p.Items(index, 1)[0]
}

// AddPath appends presets from the given path to the end of the current playlist.
func (p *Playlist) AddPath(path string, recurseSubDirs, allowDuplicates bool) uint32 {
	return uint32(C.projectm_playlist_add_path(
		C.projectm_playlist_handle(p.playListHandle),
		C.CString(path),
		C.bool(recurseSubDirs), C.bool(allowDuplicates),
	))
}

// InsertPath inserts presets from the given path to the end of the current playlist.
func (p *Playlist) InsertPath(path string, index uint32, recurseSubDirs, allowDuplicates bool) uint32 {
	return uint32(C.projectm_playlist_insert_path(
		C.projectm_playlist_handle(p.playListHandle),
		C.CString(path), C.uint32_t(index),
		C.bool(recurseSubDirs), C.bool(allowDuplicates),
	))
}

// AddPreset adds a single preset to the end of the playlist.
func (p *Playlist) AddPreset(fileName string, allowDuplicates bool) bool {
	return bool(C.projectm_playlist_add_preset(
		C.projectm_playlist_handle(p.playListHandle),
		C.CString(fileName), C.bool(allowDuplicates),
	))
}

// InsertPreset adds a single preset to the playlist at the specified position.
func (p *Playlist) InsertPreset(fileName string, index uint32, allowDuplicates bool) bool {
	return bool(C.projectm_playlist_insert_preset(
		C.projectm_playlist_handle(p.playListHandle),
		C.CString(fileName), C.uint32_t(index), C.bool(allowDuplicates),
	))
}

// AddPresets adds a single preset to the end of the playlist.
func (p *Playlist) AddPresets(fileNames []string, allowDuplicates bool) uint32 {
	var cnt uint32 = 0
	for _, fn := range fileNames {
		ok := p.AddPreset(fn, allowDuplicates)
		if ok {
			cnt++
		}
	}
	return cnt
}

// InsertPresets adds a single preset to the playlist at the specified position.
func (p *Playlist) InsertPresets(fileNames []string, index uint32, allowDuplicates bool) uint32 {
	var cnt uint32 = 0
	for i := len(fileNames) - 1; i >= 0; i-- {
		ok := p.InsertPreset(fileNames[0], index, allowDuplicates)
		if ok {
			cnt++
		}
	}
	return cnt
}

// RemovePreset removes a single preset from the playlist at the specified position.
func (p *Playlist) RemovePreset(index uint32) bool {
	return bool(C.projectm_playlist_remove_preset(
		C.projectm_playlist_handle(p.playListHandle),
		C.uint32_t(index),
	))
}

// RemovePresets removes a number of presets from the playlist from the specified position.
func (p *Playlist) RemovePresets(index, count uint32) uint32 {
	return uint32(C.projectm_playlist_remove_presets(
		C.projectm_playlist_handle(p.playListHandle),
		C.uint32_t(index), C.uint32_t(count),
	))
}

// Sort sorts part or the whole playlist according to the given predicate and order.
func (p *Playlist) Sort(startIndex, count uint32, predicate, order int) {
	C.projectm_playlist_sort(
		C.projectm_playlist_handle(p.playListHandle),
		C.uint32_t(startIndex), C.uint32_t(count),
		C.projectm_playlist_sort_predicate(predicate),
		C.projectm_playlist_sort_order(order),
	)
}

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
