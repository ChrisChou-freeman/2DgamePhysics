package main

type FPoint struct {
	X float64
	Y float64
}

func (fp *FPoint) Add(fpoint FPoint) {
	fp.X += fpoint.X
	fp.Y += fpoint.Y
}
