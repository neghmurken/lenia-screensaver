package render

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"

	m "github.com/neghmurken/lenia-screensaver/pkg/model"
)

var (
	Gradient0 color.RGBA = color.RGBA{0, 0, 0, 255}
	Gradient1 color.RGBA = color.RGBA{0, 112, 255, 255}
	Gradient2 color.RGBA = color.RGBA{255, 0, 61, 255}
)

type Canvas struct {
	tex rl.Texture2D
}

func NewCanvas(state *m.State) *Canvas {
	image := rl.GenImageColor(
		int(state.Grid.W),
		int(state.Grid.H),
		rl.Black,
	)
	defer rl.UnloadImage(image)

	return &Canvas{
		tex: rl.LoadTextureFromImage(image),
	}
}

func (c *Canvas) Unload() {
	rl.UnloadTexture(c.tex)
}

func (c *Canvas) Render(state *m.State) {
	var pixels []rl.Color

	for _, c := range state.AllCells() {
		pixels = append(pixels, CellToColor(c))
	}

	rl.UpdateTexture(c.tex, pixels)
	pixels = nil

	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)
	rl.DrawTextureEx(
		c.tex,
		rl.Vector2Zero(),
		0,
		8,
		rl.White,
	)
	rl.DrawFPS(0, 0)
	rl.EndDrawing()
}

func CellToColor(cell float32) rl.Color {
	if cell < 0.5 {
		return LerpColor(
			Gradient0,
			Gradient1,
			cell*2,
		)
	}

	return LerpColor(
		Gradient1,
		Gradient2,
		(cell-0.5)*2,
	)
}

func LerpColor(left rl.Color, right rl.Color, factor float32) rl.Color {
	return rl.NewColor(
		uint8(rl.Lerp(float32(left.R), float32(right.R), factor)),
		uint8(rl.Lerp(float32(left.G), float32(right.G), factor)),
		uint8(rl.Lerp(float32(left.B), float32(right.B), factor)),
		255,
	)
}
