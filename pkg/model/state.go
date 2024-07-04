package model

import (
	"math/rand"
	"sync"

	rl "github.com/gen2brain/raylib-go/raylib"

	u "github.com/neghmurken/lenia-screensaver/pkg/utils"
)

const MASK_RADIUS = 7

type State struct {
	cells []float32
	Grid  *Grid
	Mask  *Mask
}

func NewState(w, h int32) *State {
	return &State{
		cells: make([]float32, w*h),
		Grid:  &Grid{w, h},
		Mask:  CreateMask(MASK_RADIUS),
	}
}

func (s *State) InitRandom() {
	for i := range s.cells {
		s.cells[i] = rand.Float32() * 0.54
	}
}

func (s *State) InitMask(x, y, size int32) {
	mask := CreateMask(size)

	for i, m := range mask.Mask {
		u, v := mask.Grid.I2XY(i)
		s.cells[s.Grid.XY2I(u+x-size, v+y-size)] = m * rand.Float32()
	}
}

func (s *State) Update(dt float32) {
	next := make([]float32, s.Grid.Size())

	var wg sync.WaitGroup
	for x := int32(0); x < s.Grid.W; x++ {
		wg.Add(1)
		go func(x int32) {
			defer wg.Done()
			for y := int32(0); y < s.Grid.H; y++ {
				others := s.Convolve(s.Mask, x, y)
				i := s.Grid.XY2I(x, y)
				next[i] = rl.Clamp(Grow(s.cells[i], others, dt), 0, 1)
			}
		}(x)
	}
	wg.Wait()

	s.cells = next
}

func (s *State) CellAt(x, y int32) float32 {
	return s.cells[s.Grid.XY2I(x, y)]
}

func (s *State) AllCells() []float32 {
	return s.cells
}

func (s *State) Convolve(mask *Mask, x, y int32) float32 {
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

func Grow(life, others, dt float32) float32 {
	return life + u.GaussBell(others, 2, 0.3, 0.1)*dt
}
