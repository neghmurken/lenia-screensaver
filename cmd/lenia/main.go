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

	rl.ToggleFullscreen()
	state := m.NewState(int32(rl.GetScreenWidth()/8), int32(rl.GetScreenHeight()/8))

	state.InitRandom()
	state.InitMask(rl.GetRandomValue(0, state.Grid.W), rl.GetRandomValue(0, state.Grid.H), m.MASK_RADIUS+12)
	state.InitMask(rl.GetRandomValue(0, state.Grid.W), rl.GetRandomValue(0, state.Grid.H), m.MASK_RADIUS+6)

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
