package model

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"

	u "github.com/neghmurken/lenia-screensaver/pkg/utils"
)

const MASK_RADIUS = 1

type State struct {
	cells []float32
	Grid  *Grid
	Mask  *Mask
}

func NewState(w int32, h int32) *State {
	return &State{
		cells: make([]float32, w*h),
		Grid:  &Grid{w, h},
		Mask:  CreateMask(MASK_RADIUS),
	}
}

func (s *State) InitRandom() {
	for i := range s.cells {
		s.cells[i] = rand.Float32() * 0.8
	}
}

func (s *State) InitMask(x int32, y int32, size int32) {
	mask := CreateMask(size)

	for i, m := range mask.Mask {
		u, v := mask.Grid.I2XY(i)
		s.cells[s.Grid.XY2I(u+x-size, v+y-size)] = m
	}
}

func (s *State) Update(dt float32) {
	next := make([]float32, s.Grid.Size())
	for i, life := range s.cells {
		x, y := s.Grid.I2XY(int32(i))
		others := s.Convolve(s.Mask, x, y)
		next[i] = rl.Clamp(Grow(life, others, dt), 0, 1)
	}

	s.cells = next
}

func (s *State) CellAt(x int32, y int32) float32 {
	return s.cells[s.Grid.XY2I(x, y)]
}

func (s *State) AllCells() []float32 {
	return s.cells
}

func (s *State) Convolve(mask *Mask, x int32, y int32) float32 {
	growth := float32(0)
	for i, v := range mask.Mask {
		u, w := mask.Grid.I2XY(i)
		growth += s.cells[s.Grid.XY2I(
			u-mask.radius+x,
			w-mask.radius+y,
		)] * v
	}

	return growth / float32(mask.Grid.Size())
}

func Grow(life float32, others float32, dt float32) float32 {
	return life + u.GaussBell(others, 2, 0.3, 0.08)*dt
}
