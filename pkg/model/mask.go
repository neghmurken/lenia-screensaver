package model

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/neghmurken/lenia-screensaver/pkg/utils"
)

type Mask map[int32]float32

func CreateMask(grid *Grid, x int32, y int32, r int32) Mask {
	mask := make(Mask, int(math.Pow(float64(r*2), 2)))

	for i := -r; i < r; i++ {
		for j := -r; j < r; j++ {
			length := rl.Normalize(
				float32(math.Sqrt(math.Pow(float64(i), 2)+math.Pow(float64(j), 2))),
				0,
				float32(r),
			)
			if length <= 1 {
				mask[grid.xy2i(i+x, j+y)] = utils.SquareDistribution(length)
			}
		}
	}

	return mask
}

func (m *Mask) Size() int {
	return len(*m)
}
