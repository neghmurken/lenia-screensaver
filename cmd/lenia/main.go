package main

import (
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"

	m "github.com/neghmurken/lenia-screensaver/pkg/model"
	r "github.com/neghmurken/lenia-screensaver/pkg/render"
)

func main() {
	rl.InitWindow(0, 0, "Lenia")
	defer rl.CloseWindow()

	rl.ToggleFullscreen()

	state := m.NewState(400, 400)
	state.InitRandom()
	canvas := r.NewCanvas(state)
	defer canvas.Unload()

	rl.SetTargetFPS(20)

	shouldExit := false

	for !shouldExit {
		if rl.IsKeyPressed(rl.KeyEscape) || rl.WindowShouldClose() {
			shouldExit = true
		}

		state.Update(rl.GetFrameTime())
		canvas.Render(state)
	}

	os.Exit(0)
}
