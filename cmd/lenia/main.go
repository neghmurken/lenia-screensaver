package main

import (
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"

	m "github.com/neghmurken/lenia-screensaver/pkg/model"
	r "github.com/neghmurken/lenia-screensaver/pkg/render"
)

func main() {
	rl.InitWindow(0, 0, "Lenia")
	rl.SetConfigFlags(rl.FlagVsyncHint)
	defer rl.CloseWindow()

	//rl.ToggleFullscreen()
	state := m.NewState(800, 800)

	state.InitMask(400, 400, 400)

	canvas := r.NewCanvas(state)
	defer canvas.Unload()

	rl.SetTargetFPS(60)

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
