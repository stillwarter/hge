// For now we only build SDL, if we need to in the future we can use build tags
// such as: +build sdl
package hge

import (
	"runtime"

	"github.com/veandco/go-sdl2/sdl"
)

type Hwnd sdl.Window

func setTitle() {
	hwnd := (*sdl.Window)(stateHwnds[HWND])
	if hwnd != nil {
		hwnd.SetTitle(stateStrings[TITLE])
	}
}

func initNative(h *HGE) error {
	// Prevent crashes due to poor SDL & Go thread interactions
	runtime.LockOSThread()

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return err
	}

	// Create window
	bpp := 4
	if stateInts[SCREENBPP] >= 32 {
		bpp = 8
	}

	zbuffer := 0
	if stateBools[ZBUFFER] {
		zbuffer = 16
	}

	sdl.GL_SetAttribute(sdl.GL_RED_SIZE, bpp)
	sdl.GL_SetAttribute(sdl.GL_GREEN_SIZE, bpp)
	sdl.GL_SetAttribute(sdl.GL_BLUE_SIZE, bpp)
	sdl.GL_SetAttribute(sdl.GL_ALPHA_SIZE, bpp)
	sdl.GL_SetAttribute(sdl.GL_DEPTH_SIZE, zbuffer)
	sdl.GL_SetAttribute(sdl.GL_ACCELERATED_VISUAL, 1)
	sdl.GL_SetAttribute(sdl.GL_DOUBLEBUFFER, 1)

	var flags uint32
	flags |= uint32(sdl.WINDOW_OPENGL)
	if !stateBools[WINDOWED] {
		flags |= sdl.WINDOW_FULLSCREEN
	}

	title := stateStrings[TITLE]
	x, y := stateInts[SCREENX], stateInts[SCREENY]
	width, height := stateInts[SCREENWIDTH], stateInts[SCREENHEIGHT]

	window, err := sdl.CreateWindow(title, x, y, width, height, flags)
	if err != nil {
		sdl.Quit()
		return err
	}

	context, err := sdl.GL_CreateContext(window)
	if err != nil {
		sdl.Quit()
		return err
	}

	if err := sdl.GL_MakeCurrent(window, context); err != nil {
		sdl.Quit()
		return err
	}

	h.SetState((*Hwnd)(window))
	stateHwnds[HWND] = (*Hwnd)(window)

	if !stateBools[WINDOWED] {
		// 		bMouseOver = true;
		// 		if !pHGE->bActive {
		// 			pHGE->_FocusChange(true);
		// 		}
	}

	cursor := sdl.ENABLE
	if stateBools[HIDEMOUSE] {
		cursor = sdl.DISABLE
	}

	sdl.ShowCursor(cursor)

	return nil
}

func shutdownNative() {
	sdl.Quit()
}

func initPowerStatus() error {
	return nil
}
