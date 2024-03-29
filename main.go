package main

import (
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	physicsTypeList []GameManager
	testType        GameManager
	currentTest     int
}

func (g *Game) init() {
	g.physicsTypeList = []GameManager{
		NewThrow(),
		NewSpring(),
		NewFire(),
		NewExplode(),
	}
	g.testType = g.physicsTypeList[g.currentTest]
}

func (g *Game) keyEvent() {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.testType.Init()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		g.currentTest--
		if g.currentTest < 0 {
			g.currentTest = len(g.physicsTypeList) - 1
		}
		g.init()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		g.currentTest++
		if g.currentTest > len(g.physicsTypeList)-1 {
			g.currentTest = 0
		}
		g.init()
	}
}

func (g *Game) Update() error {
	g.keyEvent()
	g.testType.Update()
	return nil
}

func (g *Game) drawMessages(screen *ebiten.Image) {
	msg := "up/down key change\n r key replay"
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{180, 180, 180, 255})
	g.drawMessages(screen)
	g.testType.Draw(screen)
}

func (g *Game) Layout(width, height int) (int, int) {
	return 800, 480
}

func main() {
	ebiten.SetWindowSize(800, 480)
	ebiten.SetWindowTitle("test")
	game := new(Game)
	game.init()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
