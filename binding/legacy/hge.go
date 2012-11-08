package legacy

import (
	"github.com/losinggeneration/hge-go/binding/hge"
	. "github.com/losinggeneration/hge-go/binding/hge/gfx"
	. "github.com/losinggeneration/hge-go/binding/hge/ini"
	. "github.com/losinggeneration/hge-go/binding/hge/input"
	. "github.com/losinggeneration/hge-go/binding/hge/rand"
	. "github.com/losinggeneration/hge-go/binding/hge/resource"
	. "github.com/losinggeneration/hge-go/binding/hge/sound"
	. "github.com/losinggeneration/hge-go/binding/hge/timer"
)

// HGE struct
type HGE struct {
	h *hge.HGE
}

// Creates a new instance of an HGE structure
func Create(ver int) *HGE {
	return &HGE{hge.New()}
}

// Releases the memory the C++ library allocated for the HGE struct
func (h *HGE) Release() {
	h.h.Free()
}

// Initializes hardware and software needed to run engine.
func (h *HGE) System_Initiate() bool {
	return h.h.Initiate() == nil
}

//  Restores video mode and frees allocated resources.
func (h *HGE) System_Shutdown() {
	h.h.Shutdown()
}

// Starts running user defined frame function.
func (h *HGE) System_Start() bool {
	return h.h.Start() == nil
}

//  Returns last occured HGE error description.
func (h *HGE) System_GetErrorMessage() string {
	return h.h.GetErrorMessage()
}

// Writes a formatted message to the log file.
func (h *HGE) System_Log(format string, v ...interface{}) {
	h.h.Log(format, v...)
}

// Launches an URL or external executable/data file.
func (h *HGE) System_Launch(url string) bool {
	return h.h.Launch(url)
}

//  Saves current screen snapshot into a file.
func (h *HGE) System_Snapshot(a ...interface{}) {
	h.h.Snapshot(a...)
}

// Sets internal system states.
// First param should be one of: BoolState, IntState, StringState, FuncState, HwndState
// Second parameter must be of the matching type, bool, int, string, StateFunc/func() int, *Hwnd
func (h *HGE) System_SetState(a ...interface{}) {
	h.h.SetState(a...)
}

// Returns internal system state values.
func (h *HGE) System_GetState(a ...interface{}) interface{} {
	return h.h.GetState(a...)
}

// Loads a resource into memory from disk.
func (h *HGE) Resource_Load(filename string) (*Resource, hge.Dword) {
	return NewResource(filename)
}

// Deletes a previously loaded resource from memory.
func (h *HGE) Resource_Free(res Resource) {
	res.Free()
}

// Loads a resource, puts the loaded data into a byte array, and frees the data.
func (h *HGE) ResourceLoadBytes(filename string) []byte {
	return LoadBytes(filename)
}

// Loads a resource, puts the data into a string, and frees the data.
func (h *HGE) ResourceLoadString(filename string) *string {
	return LoadString(filename)
}

// Attaches a resource pack.
func (h *HGE) Resource_AttachPack(filename string, a ...interface{}) bool {
	return AttachPack(filename, a...)
}

// Removes a resource pack.
func (h *HGE) Resource_RemovePack(filename string) {
	RemovePack(filename)
}

// Removes all resource packs previously attached.
func (h *HGE) Resource_RemoveAllPacks() {
	RemoveAllPacks()
}

// Builds absolute file path.
func (h *HGE) Resource_MakePath(a ...interface{}) string {
	return MakePath(a...)
}

// Enumerates files by given wildcard.
func (h *HGE) Resource_EnumFiles(a ...interface{}) string {
	return EnumFiles(a...)
}

// Enumerates folders by given wildcard.
func (h *HGE) Resource_EnumFolders(a ...interface{}) string {
	return EnumFolders(a...)
}

func (h *HGE) Ini_SetInt(section, name string, value int) {
	NewIni(section, name).SetInt(value)
}

func (h *HGE) Ini_GetInt(section, name string, def_val int) int {
	return NewIni(section, name).GetInt(def_val)
}

func (h *HGE) Ini_SetFloat(section, name string, value float64) {
	NewIni(section, name).SetFloat(value)
}

func (h *HGE) Ini_GetFloat(section, name string, def_val float64) float64 {
	return NewIni(section, name).GetFloat(def_val)
}

func (h *HGE) Ini_SetString(section, name, value string) {
	NewIni(section, name).SetString(value)
}

func (h *HGE) Ini_GetString(section, name, def_val string) string {
	return NewIni(section, name).GetString(def_val)
}

var random = New(0)

func (h *HGE) Random_Seed(a ...interface{}) {
	seed := 1
	if len(a) == 1 {
		if s, ok := a[0].(int); ok {
			seed = s
		}
		if s, ok := a[0].(int64); ok {
			seed = int(s)
		}
	}

	random = New(seed)
	random.Seed()
}

func (h *HGE) Random_Int(min, max int) int {
	return random.Int(min, max)
}

func (h *HGE) Random_Float(min, max float64) float64 {
	return random.Float64(min, max)
}

func (h *HGE) Timer_GetTime() float64 {
	return Time()
}

func (h *HGE) Timer_GetDelta() float64 {
	return Delta()
}

func (h *HGE) Timer_GetFPS() int {
	return GetFPS()
}

func (h *HGE) Effect_Load(filename string, a ...interface{}) *Effect {
	return NewEffect(filename, a...)
}

func (h *HGE) Effect_Free(eff *Effect) {
	eff.Free()
}

func (h *HGE) Effect_Play(eff *Effect) Channel {
	return eff.Play()
}

func (h *HGE) Effect_PlayEx(eff *Effect, a ...interface{}) Channel {
	return eff.PlayEx(a...)
}

func (h *HGE) Music_Load(filename string, size hge.Dword) *Music {
	return NewMusic(filename, size)
}

func (h *HGE) Music_Free(music *Music) {
	music.Free()
}

func (h *HGE) Music_Play(music *Music, loop bool, a ...interface{}) Channel {
	return music.Play(loop, a...)
}

func (h *HGE) Music_SetAmplification(music *Music, ampl int) {
	music.SetAmplification(ampl)
}

func (h *HGE) Music_GetAmplification(music *Music) int {
	return music.Amplification()
}

func (h *HGE) Music_GetLength(music *Music) int {
	return music.Len()
}

func (h *HGE) Music_SetPos(music *Music, order, row int) {
	music.SetPos(order, row)
}

func (h *HGE) Music_GetPos(music *Music) (order, row int, ok bool) {
	return music.Pos()
}

func (h *HGE) Music_SetInstrVolume(music *Music, instr int, volume int) {
	music.SetInstrVolume(instr, volume)
}

func (h *HGE) Music_GetInstrVolume(music *Music, instr int) int {
	return music.InstrVolume(instr)
}

func (h *HGE) Music_SetChannelVolume(music *Music, channel, volume int) {
	music.SetChannelVolume(channel, volume)
}

func (h *HGE) Music_GetChannelVolume(music *Music, channel int) int {
	return music.ChannelVolume(channel)
}

func (h *HGE) Stream_Load(filename string, size hge.Dword) *Stream {
	return NewStream(filename, size)
}

func (h *HGE) Stream_Free(stream *Stream) {
	stream.Free()
}

func (h *HGE) Stream_Play(stream *Stream, loop bool, a ...interface{}) Channel {
	return stream.Play(loop, a...)
}

func (h *HGE) Channel_SetPanning(chn Channel, pan int) {
	chn.SetPanning(pan)
}

func (h *HGE) Channel_SetVolume(chn Channel, volume int) {
	chn.SetVolume(volume)
}

func (h *HGE) Channel_SetPitch(chn Channel, pitch float64) {
	chn.SetPitch(pitch)
}

func (h *HGE) Channel_Pause(chn Channel) {
	chn.Pause()
}

func (h *HGE) Channel_Resume(chn Channel) {
	chn.Resume()
}

func (h *HGE) Channel_Stop(chn Channel) {
	chn.Stop()
}

func (h *HGE) Channel_PauseAll() {
	PauseAll()
}

func (h *HGE) Channel_ResumeAll() {
	ResumeAll()
}

func (h *HGE) Channel_StopAll() {
	StopAll()
}

func (h *HGE) Channel_IsPlaying(chn Channel) bool {
	return chn.IsPlaying()
}

func (h *HGE) Channel_GetLength(chn Channel) float64 {
	return chn.Len()
}

func (h *HGE) Channel_GetPos(chn Channel) float64 {
	return chn.Pos()
}

func (h *HGE) Channel_SetPos(chn Channel, seconds float64) {
	chn.SetPos(seconds)
}

func (h *HGE) Channel_SlideTo(chn Channel, time float64, a ...interface{}) {
	chn.SlideTo(time, a...)
}

func (h *HGE) Channel_IsSliding(chn Channel) bool {
	return chn.IsSliding()
}

func (h *HGE) Input_GetMousePos() (x, y float64) {
	return NewMouse(0, 0).Pos()
}

func (h *HGE) Input_SetMousePos(x, y float64) {
	Mouse{}.SetPos(x, y)
}

func (h *HGE) Input_GetMouseWheel() int {
	return NewMouse(0, 0).WheelMovement()
}

func (h *HGE) Input_IsMouseOver() bool {
	return NewMouse(0, 0).IsOver()
}

func (h *HGE) Input_KeyDown(key int) bool {
	return NewKey(key).Down()
}

func (h *HGE) Input_KeyUp(key int) bool {
	return NewKey(key).Up()
}

func (h *HGE) Input_GetKeyState(key int) bool {
	return NewKey(key).State()
}

func (h *HGE) Input_GetKeyName(key int) string {
	return NewKey(key).Name()
}

func (h *HGE) Input_GetKey() int {
	return int(GetKey())
}

func (h *HGE) Input_GetChar() int {
	return GetChar()
}

func (h *HGE) Input_GetEvent(event *InputEvent) bool {
	event, b := GetEvent()
	return b
}

func (h *HGE) Gfx_BeginScene(a ...interface{}) bool {
	return BeginScene(a...)
}

func (h *HGE) Gfx_EndScene() {
	EndScene()
}

func (h *HGE) Gfx_Clear(color hge.Dword) {
	Clear(color)
}

func (h *HGE) Gfx_RenderLine(x1, y1, x2, y2 float64, a ...interface{}) {
	NewLine(x1, y1, x2, y2, a...).Render()
}

func (h *HGE) Gfx_RenderTriple(triple *Triple) {
	triple.Render()
}

func (h *HGE) Gfx_RenderQuad(quad *Quad) {
	quad.Render()
}

func (h *HGE) Gfx_StartBatch(prim_type int, tex *Texture, blend int) (ver *Vertex, max_prim int, ok bool) {
	return StartBatch(prim_type, tex, blend)
}

func (h *HGE) Gfx_FinishBatch(prim int) {
	FinishBatch(prim)
}

func (h *HGE) Gfx_SetClipping(a ...interface{}) {
	SetClipping(a...)
}

func (h *HGE) Gfx_SetTransform(a ...interface{}) {
	SetTransform(a...)
}

func (h *HGE) Target_Create(width, height int, zbuffer bool) *Target {
	return NewTarget(width, height, zbuffer)
}

func (h *HGE) Target_Free(target *Target) {
	target.Free()
}

func (h *HGE) Target_GetTexture(target Target) *Texture {
	return target.Texture()
}

func (h *HGE) Texture_Create(width, height int) *Texture {
	return NewTexture(width, height)
}

func (h *HGE) Texture_Load(filename string, a ...interface{}) *Texture {
	return LoadTexture(filename, a...)
}

func (h *HGE) Texture_Free(tex *Texture) {
	tex.Free()
}

func (h *HGE) Texture_GetWidth(tex Texture, a ...interface{}) int {
	return tex.Width(a...)
}

func (h *HGE) Texture_GetHeight(tex Texture, a ...interface{}) int {
	return tex.Height(a...)
}

func (h *HGE) Texture_Lock(tex Texture, a ...interface{}) *hge.Dword {
	return tex.Lock(a...)
}

func (h *HGE) Texture_Unlock(tex Texture) {
	tex.Unlock()
}