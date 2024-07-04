package model

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Mask struct {
	Mask   map[int32]float32
	Grid   Grid
	radius int32
}

func CreateMask(r int32) *Mask {
	d := r*2 + 1
	grid := Grid{d, d}
	mask := make(map[int32]float32, int(math.Pow(float64(d), 2)))

	for i := int32(0); i < d; i++ {
		for j := int32(0); j < d; j++ {
			length := rl.Normalize(
				float32(math.Sqrt(math.Pow(float64(i-r), 2)+math.Pow(float64(j-r), 2))),
				0,
				float32(r),
			)
			if length <= 1 {
				mask[grid.XY2I(i, j)] = 1 //utils.QuadBell(length)
			}
		}
	}

	mask[grid.XY2I(r+1, r+1)] = 0.0

	return &Mask{
		Mask:   mask,
		Grid:   grid,
		radius: r,
	}
}
