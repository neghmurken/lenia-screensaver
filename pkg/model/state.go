package model

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/neghmurken/lenia-screensaver/pkg/utils"
)

const MASK_RADIUS = 2

type State struct {
	Cells []float32
	Grid  *Grid
}

func NewState(w int32, h int32) *State {
	return &State{
		Cells: make([]float32, w*h),
		Grid:  &Grid{w, h},
	}
}

func (s *State) InitRandom() {
	for i := range s.Cells {
		s.Cells[i] = rand.Float32()
	}
}

func (s *State) Update(dt float32) {
	next := make([]float32, s.Grid.Size())
	for i, v := range s.Cells {
		x, y := s.Grid.i2xy(int32(i))
		growth := s.ApplyMask(CreateMask(s.Grid, x, y, MASK_RADIUS))
		next[i] = rl.Clamp(v+(growth*dt*0.1), 0, 1)
	}

	s.Cells = next
}

func (s *State) CellAt(x int32, y int32) float32 {
	return s.Cells[s.Grid.xy2i(x, y)]
}

func (s *State) AllCells() []float32 {
	return s.Cells
}

func (s *State) ApplyMask(mask Mask) float32 {
	applied := make([]float32, mask.Size())
	for i, v := range mask {
		applied = append(applied, s.Cells[i]*v)
	}

	return utils.Mean(applied)*2 - 1
}
