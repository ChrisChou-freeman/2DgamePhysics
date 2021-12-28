package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type MotivationSprite struct {
	*Sprite
	Vector    *image.Point
	lifecycle int
	counter   int
	endLife   bool
}

func NewMotivationSprite(sprite *Sprite, lifeCircle int, vector *image.Point) *MotivationSprite {
	ms := new(MotivationSprite)
	ms.init(sprite, lifeCircle, vector)
	return ms
}

func (ms *MotivationSprite) init(sprite *Sprite, lifeCircle int, vector *image.Point) {
	ms.Sprite = sprite
	ms.lifecycle = lifeCircle
	ms.Vector = vector
}

func (ms *MotivationSprite) Update() {
	if ms.endLife {
		return
	}
	ms.Position.Add(FPoint{X: float64(ms.Vector.X), Y: float64(ms.Vector.Y)})
	ms.counter++
	if ms.counter >= ms.lifecycle {
		ms.endLife = true
	}
}

func (ms *MotivationSprite) Islife() bool {
	return !ms.endLife
}

func (ms *MotivationSprite) Kill() {
	ms.endLife = true
}

func (ms *MotivationSprite) Draw(screen *ebiten.Image) {
	if ms.Islife() {
		ms.Sprite.Draw(screen)
	}
}

// ------------------------

type ShootBullet struct {
	bulletList []*MotivationSprite
}

func NewShootBullet() *ShootBullet {
	newShootBult := new(ShootBullet)
	newShootBult.Init()
	return newShootBult
}

func (sb *ShootBullet) Init() {
}

func (sb *ShootBullet) Update() {
	for b := range sb.bulletList {
		fmt.Println(b)
	}
}
