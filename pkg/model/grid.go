package model

type Grid struct {
	w int32
	h int32
}

func (g *Grid) Width() int32 {
	return g.w
}

func (g *Grid) Height() int32 {
	return g.h
}

func (g *Grid) Size() int32 {
	return g.w * g.h
}

func (g *Grid) i2xy(i int32) (x, y int32) {
	x = (i%g.w + g.w) % g.w
	y = (i/g.w + g.h) % g.h
	return
}

func (g *Grid) xy2i(x int32, y int32) int32 {
	return ((y+g.h)%g.h)*g.w + ((x + g.w) % g.w)
}
