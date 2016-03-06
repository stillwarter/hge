package hge

import "github.com/veandco/go-sdl2/sdl"

// HGE System state constants
const (
	WINDOWED      BoolState = iota // bool run in window? (default: true)
	ZBUFFER       BoolState = iota // bool use z-buffer? (default: false)
	TEXTUREFILTER BoolState = iota // bool texture filtering? (default: true)

	USESOUND BoolState = iota // bool use sound? (default: true)

	DONTSUSPEND BoolState = iota // bool focus lost:suspend? (default: false)
	HIDEMOUSE   BoolState = iota // bool hide system cursor? (default: true)

	SHOWSPLASH BoolState = iota // bool show splash? (default: true)

	boolstate BoolState = iota
)

// When any of these return true, it indicates to stop the main loop from
// continuing to run.
const (
	FRAMEFUNC      FuncState = iota // func() bool frame function (default: nil) (you MUST set this)
	RENDERFUNC     FuncState = iota // func() bool render function (default: nil)
	FOCUSLOSTFUNC  FuncState = iota // func() bool focus lost function (default: nil)
	FOCUSGAINFUNC  FuncState = iota // func() bool focus gain function (default: nil)
	GFXRESTOREFUNC FuncState = iota // func() bool gfx restore function (default: nil)
	EXITFUNC       FuncState = iota // func() bool exit function (default: nil)

	funcstate FuncState = iota
)

const (
	HWND       HwndState = iota // int		window handle: read only
	HWNDPARENT HwndState = iota // int		parent win handle	(default: 0)

	hwndstate HwndState = iota
)

const (
	SCREENWIDTH      IntState = iota // int screen width (default: 800)
	SCREENHEIGHT     IntState = iota // int screen height (default: 600)
	SCREENX          IntState = iota // int screen x location (default: centered)
	SCREENY          IntState = iota // int screen y location (default: centered)
	SCREENBPP        IntState = iota // int screen bitdepth (default: 32) (desktop bpp in windowed mode)
	ORIGSCREENWIDTH  IntState = iota // int original screen width (default: 800 ... not valid until hge.System_Initiate()!)
	ORIGSCREENHEIGHT IntState = iota // int original screen height (default: 600 ... not valid until hge.System_Initiate()!))
	FPS              IntState = iota // int fixed fps (default: hge.FPS_UNLIMITED)
	MINDELTATIME     IntState = iota // int minimum delta time in miliseconds between frames (default: 1000)

	SAMPLERATE   IntState = iota // int sample rate (default: 44100)
	FXVOLUME     IntState = iota // int global fx volume (default: 100)
	MUSVOLUME    IntState = iota // int global music volume (default: 100)
	STREAMVOLUME IntState = iota // int stream music volume (default: 100)

	POWERSTATUS IntState = iota // int battery life percent + status

	intstate IntState = iota
)

const (
	ICON  StringState = iota // string icon resource (default: nil)
	TITLE StringState = iota // string window title (default: "HGE")

	INIFILE StringState = iota // string ini file (default: nil) (meaning no file)
	LOGFILE StringState = iota // string log file (default: nil) (meaning no file)

	stringstate StringState = iota
)

type (
	BoolState   int
	FuncState   int
	HwndState   int
	IntState    int
	StringState int
)

type StateFunc func() bool

var (
	stateBools   = new([boolstate]bool)
	stateFuncs   = new([funcstate]StateFunc)
	stateHwnds   = new([hwndstate]*Hwnd)
	stateInts    = new([intstate]int)
	stateStrings = new([stringstate]string)

	setupBools   = new([boolstate]func(*HGE) error)
	setupHwnds   = new([hwndstate]func(*HGE) error)
	setupInts    = new([intstate]func(*HGE) error)
	setupStrings = new([stringstate]func(*HGE) error)
)

func init() {
	// Bool states
	setupBools[WINDOWED] = setupWindowed
	setupBools[ZBUFFER] = setupZBuffer
	setupBools[TEXTUREFILTER] = setupTextureFilter
	setupBools[USESOUND] = setupUseSound
	setupBools[DONTSUSPEND] = setupDonetSuspend
	setupBools[HIDEMOUSE] = setupHideMouse
	setupBools[SHOWSPLASH] = setupShowSplash

	// Func states: no setup needed

	// Hwnd states
	setupHwnds[HWND] = setupHwnd
	setupHwnds[HWNDPARENT] = setupHwndParent

	// Int states
	setupInts[SCREENWIDTH] = setupScreenWidth
	setupInts[SCREENHEIGHT] = setupScreenHeight
	setupInts[SCREENX] = setupScreenX
	setupInts[SCREENY] = setupScreenY
	setupInts[SCREENBPP] = setupScreenBPP
	setupInts[ORIGSCREENWIDTH] = setupOrigScreenWidth
	setupInts[ORIGSCREENHEIGHT] = setupOrigScreenHeight
	setupInts[FPS] = setupFPS
	setupInts[MINDELTATIME] = setupMinDeltaTime
	setupInts[SAMPLERATE] = setupSampleRate
	setupInts[FXVOLUME] = setupFxVolume
	setupInts[MUSVOLUME] = setupMusVolume
	setupInts[STREAMVOLUME] = setupStreamVolume
	setupInts[POWERSTATUS] = setupPowerStatus

	// String states
	setupStrings[ICON] = setupIcon
	setupStrings[TITLE] = setupTitle
	setupStrings[INIFILE] = setupInifile
	setupStrings[LOGFILE] = setupLogfile
}

// Sets internal system states.
// First param should be one of: BoolState, IntState, StringState, FuncState, HwndState
// Second parameter must be of the matching type, bool, int, string, StateFunc/func (h *HGE)() int, *Hwnd
func (h *HGE) SetState(a ...interface{}) error {
	if len(a) == 2 {
		switch state := a[0].(type) {
		case BoolState:
			if bs, ok := a[1].(bool); ok {
				return h.setStateBool(state, bs)
			}

		case IntState:
			if is, ok := a[1].(int); ok {
				return h.setStateInt(state, is)
			}

		case StringState:
			if ss, ok := a[1].(string); ok {
				return h.setStateString(state, ss)
			}
			if ss, ok := a[1].(*string); ok && ss != nil {
				return h.setStateString(state, *ss)
			} else {
				// A nil string state
				return h.setStateString(state, "")
			}

		case FuncState:
			switch a[1].(type) {
			case StateFunc:
				return h.setStateFunc(state, a[1].(StateFunc))
			case func() bool:
				return h.setStateFunc(state, a[1].(func() bool))
			default:
				return h.setStateFunc(state, nil)
			}

		case HwndState:
			switch a[1].(type) {
			case *Hwnd:
				return h.setStateHwnd(state, a[1].(*Hwnd))
			default:
				return h.setStateHwnd(state, nil)
			}
		}
	}

	return h.logError("Invalid arguments passed to SetState:", a...)
}

func (h *HGE) setStateBool(state BoolState, value bool) error {
	if state >= boolstate || state < 0 {
		h.Log("Invalid bool state")
		return h.logError("Invald bool state: %d %s", state, value)
	}

	stateBools[state] = value
	return setupBools[state](h)
}

func (h *HGE) setStateFunc(state FuncState, value StateFunc) error {
	if state >= funcstate || state < 0 {
		h.Log("Invalid function state")
		return h.logError("Invald function state: %d %s", state, value)
	}

	stateFuncs[state] = value

	return nil
}

func (h *HGE) setStateHwndPrivate(state HwndState, value *Hwnd) error {
	stateHwnds[state] = value

	return setupHwnds[state](h)
}

func (h *HGE) setStateHwnd(state HwndState, value *Hwnd) error {
	if state != HWNDPARENT {
		h.Log("Invalid hwnd state")
		return h.logError("Invald hwnd state: %d %s", state, value)
	}

	return h.setStateHwndPrivate(state, value)
}

func (h *HGE) setStateInt(state IntState, value int) error {
	if state >= intstate || state < 0 {
		h.Log("Invalid int state")
		return h.logError("Invald int state: %d %s", state, value)
	}

	stateInts[state] = value
	return setupInts[state](h)
}

func (h *HGE) setStateString(state StringState, value string) error {
	if state >= stringstate || state < 0 {
		h.Log("Invalid string state")
		return h.logError("Invald string state: %d %s", state, value)
	}

	stateStrings[state] = value
	return setupStrings[state](h)
}

// Returns internal system state values.
func (h *HGE) GetState(a ...interface{}) interface{} {
	if len(a) == 1 {
		switch a[0].(type) {
		case BoolState:
			return h.getStateBool(a[0].(BoolState))

		case IntState:
			return h.getStateInt(a[0].(IntState))

		case StringState:
			return h.getStateString(a[0].(StringState))

		case FuncState:
			return h.getStateFunc(a[0].(FuncState))

		case HwndState:
			return h.getStateHwnd(a[0].(HwndState))
		}
	}

	return nil
}

func (h *HGE) getStateBool(state BoolState) bool {
	if state >= boolstate || state < 0 {
		return false
	}

	return stateBools[state]
}

func (h *HGE) getStateFunc(state FuncState) StateFunc {
	if state >= funcstate || state < 0 {
		return nil
	}

	return stateFuncs[state]
}

func (h *HGE) getStateHwnd(state HwndState) Hwnd {
	return Hwnd{}
}

func (h *HGE) getStateInt(state IntState) int {
	if state >= intstate || state < 0 {
		return 0
	}

	return stateInts[state]
}

func (h *HGE) getStateString(state StringState) string {
	if state >= stringstate || state < 0 {
		return ""
	}

	return stateStrings[state]
}

func (h *HGE) setDefaultStates() {
	// Bool states
	h.SetState(WINDOWED, true)      // bool run in window? (default: true)
	h.SetState(ZBUFFER, false)      // bool use z-buffer? (default: false)
	h.SetState(TEXTUREFILTER, true) // bool texture filtering? (default: true)
	h.SetState(USESOUND, true)      // bool use sound? (default: true)
	h.SetState(DONTSUSPEND, false)  // bool focus lost:suspend? (default: false)
	h.SetState(HIDEMOUSE, true)     // bool hide system cursor? (default: true)
	h.SetState(SHOWSPLASH, true)    // bool show splash? (default: true)

	// Func States
	h.SetState(FRAMEFUNC, nil)      // func() bool frame function (default: nil) (you MUST set this)
	h.SetState(RENDERFUNC, nil)     // func() bool render function (default: nil)
	h.SetState(FOCUSLOSTFUNC, nil)  // func() bool focus lost function (default: nil)
	h.SetState(FOCUSGAINFUNC, nil)  // func() bool focus gain function (default: nil)
	h.SetState(GFXRESTOREFUNC, nil) // func() bool gfx restore function (default: nil)
	h.SetState(EXITFUNC, nil)       // func() bool exit function (default: nil)

	// Hwnd States
	h.setStateHwndPrivate(HWND, nil) // int		window handle: read only
	h.SetState(HWNDPARENT, nil)      // int		parent win handle	(default: 0)

	// Int states
	h.SetState(SCREENWIDTH, 800)                // int screen width (default: 800)
	h.SetState(SCREENHEIGHT, 600)               // int screen height (default: 600)
	h.SetState(SCREENX, sdl.WINDOWPOS_CENTERED) // int screen x (default: centered)
	h.SetState(SCREENY, sdl.WINDOWPOS_CENTERED) // int screen y (default: centered)
	h.SetState(SCREENBPP, 32)                   // int screen bitdepth (default: 32) (desktop bpp in windowed mode)
	h.SetState(SAMPLERATE, 44100)               // int sample rate (default: 44100)
	h.SetState(FXVOLUME, 100)                   // int global fx volume (default: 100)
	h.SetState(MUSVOLUME, 100)                  // int global music volume (default: 100)
	h.SetState(STREAMVOLUME, 100)               // int stream music volume (default: 100)
	h.SetState(FPS, FPS_UNLIMITED)              // int fixed fps (default: hge.FPS_UNLIMITED)
	h.SetState(MINDELTATIME, 1000)
	h.SetState(POWERSTATUS, 0)      // int battery life percent + status
	h.SetState(ORIGSCREENWIDTH, 0)  // int original screen width (default: 800 ... not valid until hge.System_Initiate()!)
	h.SetState(ORIGSCREENHEIGHT, 0) // int original screen height (default: 600 ... not valid until hge.System_Initiate()!))

	// String states
	h.SetState(ICON, "")     // string icon resource (default: nil)
	h.SetState(TITLE, "HGE") // string window title (default: "HGE")
	h.SetState(INIFILE, "")  // string ini file (default: nil) (meaning no file)
	h.SetState(LOGFILE, nil) // string log file (default: nil) (meaning no file)
}
