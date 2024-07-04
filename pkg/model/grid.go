package model

type Grid struct {
	W int32
	H int32
}

func (g *Grid) Size() int32 {
	return g.W * g.H
}

func (g *Grid) I2XY(i int32) (x, y int32) {
	x = (i%g.W + g.H) % g.W
	y = (i/g.W + g.H) % g.H
	return
}

func (g *Grid) XY2I(x int32, y int32) int32 {
	return ((y+g.H)%g.H)*g.W + ((x + g.W) % g.W)
}
