package main

import(
  "image"

  "github.com/hajimehoshi/ebiten/v2"
)

type Explode struct{
  position *image.Point
  circles []*Circle
}

func NewExplode() *Explode{
  e := new(Explode)
  e.Init()
  return e
}

func (e *Explode) Init(){
  e.position = &image.Point{400, 240}
  e.circles = []*Circle{}
}

func (e *Explode) Update(){
}

func (e *Explode) Draw(screen *ebiten.Image){
}
