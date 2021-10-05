package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"math/rand"
)

type Fire struct {
	circles  []*Circle
	position *FPoint
}

func NewFire() *Fire {
	f := new(Fire)
	f.Init()
	return f
}

func (f *Fire) Init() {
	f.position = &FPoint{400, 240}
	f.circles = []*Circle{}
}

func (f *Fire) UpdateFire(c *Circle) {
	c.Postion.Y += c.Velocity.Y
	c.Postion.X += c.Velocity.X
	c.Radius -= 0.28
	if c.Velocity.Y < 0 {
		c.Velocity.Y += 0.005
	}
}

func (f *Fire) Update() {
	radius := 5.0
	radius += float64(rand.Intn(5))
	vector := &FPoint{}
	vector.X = float64(rand.Intn(10))/10.0 - 0.5
	vector.Y = -2.5
	colorFire := color.RGBA{255, 175, 116, 255}
	newCircle := NewCircle(radius, &FPoint{float64(f.position.X), float64(f.position.Y)}, vector, colorFire)
	f.circles = append(f.circles, newCircle)
	need_remove := []int{}
	for index, item := range f.circles {
		f.UpdateFire(item)
		if item.Radius <= 0 {
			need_remove = append(need_remove, index)
		}
	}
	for index, item := range need_remove {
		if item == len(f.circles)-1 {
			f.circles = f.circles[:item]
		} else {
			f.circles = append(f.circles[:item], f.circles[item+1:]...)
			for i := index + 1; i < len(need_remove); i++ {
				need_remove[i]--
			}
		}
	}
}

func (f *Fire) Draw(screen *ebiten.Image) {
	for _, circle := range f.circles {
		circle.Draw(screen)
	}
}
