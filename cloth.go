package main

import (
	"os"
	// "fmt"
	"encoding/json"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Rag struct {
	Points      [][]float64
	Connections [][]int
	Scale       int
	Grounded    []int
	Color       []int
}

type Cloth struct {
	rags       *Rag
	points     [][]float64
	origPoints [][]float64
	sticks     [][]float64
}

func NewCloth() *Cloth {
	c := new(Cloth)
	c.Init()
	return c
}

func (c *Cloth) Init() {
	c.rags = new(Rag)
	c.loadRags()
}

func (c *Cloth) loadRags() {
	rags, err := os.ReadFile("content/rags.json")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(rags, c.rags)
}

func (c *Cloth) Update() {
}

func (c *Cloth) Draw(screen *ebiten.Image) {
}
