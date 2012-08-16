package main

import (
	"fmt"
	dist "github.com/losinggeneration/hge-go/helpers/distortionmesh"
	"github.com/losinggeneration/hge-go/helpers/font"
	HGE "github.com/losinggeneration/hge-go/hge"
	hge "github.com/losinggeneration/hge-go/legacy"
	"math"
)

var (
	h   *hge.HGE
	tex HGE.Texture
	dis dist.DistortionMesh
	fnt *font.Font
)

const (
	rows  = 16
	cols  = 16
	cellw = 512.0 / (cols - 1)
	cellh = 512.0 / (rows - 1)

	meshx = 144.0
	meshy = 44.0
)

var trans = 0
var t = 0.0

func FrameFunc() int {
	t += h.Timer_GetDelta()

	// Process keys
	switch h.Input_GetKey() {
	case HGE.K_ESCAPE:
		return 1

	case HGE.K_SPACE:
		trans++

		if trans > 2 {
			trans = 0
		}

		dis.Clear(HGE.Dword(0xFF000000))
	}

	// Calculate new displacements and coloring for one of the three effects
	switch trans {
	case 0:
		for i := 1; i < rows-1; i++ {
			for j := 1; j < cols-1; j++ {
				dis.SetDisplacement(j, i, math.Cos(t*float64(10+(i+j))/2)*5, math.Sin(t*float64(10+(i+j))/2)*5, dist.DISP_NODE)
			}
		}

	case 1:
		for i := 0; i < rows; i++ {
			for j := 1; j < cols-1; j++ {
				dis.SetDisplacement(j, i, math.Cos(t*float64(5+j)/2)*15, 0, dist.DISP_NODE)
				col := HGE.Dword((math.Cos(t*float64(5+(i+j))/2) + 1) * 35)
				dis.SetColor(j, i, 0xFF<<24|col<<16|col<<8|col)
			}
		}

	case 2:
		for i := 0.0; i < rows; i++ {
			for j := 0.0; j < cols; j++ {
				r := math.Sqrt(math.Pow(j-float64(cols)/2, 2) + math.Pow(i-float64(rows)/2, 2))
				a := r * math.Cos(t*2) * 0.1
				dx := math.Sin(a)*(i*cellh-256) + math.Cos(a)*(j*cellw-256)
				dy := math.Cos(a)*(i*cellh-256) - math.Sin(a)*(j*cellw-256)
				dis.SetDisplacement(int(j), int(i), dx, dy, dist.DISP_CENTER)
				col := HGE.Dword((math.Cos(r+t*4) + 1) * 40)
				dis.SetColor(int(j), int(i), 0xFF<<24|col<<16|(col/2)<<8)
			}
		}
	}

	return 0
}

func RenderFunc() int {
	// Render graphics
	h.Gfx_BeginScene()
	h.Gfx_Clear(0)
	dis.Render(meshx, meshy)
	fnt.Printf(5, 5, font.TEXT_LEFT, "dt:%.3f\nFPS:%d\n\nUse your\nSPACE!", h.Timer_GetDelta(), h.Timer_GetFPS())
	h.Gfx_EndScene()

	return 0
}

func main() {
	h = hge.Create(HGE.VERSION)
	defer h.Release()

	h.System_SetState(HGE.LOGFILE, "tutorial05.log")
	h.System_SetState(HGE.FRAMEFUNC, FrameFunc)
	h.System_SetState(HGE.RENDERFUNC, RenderFunc)
	h.System_SetState(HGE.TITLE, "HGE Tutorial 05 - Using distortion mesh")
	h.System_SetState(HGE.WINDOWED, true)
	h.System_SetState(HGE.SCREENWIDTH, 800)
	h.System_SetState(HGE.SCREENHEIGHT, 600)
	h.System_SetState(HGE.SCREENBPP, 32)
	h.System_SetState(HGE.USESOUND, false)

	if h.System_Initiate() {
		defer h.System_Shutdown()
		tex = h.Texture_Load("texture.jpg")
		if tex == 0 {
			fmt.Println("Error: Can't load texture.jpg")
			return
		}
		defer h.Texture_Free(tex)

		dis = dist.NewDistortionMesh(cols, rows)
		dis.SetTexture(tex)
		dis.SetTextureRect(0, 0, 512, 512)
		dis.SetBlendMode(HGE.BLEND_COLORADD | HGE.BLEND_ALPHABLEND | HGE.BLEND_ZWRITE)
		dis.Clear(HGE.Dword(0xFF000000))

		// Load a font
		fnt = font.NewFont("font1.fnt")

		if fnt == nil {
			fmt.Println("Error: Can't load font1.fnt or font1.png")
			return
		}

		// Let's rock now!
		h.System_Start()
	}
}
