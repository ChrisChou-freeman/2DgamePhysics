package main

import (
	// "fmt"
	"image"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Explode struct {
	position   *image.Point
	circles    []*Circle
	randValues float64
	expandTime int
	explodeDot int
}

func NewExplode() *Explode {
	e := new(Explode)
	e.Init()
	return e
}

func (e *Explode) Init() {
	e.position = &image.Point{400, 240}
	e.circles = []*Circle{}
	e.explodeDot = 70
	e.expandTime = 240
	e.LoadExplodeDot()
}

func (e *Explode) LoadExplodeDot() {
	for i := 0; i < e.explodeDot; i++ {
		rand.Seed(int64(i))
		radius := 30.0 + float64(rand.Intn(10))
		vector := &FPoint{}
		vector.X = float64(rand.Intn(40)/5 - 4)
		vector.Y = float64(-rand.Intn(6))
		colorFire := color.RGBA{162, 162, 162, 255}
		offsetX := rand.Intn(10)
		offsetY := rand.Intn(30)
		if i%2 == 0 {
			offsetX *= -1
		}
		explodPosition := &FPoint{float64(e.position.X + offsetX), float64(e.position.Y + offsetY)}
		newCircle := NewCircle(radius, explodPosition, vector, colorFire)
		e.circles = append(e.circles, newCircle)
	}
}

func (e *Explode) UpdateExplode(c *Circle) {
	c.Postion.Y += c.Velocity.Y
	c.Postion.X += c.Velocity.X
	randValue := rand.Float64() * float64(rand.Intn(10))
	if e.expandTime > 0 {
		c.Radius += randValue
	} else {
		c.Radius -= randValue
	}
	c.Velocity.Y += 0.03
	e.expandTime--
	c.cColor.A -= 2
}

func (e *Explode) Update() {
	need_remove := []int{}
	for index, circle := range e.circles {
		e.UpdateExplode(circle)
		if circle.Radius <= 0 {
			need_remove = append(need_remove, index)
		}
	}
	for index, cIndex := range need_remove {
		if cIndex == len(e.circles)-1 {
			e.circles = e.circles[:cIndex]
		} else {
			e.circles = append(e.circles[:cIndex], e.circles[cIndex+1:]...)
			for i := index + 1; i < len(need_remove); i++ {
				need_remove[i]--
			}
		}
	}
}

func (e *Explode) Draw(screen *ebiten.Image) {
	for _, circle := range e.circles {
		circle.Draw(screen)
	}
}
