package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/zlasd/projectm-go"
	"os"
	"runtime"
)

var handle *projectm.Handle
var gCtx sdl.GLContext

func init() {
	runtime.LockOSThread()
}

func initGL() {
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 2)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 1)
	sdl.GLSetAttribute(sdl.GL_RED_SIZE, 5)
	sdl.GLSetAttribute(sdl.GL_GREEN_SIZE, 5)
	sdl.GLSetAttribute(sdl.GL_BLUE_SIZE, 5)
	sdl.GLSetAttribute(sdl.GL_DEPTH_SIZE, 16)
	sdl.GLSetAttribute(sdl.GL_DOUBLEBUFFER, 1)
}

func getPCMData(n int) ([]uint8, []uint8) {
	d1, d2 := make([]uint8, 0, n), make([]uint8, 0, n)
	for i := 0; i < n; i++ {
		d1 = append(d1, uint8(i%256))
		d2 = append(d2, uint8(255-i%256))
	}
	return d1, d2
}

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	initGL()
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN|sdl.WINDOW_OPENGL)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	gCtx, err = window.GLCreateContext()
	if err != nil {
		panic(err)
	}
	defer sdl.GLDeleteContext(gCtx)
	sdl.GLSetSwapInterval(1)

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)
	//
	//rect := sdl.Rect{0, 0, 200, 200}
	//colour := sdl.Color{R: 255, G: 0, B: 255, A: 255} // purple
	//pixel := sdl.MapRGBA(surface.Format, colour.R, colour.G, colour.B, colour.A)
	//surface.FillRect(&rect, pixel)
	//window.UpdateSurface()

	handle = projectm.Create()
	presets, _ := os.ReadFile("./resource/presets")
	handle.LoadPresetData(presets, true)
	d1, d2 := getPCMData(10000)
	handle.PCMAddUint8(d1, 0)
	handle.PCMAddUint8(d2, 1)
	running := true
	for running {
		handle.RenderFrame()
		window.GLSwap()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
	}
}
