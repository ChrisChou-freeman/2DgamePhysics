package main 

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type GameManager interface {
  Init()
	Update()
	Draw(scrren *ebiten.Image)
}
