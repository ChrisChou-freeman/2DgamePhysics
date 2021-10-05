package main

import (
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Circle struct {
	Radius         float64
	verticesNumber int
	Postion        *FPoint
	Velocity       *FPoint
	emptyImage     *ebiten.Image
	cColor         *color.RGBA
}

func NewCircle(radius float64, postion *FPoint, velocity *FPoint, cColor color.RGBA) *Circle {
	c := new(Circle)
	c.init(radius, postion, velocity, cColor)
	return c
}

func (c *Circle) init(radius float64, postion *FPoint, velocity *FPoint, cColor color.RGBA) {
	c.emptyImage = ebiten.NewImage(3, 3)
	c.emptyImage.Fill(color.White)
	c.Radius = radius
	c.Postion = postion
	c.verticesNumber = 100
	c.Velocity = velocity
	c.cColor = &cColor
}

func (c *Circle) genVertices() ([]ebiten.Vertex, []uint16) {
	vs := []ebiten.Vertex{}
	for i := 0; i < c.verticesNumber+1; i++ {
		rate := float64(i) / float64(c.verticesNumber)
		vs = append(vs, ebiten.Vertex{
			DstX:   float32(float64(c.Radius)*math.Cos(2*math.Pi*rate)) + float32(c.Postion.X),
			DstY:   float32(float64(c.Radius)*math.Sin(2*math.Pi*rate)) + float32(c.Postion.Y),
			SrcX:   0,
			SrcY:   0,
			ColorR: float32(c.cColor.R) / 255.0,
			ColorG: float32(c.cColor.G) / 255.0,
			ColorB: float32(c.cColor.B) / 255.0,
			ColorA: float32(c.cColor.A) / 255.0,
		})
	}
	indices := []uint16{}
	for i := 0; i < c.verticesNumber; i++ {
		indices = append(indices, uint16(i), uint16(i+1)%uint16(c.verticesNumber), uint16(c.verticesNumber))
	}
	return vs, indices
}

func (c *Circle) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawTrianglesOptions{}
	op.Address = ebiten.AddressUnsafe
	vertices, indices := c.genVertices()
	pix2Point := c.emptyImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
	screen.DrawTriangles(vertices, indices, pix2Point, op)
}
