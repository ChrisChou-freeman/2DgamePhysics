package main

import(
  "image"
  "image/color"
  "math/rand"

  "github.com/hajimehoshi/ebiten/v2"
)

type Spring struct{
  circles []*Circle
  position *image.Point
}

func NewSpring() *Spring{
  s := new(Spring)
  s.Init()
  return s
}

func (s *Spring) Init(){
  s.circles = []*Circle{}
  s.position = &image.Point{400, 240}
}

func (s *Spring) UpdateSpring(c *Circle){
  c.Postion.X += c.Velocity.X
  c.Postion.Y += c.Velocity.Y
  c.Radius -= 0.1
  c.Velocity.Y += 0.15
}


func (s *Spring) Update(){
  radius := 5.0
  radius += float64(rand.Intn(5))
  vector := &FPoint{}
  vector.X = float64(rand.Intn(20))/10.0 - 1
  vector.Y = -5
  whiteRGBA := color.RGBA{255,255,255,255}
  newCircle := NewCircle(radius, &FPoint{float64(s.position.X), float64(s.position.Y)}, vector, whiteRGBA)
  s.circles = append(s.circles, newCircle)
  need_remove := []int{}
  for index, item := range(s.circles){
    s.UpdateSpring(item)
    if item.Radius <= 0{
      need_remove = append(need_remove, index)
    }
  }
  for index, item := range(need_remove){
    if item == len(s.circles) -1 {
      s.circles = s.circles[:item]
    }else{
      s.circles = append(s.circles[:item], s.circles[item+1:]...)
      for i:= index+1; i<len(need_remove)-1; i++{
        need_remove[i]--
      }
    }
  }
}

func (s *Spring) Draw(screen *ebiten.Image){
  for _, circle := range(s.circles){
    circle.Draw(screen)
  }
}
