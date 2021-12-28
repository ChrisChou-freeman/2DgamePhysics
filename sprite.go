package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type Sprite struct {
	Texture       *ebiten.Image
	Position      *FPoint
	SpriteName    string
	CollisionInfo string
}

func (s *Sprite) GetRec() image.Rectangle {
	width, height := s.Texture.Size()
	rec := image.Rectangle{
		Min: image.Point{int(s.Position.X), int(s.Position.Y)},
		Max: image.Point{int(s.Position.X) + width, int(s.Position.Y) + height},
	}
	return rec
}

func (s *Sprite) Update() {}

func (s *Sprite) Draw(screen *ebiten.Image) {
	var iop *ebiten.DrawImageOptions = new(ebiten.DrawImageOptions)
	iop.GeoM.Translate(s.Position.X, s.Position.Y)
	screen.DrawImage(s.Texture, iop)
}

func (s *Sprite) Dispose() {
	s.Texture.Dispose()
}
